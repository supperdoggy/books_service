package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/supperdoggy/grpc_CRUD/book_proto"
	cfg "github.com/supperdoggy/grpc_CRUD/book_upload_client/internal/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"io/ioutil"
	"time"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err.Error())
	}
	config := cfg.GetConfig()

	urlport := fmt.Sprintf("%s:%s", config.Url, config.Port)
	conn, err := grpc.Dial(urlport, grpc.WithInsecure())
	if err != nil {
		logger.Fatal("error dialing with server", zap.Error(err))
	}

	client := book_proto.NewBooksServiceClient(conn)

	bts, err := ioutil.ReadFile(config.PathToTxt)
	if err != nil {
		logger.Fatal("error ")
	}

	hash := sha256.Sum256(bts)
	in := book_proto.UploadBookRequest{
		Name: "War and Peace",
		Author: "Lev Tolstoy",
		Data: bts,
		Hash: hex.EncodeToString(hash[:]),
		CreationTime: time.Date(1873, time.April, 1, 12, 0, 0, 0, time.UTC).Unix(),
		SecurityToken: config.SecurityToken,
	}

	out, err := client.UploadBook(context.Background(), &in)
	if err != nil {
		logger.Fatal("error while uploading book", zap.Error(err))
	}

	logger.Info("out", zap.Any("out", out))
}
