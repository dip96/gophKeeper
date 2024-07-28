package interceptors

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gophKeeper/internal/lib/auth/jwt"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// Список публичных методов, которые не требуют аутентификации
	publicMethods := map[string]bool{
		"/v1.users.UserService/Registration": true,
		"/v1.users.UserService/Login":        true,
	}

	// Если метод публичный, пропускаем аутентификацию
	if publicMethods[info.FullMethod] {
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	token := md.Get("authorization")
	if len(token) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	claims, err := auth.VerifyToken(token[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token")
	}

	newCtx := context.WithValue(ctx, "user_id", claims["user_id"])

	return handler(newCtx, req)
}
