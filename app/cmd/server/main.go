package main

import (
	"google.golang.org/grpc"
	"gophKeeper/internal/config"
	auth "gophKeeper/internal/lib/auth/jwt"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/migrator"
	cGrpc "gophKeeper/internal/server/grpc"
	servicesGrpc "gophKeeper/internal/server/grpc/services"
	userService "gophKeeper/internal/service/user"
	"gophKeeper/internal/storage/postgres"
	"gophKeeper/internal/uow"
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

	//TODO Добавить в конфиги
	auth.Init("secret")

	// init and start server
	log.Info("starting server")
	srv := cGrpc.NewGRPCServer(cnf)

	//TODO переместить в отдельный слой
	uowService := uow.NewUnitOfWork(storage)
	userSrv := servicesGrpc.NewUserService(userService.NewUserService(uowService))

	srv.RegisterService(func(grpcServer *grpc.Server) {
		pU.RegisterUserServiceServer(grpcServer, userSrv)
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
