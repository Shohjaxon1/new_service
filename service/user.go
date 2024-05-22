package service

import (
	"context"
	pb "new_service/genproto"
	l "new_service/pkg/logger"
	"new_service/storage"

	"github.com/jmoiron/sqlx"
)


type PhoneService struct {
    storage storage.IStorage
    logger  l.Logger
}

func NewPhoneService(db *sqlx.DB, log l.Logger) *PhoneService {
    return &PhoneService{
        storage: storage.NewStoragePg(db),
        logger:  log,
    }
}

func (s *PhoneService) Create(ctx context.Context, req *pb.Phone) (*pb.Phone, error) {
	return s.storage.User().Create(req)
}

func (s *PhoneService) GetById(ctx context.Context, req *pb.GetByIdRequest) (*pb.Phone, error) {
	return s.storage.User().GetById(req)
}

func (s *PhoneService) GetByPhoneName(ctx context.Context, req *pb.GetByPhoneNameRequest) (*pb.Phone, error) {
	return s.storage.User().GetByPhoneName(req)
}
func (s *PhoneService) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.AllPhones, error) {
	return s.storage.User().GetAll(req)
}

func (s *PhoneService) Update(ctx context.Context, req *pb.Phone) (*pb.Phone, error) {
	return s.storage.User().Update(req)
}
func (s *PhoneService) Delete(ctx context.Context, req *pb.GetByIdRequest) (*pb.Phone, error) {
	return s.storage.User().Delete(req)
}

