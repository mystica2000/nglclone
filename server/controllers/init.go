package controllers

import (
	"database/sql"

	. "nglclone/logger"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"
)

var Db *sql.DB;

func init() {

	var err error;
	Db, err = sql.Open("sqlite3","./ngl.db")

	if(Log==nil) {
		InitLogger()
	}

	if err != nil {
		Log.Error("Error Database", zap.Error(err))
	} else {
		Log.Info("Database started")
		Db.Exec("PRAGMA foreign_keys=ON");
	}
}