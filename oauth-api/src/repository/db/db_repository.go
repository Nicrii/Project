package db

import (
	"github.com/Nicrii/Project/oauth-api/src/clients/postgresql"
	"github.com/Nicrii/Project/oauth-api/src/domain/access_token"
	"github.com/Nicrii/Project/oauth-api/src/logger"
	"github.com/Nicrii/Project/oauth-api/src/utils/errors"
	"github.com/Nicrii/Project/users-api/datasources/users_db"
)

const (
	queryGetAccessToken    = "SELECT access_token, user_id, client_id, expires FROM access_tokens WHERE access_token=$1;"
	queryCreateAccessToken = "INSERT INTO access_tokens(access_token, user_id, client_id, expires) VALUES ($1, $2, $3, $4);"
	queryUpdateExpires     = "UPDATE access_tokens SET expires=$1 WHERE access_token=$2;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}

func (r *dbRepository) Create(id string) *errors.RestErr {
	stmt, err := users_db.Db.Prepare(queryCreateAccessToken)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("database error")

	}
	defer stmt.Close()
	insertResult := stmt.QueryRow(user.FirstName, user.LastName, user.Email)
	err = insertResult.Scan(&user.Id)
	if err != nil {
		logger.Error("error when trying to save user", saveErr)
		return errors.NewInternalServerError("database error")

	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, errors.NewInternalServerError("database connection not implemented yet")
}
