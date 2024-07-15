package main

import (
	"gophKeeper/internal/config"
	"gophKeeper/internal/logger"
	"gophKeeper/internal/migrator"
	"gophKeeper/internal/server/grpc"
	"gophKeeper/internal/storage/postgres"
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
	_, err = postgres.NewDB(cnf)

	if err != nil {
		log.Error("failed to create storage:", err)
		return
	}

	// init and start server
	log.Info("starting server")
	srv := grpc.NewGRPCServer(cnf)
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
