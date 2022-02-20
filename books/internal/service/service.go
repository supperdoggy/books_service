package service

import (
	"context"
	"github.com/supperdoggy/grpc_CRUD/book_proto"
	db2 "github.com/supperdoggy/grpc_CRUD/books/internal/db"
	"github.com/supperdoggy/grpc_CRUD/books/internal/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IService interface {
	UploadBook(context.Context, *book_proto.UploadBookRequest) (*book_proto.UploadBookResponse, error)
}

type service struct {
	l  *zap.Logger
	db db2.IDB
}

func NewService(l *zap.Logger, d db2.IDB) IService {
	return &service{l: l, db: d}
}

func (s *service) UploadBook(ctx context.Context, in *book_proto.UploadBookRequest) (*book_proto.UploadBookResponse, error) {
	if !utils.CreateAndCompareSHA256(in.GetData(), in.GetHash()) {
		return nil, status.Error(codes.DataLoss, "different hashes got")
	}

	err := s.db.NewBook(in.GetName(), in.GetAuthor(), in.GetData())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := book_proto.UploadBookResponse{Ok: true}
	return &out, nil
}
