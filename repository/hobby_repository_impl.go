package repository

import (
	"context"
	"database/sql"
	"net/http"

	cError "github.com/vincen320/user-service-graphql/helper/error"
	"github.com/vincen320/user-service-graphql/model"
)

type hobbyRepository struct {
	db *sql.DB
}

func NewHobbyRepository(db *sql.DB) HobbyRepository {
	return &hobbyRepository{
		db: db,
	}
}

func (h *hobbyRepository) FindHobbies(ctx context.Context) (response []model.Hobby, err error) {
	rows, err := h.db.Query("SELECT id, name FROM hobbies")
	if err == sql.ErrNoRows {
		err = nil
	}
	if err != nil {
		return nil, cError.New(http.StatusInternalServerError, "internal server error", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var hobby model.Hobby
		if err = rows.Scan(
			&hobby.ID,
			&hobby.Name,
		); err != nil {
			return nil, cError.New(http.StatusInternalServerError, "internal server error", err.Error())
		}
		response = append(response, hobby)
	}
	return
}
