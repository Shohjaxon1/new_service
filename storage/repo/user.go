package repo

import (
	pb "new_service/genproto"
)


// UserStorageI ...
type UserStorageI interface {
	Create(*pb.Phone) (*pb.Phone, error)
	GetById(*pb.GetByIdRequest) (*pb.Phone, error)
	GetByPhoneName(*pb.GetByPhoneNameRequest) (*pb.Phone, error)
	GetAll(*pb.GetAllRequest) (*pb.AllPhones, error)
	Update(*pb.Phone) (*pb.Phone, error)
	Delete(*pb.GetByIdRequest) (*pb.Phone, error)
}