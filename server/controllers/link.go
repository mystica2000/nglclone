package controllers

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"

	. "nglclone/logger"
)

func CreateURL(id string) bool {

	name, _ := FindUserById(id)

	stmt, err := Db.Prepare("INSERT OR REPLACE INTO link(url,active,user_id) VALUES(?,?,?)")

	if err != nil {
		Log.Error("Error : ", zap.Error(err))
	}

	result, err := stmt.Exec(name, 1, id)

	if err != nil {
		Log.Error("Error : ", zap.Error(err))
	}

	if row, err := result.RowsAffected(); err != nil {
		Log.Error("Error : ", zap.Error(err))
	} else if row > 0 {
		Log.Info("URL Created for the user : ", zap.String("id", id))
		return true
	}

	return false
}

func FindIfValidLink(linkName string) (int, int64, string) {

	row := Db.QueryRow("SELECT link.id,link.active,users.image FROM link JOIN users on users.id = link.user_id where url = ?;", linkName)

	var id int64
	var active int
	var image string
	if err := row.Scan(&id, &active, &image); err != nil {
		if err == sql.ErrNoRows {
			return -1, -1, ""
		} else {
			Log.Error("Error", zap.Error(err))
			return -1, -1, ""
		}
	}

	return active, id, image
}

// toggle url!
func ToggleURL(linkName string) int {
	fmt.Print(linkName)
	found, id, _ := FindIfValidLink(linkName)

	if found > -1 && id > -1 {
		if found == 0 {
			found = 1
		} else {
			found = 0
		}

		stmt, _ := Db.Prepare("UPDATE link SET active = ? where id = ? ")

		res, err := stmt.Exec(found, id)

		if err != nil {
			Log.Error("Error: ", zap.Error(err))
		}

		r, _ := res.RowsAffected()

		if r > 0 {
			return found
		} else {
			return -1
		}
	} else {
		return -1
	}
}

func FindIfLinkPresent(id string) (bool, error) {
	row := Db.QueryRow("SELECT active FROM link where user_id = ?", id)

	var active int
	if err := row.Scan(&active); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		} else {
			Log.Error("Error", zap.Error(err))
			return false, fmt.Errorf("database error")
		}
	}

	return true, nil
}

func UserIdByLink(id int64) (string, error) {

	row := Db.QueryRow("Select user_id from link where id = ?", id)

	var user_id string
	if err := row.Scan(&user_id); err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		} else {
			Log.Info("Error ", zap.Error(err))
		}
		return "", fmt.Errorf("database Error")
	}
	return user_id, nil
}
