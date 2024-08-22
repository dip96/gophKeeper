package main

import (
	"crypto/tls"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gophKeeper/internal/config"
	auth "gophKeeper/internal/lib/auth/jwt"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/migrator"
	cGrpc "gophKeeper/internal/server/grpc"
	servicesGrpc "gophKeeper/internal/server/grpc/services"
	binaryDataService "gophKeeper/internal/service/binary_data"
	creditCardDataService "gophKeeper/internal/service/credit_card_data"
	loginDataService "gophKeeper/internal/service/login_data"
	textDataService "gophKeeper/internal/service/text_data"
	userService "gophKeeper/internal/service/user"
	"gophKeeper/internal/storage/postgres"
	"gophKeeper/internal/uow"
	pBd "gophKeeper/protobuf/V1/binary_data"
	pCc "gophKeeper/protobuf/V1/credit_card_data"
	pLd "gophKeeper/protobuf/V1/login_data"
	pTd "gophKeeper/protobuf/V1/text_data"
	pU "gophKeeper/protobuf/V1/users"
	bLog "log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//init config
	cnf := config.MustLoad()

	//init logger
	log, err := logger.Init()
	if err != nil {
		bLog.Println(err.Error())
		return
	}

	//init migrator
	log.Info("init migrator")
	migIns, err := migrator.NewMigrator(cnf)
	if err != nil {
		log.Error("failed to create migrator:", err)
		return
	}

	err = migIns.Up()
	if err != nil {
		log.Error("Error migrator")
		panic("Error migrator")
	}

	//init storage
	log.Info("init storage")
	storage, err := postgres.NewDB(cnf)
	if err != nil {
		log.Error("failed to create storage:", err)
		return
	}

	auth.Init("secret")

	tlsCredentials, err := loadTLSCredentials(cnf)
	if err != nil {
		log.Error("failed to load TLS credentials:", err)
		return
	}

	srv := cGrpc.NewGRPCServer(cnf, grpc.Creds(tlsCredentials))

	uowService := uow.NewUnitOfWork(storage)
	userSrv := servicesGrpc.NewUserService(userService.NewUserService(uowService))
	loginDataSrv := servicesGrpc.NewLoginDataService(loginDataService.NewLoginDataService(uowService))
	binaryDataSrv := servicesGrpc.NewBinaryDataService(binaryDataService.NewBinaryDataService(uowService))
	textDataSrv := servicesGrpc.NewTextDataService(textDataService.NewTextDataService(uowService))
	creditCardDataSrv := servicesGrpc.NewCreditCardDataService(creditCardDataService.NewCreditCardDataService(uowService))

	srv.RegisterService(func(grpcServer *grpc.Server) {
		pU.RegisterUserServiceServer(grpcServer, userSrv)
		pLd.RegisterLoginDataServiceServer(grpcServer, loginDataSrv)
		pBd.RegisterBinaryDataServiceServer(grpcServer, binaryDataSrv)
		pTd.RegisterTextDataServiceServer(grpcServer, textDataSrv)
		pCc.RegisterCreditCardDataServiceServer(grpcServer, creditCardDataSrv)
	})

	go func() {
		if err := srv.Start(); err != nil {
			log.Error("failed to start server:", err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("shutting down server")
	if err := srv.Stop(); err != nil {
		log.Error("failed to stop server:", err)
	}
}

func loadTLSCredentials(cfg config.ConfigServer) (credentials.TransportCredentials, error) {
	certPath := cfg.GetSslCert()
	keyPath := cfg.GetSslKey()

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load key pair: %w", err)
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	return credentials.NewTLS(config), nil
}
