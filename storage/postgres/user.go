package postgres

import (
	pb "new_service/genproto"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cast"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(phone *pb.Phone) (*pb.Phone, error) {
	query := `
		INSERT INTO phone (id, phone_name, color, price, ram, memory, created_at) 
		VALUES ($1, $2, $3, $4, $5, $6, $7) 
		RETURNING id, phone_name, color, price, ram, memory, created_at
	`

	err := r.db.QueryRow(query, phone.Id, phone.PhoneName, phone.Color, phone.Price, phone.Ram,phone.Memory, time.Now()).Scan(
		&phone.Id,
		&phone.PhoneName,
		&phone.Color,
		&phone.Price,
		&phone.Ram,
		&phone.Memory,
		&phone.CreatedUp,
	)

	if err != nil {
		return nil, err
	}
	return phone, nil
}

func (r *userRepo) GetById(req *pb.GetByIdRequest) (*pb.Phone, error) {

	query := "SELECT id, phone_name, color, price, ram, memory, created_at FROM phone WHERE id = $1"

	row := r.db.QueryRow(query, req.Id)

	phone := &pb.Phone{}
	err := row.Scan(
		&phone.Id,
		&phone.PhoneName,
		&phone.Color,
		&phone.Price,
		&phone.Ram,
		&phone.Memory,
		&phone.CreatedUp,
	)
	if err != nil {
		return nil, err
	}

	return phone, nil
}

func (r *userRepo) GetByPhoneName(req *pb.GetByPhoneNameRequest) (*pb.Phone, error) {

	query := "SELECT id, phone_name, color, price, ram, memory, created_at FROM phone WHERE phone_name = $1"

	row := r.db.QueryRow(query, req.PhoneName)

	phone := &pb.Phone{}
	err := row.Scan(
		&phone.Id,
		&phone.PhoneName,
		&phone.Color,
		&phone.Price,
		&phone.Ram,
		&phone.Memory,
		&phone.CreatedUp,
	)
	if err != nil {
		return nil, err
	}

	return phone, nil
}

func (r *userRepo) GetAll(req *pb.GetAllRequest) (*pb.AllPhones, error) {
    intLimit := cast.ToInt(req.Limit)
    intPage := cast.ToInt(req.Page)
    offset := (intPage - 1) * intLimit

    query := `
        SELECT id, phone_name, color, price, ram, memory, created_at, updated_at
        FROM phone
        LIMIT $1 OFFSET $2
    `

    rows, err := r.db.Query(query, intLimit, offset)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    phones := []*pb.Phone{}

    for rows.Next() {
        phone := &pb.Phone{}
        err := rows.Scan(
            &phone.Id,
            &phone.PhoneName,
            &phone.Color,
            &phone.Price,
            &phone.Ram,
            &phone.Memory,
            &phone.CreatedUp,
			&phone.UpdatedAt,
        )
        if err != nil {
            return nil, err
        }
        phones = append(phones, phone)
    }

    if err = rows.Err(); err != nil {
        return nil, err
    }

    return &pb.AllPhones{Phone: phones}, nil
}

func (r *userRepo) Update(req *pb.Phone) (*pb.Phone, error) {
	query := `
        UPDATE phone 
        SET phone_name = $2, color = $3, price = $4, ram = $5, memory = $6, created_at = $7, updated_at = $8
        WHERE id = $1
        RETURNING id, phone_name, color, price, ram, memory, created_at, updated_at
    `

	row := r.db.QueryRow(query, req.Id, req.PhoneName, req.Color, req.Price, req.Ram, req.Memory, time.Now(), time.Now())

	updatedPhone := &pb.Phone{}
	err := row.Scan(&updatedPhone.Id, &updatedPhone.PhoneName, &updatedPhone.Color, &updatedPhone.Price, &updatedPhone.Ram, &updatedPhone.Memory,&updatedPhone.CreatedUp, &updatedPhone.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return updatedPhone, nil
}

func (r *userRepo) Delete(req *pb.GetByIdRequest) (*pb.Phone, error) {

	query := "DELETE FROM phone WHERE id = $1 RETURNING id, phone_name, color, price, ram, memory, created_at, updated_at"

	row := r.db.QueryRow(query, req.Id)

	deletedPhone := &pb.Phone{}
	err := row.Scan(&deletedPhone.Id, &deletedPhone.PhoneName, &deletedPhone.Color, &deletedPhone.Price, &deletedPhone.Ram, &deletedPhone.Memory, &deletedPhone.CreatedUp, &deletedPhone.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return deletedPhone, nil
}
