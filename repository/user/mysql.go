package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	domain "github.com/jwilyandi19/simple-auth/domain/user"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

func NewMySqlUserRepository(conn *sql.DB) domain.UserRepository {
	return &mysqlUserRepository{
		Conn: conn,
	}
}

func (m *mysqlUserRepository) FetchUser(ctx context.Context, req domain.FetchUserRequest) ([]domain.User, error) {
	offset := (req.Page - 1) * req.Limit
	rows, err := m.Conn.QueryContext(ctx, fetchUsers, offset, req.Limit)
	if err != nil {
		log.Println("[FetchUser-MySQL] QueryContext err", err.Error())
		return []domain.User{}, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Println("[Fetch-MySQL] rowsClose err", err.Error())
		}
	}()

	result := make([]domain.User, 0)
	for rows.Next() {
		t := domain.User{}
		err = rows.Scan(
			&t.ID,
			&t.Username,
			&t.FullName,
		)
		if err != nil {
			log.Println("[FetchUser-MySQL] ScanRow err", err.Error())
			return []domain.User{}, err
		}
		fmt.Println(t)
		result = append(result, t)
	}
	return result, nil
}

func (m *mysqlUserRepository) GetUser(ctx context.Context, username string) (domain.User, error) {
	rows, err := m.Conn.QueryContext(ctx, getUser, username)
	if err != nil {
		log.Println("[GetUser-MySQL] QueryContext err", err.Error())
		return domain.User{}, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Println("[GetUser-MySQL] rowsClose err", err.Error())
		}
	}()

	rows.Next()
	result := domain.User{}
	err = rows.Scan(
		&result.ID,
		&result.Username,
		&result.FullName,
		&result.HashedPassword,
	)
	if err != nil {
		log.Println("[GetUser-MySQL] ScanRow err", err.Error())
		return domain.User{}, err
	}

	return result, nil
}

func (m *mysqlUserRepository) CreateUser(ctx context.Context, req domain.CreateUserRequest) (domain.User, error) {
	stmt, err := m.Conn.PrepareContext(ctx, insertUser)
	if err != nil {
		log.Println("[CreateUser-MySQL] PrepareContext err", err.Error())
		return domain.User{}, err
	}

	res, err := stmt.ExecContext(ctx, req.Username, req.FullName, req.Password)
	if err != nil {
		log.Println("[CreateUser-MySQL] ExecContext err", err.Error())
		return domain.User{}, err
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		log.Println("[CreateUser-MySQL] LastInsertID err", err.Error())
		return domain.User{}, err
	}
	return domain.User{
		ID:       int(lastID),
		Username: req.Username,
		FullName: req.FullName,
	}, nil
}
