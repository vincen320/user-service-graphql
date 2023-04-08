package repository

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	cError "github.com/vincen320/user-service-graphql/helper/error"
	"github.com/vincen320/user-service-graphql/model"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) FindUsers(ctx context.Context) (response []model.User, err error) {
	rows, err := u.db.Query(
		`SELECT
			id
			, name
			, age
			, address
			, salary
		FROM users`,
	)
	if err == sql.ErrNoRows {
		err = nil
	}
	if err != nil {
		return nil, cError.New(http.StatusInternalServerError, "internal server error", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var user model.User
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Age,
			&user.Address,
			&user.Salary,
		); err != nil {
			return nil, cError.New(http.StatusInternalServerError, "internal server error", err.Error())
		}
		response = append(response, user)
	}
	return
}

func (u *userRepository) FindUserByID(ctx context.Context, userID int64) (response []model.User, err error) {
	var user model.User
	if err = u.db.QueryRow(
		`SELECT
			id
			, name
			, age
			, address
			, salary
		FROM users
		WHERE id = $1`,
		userID,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Age,
		&user.Address,
		&user.Salary,
	); err != nil {
		return response, cError.New(http.StatusInternalServerError, "internal server error", err.Error())
	}
	response = append(response, user)
	return
}

func (u *userRepository) CreateUser(ctx context.Context, request model.User) (user model.User, err error) {
	err = u.db.QueryRow(
		`INSERT INTO users(
			name
			, age
			, address
			, salary
		)VALUES($1, $2, $3, $4) RETURNING id`,
		request.Name,
		request.Age,
		request.Address,
		request.Salary,
	).Scan(&request.ID)
	if err != nil {
		return user, cError.New(http.StatusInternalServerError, "internal server error", err.Error())
	}
	user = request
	return
}

func (u *userRepository) FindUserHobbies(ctx context.Context, userIDs []int64) (response map[int64][]model.Hobby, err error) {
	placeholders := make([]string, len(userIDs))
	params := make([]any, len(userIDs))
	for i, userID := range userIDs {
		params[i] = userID
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}
	rows, err := u.db.Query(
		fmt.Sprintf(
			`SELECT
			uh.user_id
			, uh.hobby_id
			, h.name
		FROM user_hobbies uh
		JOIN hobbies h ON h.id = uh.hobby_id
		WHERE uh.user_id IN(%s)`,
			strings.Join(placeholders, ",")),
		params...)
	if err == sql.ErrNoRows {
		err = nil
	}
	if err != nil {
		return nil, cError.New(http.StatusInternalServerError, "internal server error", err.Error())
	}
	defer rows.Close()
	response = map[int64][]model.Hobby{}
	for rows.Next() {
		var userID int64
		var hobby model.Hobby
		if err = rows.Scan(
			&userID,
			&hobby.ID,
			&hobby.Name,
		); err != nil {
			return nil, cError.New(http.StatusInternalServerError, "internal server error", err.Error())
		}
		response[userID] = append(response[userID], hobby)
	}
	return
}
