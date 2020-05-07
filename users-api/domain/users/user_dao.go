package users

import (
	"database/sql"
	"github.com/Nicrii/Project/users-api/datasources/users_db"
	"github.com/Nicrii/Project/users-api/logger"
	"github.com/Nicrii/Project/users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email) VALUES($1, $2, $3) RETURNING id;"
	queryGetUser    = "SELECT * FROM users WHERE id=$1;"
	queryUpdateUser = "UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4;"
	queryDeleteUser = "DELETE FROM users WHERE id=$1;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Db.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	getResult := stmt.QueryRow(user.Id)

	err = getResult.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email)
	if err == sql.ErrNoRows {
		logger.Error("user not found", err)
		return errors.NewNotFoundError("user not found")
	}
	if err != nil {
		logger.Error("error when trying to get user by id", err)
		return errors.NewInternalServerError("database error")

	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Db.Prepare(queryInsertUser)
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

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Db.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.NewInternalServerError("database error")

	}
	count, err := res.RowsAffected()
	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.NewInternalServerError("database error")

	}
	if count == 0 {
		logger.Error("user not found", err)
		return errors.NewNotFoundError("user_not_found")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Db.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Id)
	if err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.NewInternalServerError("database error")

	}
	count, err := res.RowsAffected()
	if count == 0 {
		logger.Error("user not found", err)
		return errors.NewNotFoundError("user_not_found")
	}
	return nil
}
