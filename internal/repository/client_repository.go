package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"simple-grpc-crud/protobuf/pb"
)

type ClientStorage struct {
	db *pgx.Conn
}

func NewClientStorage(db *pgx.Conn) *ClientStorage {
	return &ClientStorage{db: db}
}

func (c ClientStorage) Create(ctx context.Context, in *pb.CreateClientRequest) (*pb.CreateClientResponse, error) {
	sql, args, _ := sq.Insert("clients").Values(in.Id, in.Name, in.Surname, in.Lastname, in.Age, in.Weight, in.Height).ToSql()

	_, err := c.db.Exec(ctx, sql, args...)
	if err != nil {
		return &pb.CreateClientResponse{}, err
	}

	return &pb.CreateClientResponse{Id: in.Id}, nil
}

func (c ClientStorage) Read(ctx context.Context, in *pb.ReadClientRequest) (*pb.ReadClientResponse, error) {
	sql, args, _ := sq.Select("*").From("clients").Where(sq.Eq{"id": in.Id}).ToSql()

	var client pb.CreateClientRequest
	err := pgxscan.Select(ctx, c.db, &client, sql, args...)
	if err != nil {
		return &pb.ReadClientResponse{}, nil
	}

	return &pb.ReadClientResponse{Client: &client}, nil
}

func (c ClientStorage) Update(ctx context.Context, in *pb.UpdateClientRequest) (*pb.UpdateClientResponse, error) {
	sql, args, _ := sq.Update("clients").
		Set("id", in.Client.Id).
		Set("name", in.Client.Name).
		Set("surname", in.Client.Surname).
		Set("lastname", in.Client.Lastname).
		Set("age", in.Client.Age).
		Set("height", in.Client.Height).
		Set("weight", in.Client.Weight).
		ToSql()

	_, err := c.db.Exec(ctx, sql, args...)
	if err != nil {
		return &pb.UpdateClientResponse{}, nil
	}

	return &pb.UpdateClientResponse{Id: in.Client.Id}, nil

}

func (c ClientStorage) Delete(ctx context.Context, in *pb.DeleteClientRequest) (*pb.DeleteClientResponse, error) {
	sql, args, _ := sq.Delete("*").From("clients").Where(sq.Eq{"id": in.Id}).ToSql()

	_, err := c.db.Exec(ctx, sql, args...)
	if err != nil {
		return &pb.DeleteClientResponse{}, nil
	}

	return &pb.DeleteClientResponse{Id: in.Id}, nil

}
