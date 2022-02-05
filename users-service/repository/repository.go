package repository

import (
	"context"
	"github.com/BinaryArchaism/users-service/users-service/models"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

type IRepository interface {
	AddUser(ctx context.Context, user *models.UserToAdd) (int, error)
	DeleteUser(ctx context.Context, id *models.UserId) error
	GetUsers(ctx context.Context) (*models.Users, error)
}

type repository struct {
	db *pgx.Conn
}

func (r *repository) AddUser(ctx context.Context, user *models.UserToAdd) (int, error) {
	var id int
	if err := r.db.QueryRow(ctx, "INSERT INTO users (first_name, last_name, email, age) values ($1, $2, $3, $4) RETURNING id",
		user.FirstName, user.LastName, user.Email, user.Age).Scan(&id); err != nil {
		logrus.Debug(err)
		return id, err
	}
	return id, nil
}

func (r *repository) DeleteUser(ctx context.Context, id *models.UserId) error {
	if err := r.db.QueryRow(ctx, "delete from users where id=$1",
		id).Scan(&id); err != nil {
		logrus.Debug(err)
		return err
	}
	return nil
}

func (r repository) GetUsers(ctx context.Context) (*models.Users, error) {
	var users models.Users
	rows, err := r.db.Query(ctx, "select * from users")
	if err != nil {
		logrus.Debug(err)
		return nil, err
	}
	var curUser models.FullUser
	for rows.Next() {
		err := rows.Scan(&curUser.Id, &curUser.FirstName, &curUser.LastName, &curUser.Email, &curUser.Age)
		if err != nil {
			logrus.Debug(err)
			return nil, err
		}
		users.Users = append(users.Users, &curUser)
	}

	return &users, nil
}

func NewRepository(db *pgx.Conn) IRepository {
	return &repository{db: db}
}
