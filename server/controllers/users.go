package controllers

import (
	"database/sql"

	"fmt"
	interfaces "nglclone/interface"
	. "nglclone/logger"
	"nglclone/utils"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

func CreateUser(user interfaces.User) (int64, error) {

	stmt, err := Db.Prepare("INSERT INTO users(id,name,image,email) VALUES(?,?,?,?)")

	if err != nil {
		Log.Info("Error", zap.Error(err))
	}

	result, err := stmt.Exec(user.Id, "", user.Image, user.Email)

	if err != nil {
		Log.Info("Error", zap.Error(err))
	}

	rowsAffected, err := result.RowsAffected()

	return rowsAffected, err
}

func Find(email string) (string, bool) {

	row := Db.QueryRow("SELECT id,name FROM users WHERE email = ?", email)

	var id string
	var name = ""
	if err := row.Scan(&id, &name); err != nil {
		if err == sql.ErrNoRows {
			return "", false
		} else {
			Log.Info("Error ", zap.Error(err))
		}
		return "", false
	}

	found := false
	if len(name) > 0 {
		found = true
	}
	Log.Info("TEST", zap.Any("id", id))
	return id, found
}

func FindById(id string) interfaces.UserResponse {

	Log.Info("test id find BY ID", zap.String("test id", id))

	isPresent, err := FindIfLinkPresent(id)

	var userInfo interfaces.UserResponse

	if err != nil {
		Log.Error("Error :", zap.Error(err))
		return userInfo
	}

	sqlString := ""
	if isPresent {
		sqlString = "SELECT email,name,active,image FROM users join link on users.id = link.user_id WHERE users.id = ?"
	} else {
		sqlString = "SELECT email,name,image FROM users where id = ?"
		Log.Info("test @2")
	}

	row := Db.QueryRow(sqlString, id)

	if isPresent {

		err = row.Scan(&userInfo.Email, &userInfo.Name, &userInfo.Active, &userInfo.Image)
	} else {
		err = row.Scan(&userInfo.Email, &userInfo.Name, &userInfo.Image)
	}

	Log.Info("test ", zap.Any("testing", userInfo.Email))

	if err != nil {
		return interfaces.UserResponse{}
	}

	return userInfo
}

func UpdateUser(id string, user interfaces.UserForm) (interfaces.UserResponse, error) {

	Log.Info("TEST", zap.Any("id", id), zap.Any("name", user.Image))

	var names string
	if err := Db.QueryRow("SELECT name FROM users where name = ? and id <> ?", user.Name, id).Scan(&names); err != nil {
		if err == sql.ErrNoRows && !utils.IsProtectedName(user.Name) {
			Log.Info("LKET GOOO")

			sqlString := `UPDATE users SET name = CASE WHEN COALESCE(?,"") = "" THEN name else ? END,image = CASE WHEN COALESCE(?,"") = "" THEN image ELSE ? end,public_id = CASE WHEN coalesce(?,"")="" then public_id ELSE ? end where id = ?`
			stmt, err := Db.Prepare(sqlString)

			if err != nil {
				Log.Error("Error: ", zap.Error(err))
			}

			result, err := stmt.Exec(user.Name, user.Name, user.Image, user.Image, user.PublicID, user.PublicID, id)

			if err != nil {
				Log.Error("Error: ", zap.Error(err))
			}

			row, err := result.RowsAffected()

			if err != nil {
				Log.Error("Error: ", zap.Error(err))
			}

			if len(user.Name) > 0 {
				if CreateURL(id) {
					return interfaces.UserResponse{Name: user.Name}, nil
				} else {
					return interfaces.UserResponse{}, fmt.Errorf("url creation failed")
				}
			}

			if row > 0 {
				return interfaces.UserResponse{Image: user.Image}, nil
			}

		}
	}

	return interfaces.UserResponse{}, fmt.Errorf("user already present")
}

func UpdateUserImage(id string) (int64, error) {

	Log.Info("TEST", zap.Any("id", id))

	sqlString := `UPDATE users SET image = '',public_id = '' where id = ?` //TODO: NOT WORKING!
	stmt, err := Db.Prepare(sqlString)

	if err != nil {
		Log.Error("Error: ", zap.Error(err))
	}

	result, err := stmt.Exec(id)

	if err != nil {
		Log.Error("Error: ", zap.Error(err))
	}

	rows, err := result.RowsAffected()

	if err != nil {
		Log.Error("Error: ", zap.Error(err))
	}

	return rows, nil
}

func FindUserById(id string) (string, error) {

	var names string
	if err := Db.QueryRow("SELECT name FROM users where id = ?", id).Scan(&names); err != nil {
		if err == sql.ErrNoRows {
			return "invalid id", sql.ErrNoRows
		}
	}
	return names, nil
}

func DeleteUser(id string) (int64, error) {

	stmt, err := Db.Prepare("DELETE from users where id = ?")

	if err != nil {
		Log.Error("Error: ", zap.Error(err))
	}

	result, err := stmt.Exec(id)

	if err != nil {
		Log.Error("Error: ", zap.Error(err))
	}

	rows, err := result.RowsAffected()

	if err != nil {
		Log.Error("Error: ", zap.Error(err))
		return 0, fmt.Errorf("try again later")
	}

	return rows, nil
}

func CheckPublicID(id string) (string, error) {

	row := Db.QueryRow("Select public_id from users where id = ?", id)

	var public_id string
	if err := row.Scan(&public_id); err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		} else {
			Log.Info("Error ", zap.Error(err))
		}
		return "", fmt.Errorf("database Error")
	}
	return public_id, nil
}
