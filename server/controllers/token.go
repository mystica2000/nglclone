package controllers

import (
	"database/sql"

	"github.com/google/uuid"
	"go.uber.org/zap"
	. "nglclone/logger"
)

func CreateToken() string {
	uuid := uuid.NewString();

	return uuid;
}

func InsertToken(user_id string) (string,error) {

	stmt,err := Db.Prepare("INSERT OR REPLACE INTO token(token,user_id) VALUES(?,?)")

	if err != nil {
		Log.Error("Error: ",zap.Error(err))
	}

	token := CreateToken()
  res , err := stmt.Exec(token,user_id)

	if err != nil {
		Log.Error("Error: ",zap.Error(err))
	}

	if row,err := res.RowsAffected(); err!=nil {
		return "",err;
	} else if row > 0 {
		return token,nil;
	} else {
		return "",err;
	}
}

func DeleteToken(user_id string) bool {

	stmt,_ := Db.Prepare("DELETE FROM token WHERE user_id = ?")

  res , err := stmt.Exec(user_id)
	if err != nil {
		Log.Error("Error: ",zap.Error(err))
		return false;
	}

	n, _ := res.RowsAffected()

	if( n > 0) {
		return true;
	} else {
		return false;
	}
}


func FindToken(token string) (string,error) {

	row := Db.QueryRow("SELECT user_id FROM token WHERE token = ?",token)

	var id string
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			Log.Error("Error :",zap.Error(err))
		} else {
			Log.Error("Error :",zap.Error(err))
		}
		return "",err;
	}

	return id,nil;
}

func DeleteTokenByCRON() {
	stmt,err := Db.Prepare("DELETE FROM token;")

	if err != nil {
		Log.Error("Error: ",zap.Error(err))
		return;
	}

  res , err := stmt.Exec()

	if err != nil {
		Log.Error("Error: ",zap.Error(err))
	}

  r,_ := res.RowsAffected();
	Log.Info("Rows affected", zap.Int64("records",r));
}