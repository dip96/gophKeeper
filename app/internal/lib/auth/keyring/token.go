package keyring

//TODO error - failed to unlock correct collection '/org/freedesktop/secrets/aliases/default'
//import (
//	"context"
//	"github.com/zalando/go-keyring"
//)
//
//const (
//	serviceName = "GophKeeper"
//	tokenKey    = "AuthToken"
//)
//
//func SaveToken(token string) error {
//	return keyring.Set(serviceName, tokenKey, token)
//}
//
//func LoadToken() (string, error) {
//	return keyring.Get(serviceName, tokenKey)
//}
//
//func DeleteToken() error {
//	return keyring.Delete(serviceName, tokenKey)
//}
//
////реализует grpc.PerRPCCredentials для аутентификации запросов
//type TokenAuth struct {
//	token string
//}
//
//func (t TokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
//	return map[string]string{
//		"authorization": "Bearer " + t.token,
//	}, nil
//}
//
//func (TokenAuth) RequireTransportSecurity() bool {
//	return false
//}
