package controllers

import (
	"fmt"
	. "nglclone/logger"
	"nglclone/utils"

	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/zap"

	interfaces "nglclone/interface"
)

func InsertResponse(link_id int64, data string) (int64, error) {

	if len(data) > 0 {
		stmt, err := Db.Prepare("INSERT INTO response(questions,link_id) VALUES(?,?)")

		if err != nil {
			Log.Error("Error: ", zap.Error(err))
		}

		result, err := stmt.Exec(data, link_id)

		if err != nil {
			Log.Error("Error : ", zap.Error(err))
			return 0, fmt.Errorf("database error")
		}

		row, _ := result.RowsAffected()

		return row, nil
	} else {
		return 0, fmt.Errorf("missing required data")
	}
}

func GetQuestionStateAndUpdate(id int64) (int64, error) {
	if id > 0 {
		row := Db.QueryRow("SELECT * from response where id = ?", id)

		var id int64
		var questions string
		var link string
		var done int64
		if err := row.Scan(&id, &questions, &link, &done); err != nil {
			return -1, fmt.Errorf("database error %w", err)
		}

		if done == 1 {
			done = 0
		} else {
			done = 1
		}

		stmt, _ := Db.Prepare("UPDATE response SET done = ? where id = ? ")

		res, err := stmt.Exec(done, id)

		if err != nil {
			Log.Error("Error: ", zap.Error(err))
			return -1, fmt.Errorf("database error")
		}

		r, _ := res.RowsAffected()
		Log.Info("Rows affected", zap.Int64("records", r))

		return done, nil
	} else {
		return -1, fmt.Errorf("invalid id")
	}
}

func GetResponses(user_id string) []interfaces.QuestionResponse {

	rows, err := Db.Query("SELECT response.id,questions,done from response join link on response.link_id = link.id where link.user_id = ?", user_id)
	if err != nil {
		Log.Info("Error ", zap.Error(err))
	}
	defer rows.Close()

	var questions []interfaces.QuestionResponse

	for rows.Next() {

		var question interfaces.QuestionResponse
		var qn string
		var done int
		var id int64

		err = rows.Scan(&id, &qn, &done)
		if err != nil {
			Log.Info("Error ", zap.Error(err))
		}
		question.Done = done
		question.Response = qn
		question.Id = id
		questions = append(questions, question)
	}

	err = rows.Err()
	if err != nil {
		Log.Info("Error ", zap.Error(err))
	}

	return questions
}

func DeleteResponse(id int64) (int64, error) {
	if id > 0 {

		stmt, _ := Db.Prepare("DELETE FROM response where id = ? ")

		res, err := stmt.Exec(id)

		if err != nil {
			Log.Error("Error: ", zap.Error(err))
			return -1, fmt.Errorf("database error")
		}

		r, _ := res.RowsAffected()
		Log.Info("Rows affected", zap.Int64("records", r))

		return r, nil
	} else {
		return -1, fmt.Errorf("invalid response")
	}
}

func DeleteResponseByUserID(user_id string) (int64, error) {
	if !utils.IsEmpty(user_id) {

		stmt, _ := Db.Prepare("DELETE FROM response WHERE response.link_id IN (SELECT id from link where link.user_id = ?);")

		res, err := stmt.Exec(user_id)

		if err != nil {
			Log.Error("Error: ", zap.Error(err))
			return -1, fmt.Errorf("database error")
		}

		r, _ := res.RowsAffected()
		Log.Info("Rows affected", zap.Int64("records", r))

		return r, nil
	} else {
		return -1, fmt.Errorf("invalid response")
	}
}

func DeleteRepliedResponseByUserID(user_id string) (int64, error) {
	if !utils.IsEmpty(user_id) {

		stmt, _ := Db.Prepare("DELETE FROM response WHERE response.link_id IN (SELECT id from link where link.user_id = ?) and response.done = 1;")

		res, err := stmt.Exec(user_id)

		if err != nil {
			Log.Error("Error: ", zap.Error(err))
			return -1, fmt.Errorf("database error")
		}

		r, _ := res.RowsAffected()
		Log.Info("Rows affected", zap.Int64("records", r))

		return r, nil
	} else {
		return -1, fmt.Errorf("invalid response")
	}
}
