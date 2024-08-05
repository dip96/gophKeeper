package interceptors

import (
	"context"
	"github.com/go-chi/jwtauth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gophKeeper/internal/lib/auth/jwt"
	"strings"
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

	values := md.Get("authorization")
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	authHeader := values[0]
	bearerToken := strings.TrimPrefix(authHeader, "Bearer ")
	if bearerToken == authHeader {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token format")
	}

	claims, err := auth.VerifyToken(bearerToken)
	if err != nil {
		if err == jwtauth.ErrExpired {
			return nil, status.Errorf(codes.Unauthenticated, "token has expired")
		}
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	newCtx := context.WithValue(ctx, "user_id", claims["user_id"])

	return handler(newCtx, req)
}
