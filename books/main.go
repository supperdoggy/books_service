package main

import (
	"fmt"
	"github.com/supperdoggy/grpc_CRUD/book_proto"
	"github.com/supperdoggy/grpc_CRUD/books/internal/cfg"
	db2 "github.com/supperdoggy/grpc_CRUD/books/internal/db"
	service2 "github.com/supperdoggy/grpc_CRUD/books/internal/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err.Error())
	}

	config := cfg.GetConfig()
	db, err := db2.NewDB(logger, config.DBUrl, config.DBName,
		config.BooksCollectionName, config.BooksDataCollectionName,
		config.ScoresCollectionName, config.CurrentPageCollectionName)
	service := service2.NewService(logger, db)
	if err != nil {
		logger.Fatal("error getting db", zap.Error(err))
	}

	urlport := fmt.Sprintf("%s:%s", config.Url, config.Port)
	listener, err := net.Listen("tcp", urlport)
	if err != nil {
		logger.Fatal("error listening to adress", zap.Error(err))
	}

	grpcServer := grpc.NewServer()

	book_proto.RegisterBooksServiceServer(grpcServer, service)

	logger.Info("serving the server", zap.Any("listener", listener.Addr()))

	if err := grpcServer.Serve(listener); err != nil {
		logger.Fatal("error serving server", zap.Error(err))
	}

}
