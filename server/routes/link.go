package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nglclone/controllers"
	interfaces "nglclone/interface"
	. "nglclone/logger"
	"nglclone/utils"

	"go.uber.org/zap"
)

type FormResponse struct {
	Link string `json:"link"`
	Text string `json:"question"`
	Id   int64  `json:"id"`
}

func addHeader(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
}

// no security! public api!
// for the valid link!
func validLink(w *http.ResponseWriter, linkName string) {
	addHeader(w)

	active, _, url := controllers.FindIfValidLink(linkName)

	if active > 0 {
		(*w).WriteHeader(http.StatusOK)
		fmt.Fprint((*w), url)
	} else {
		(*w).WriteHeader(http.StatusNotFound)
		fmt.Fprint((*w), "invalid url")
	}
}

// protected route
func ToggleURL(w http.ResponseWriter, r *http.Request) {

	link, _ := io.ReadAll(r.Body)
	linkName := string(link)

	if len(linkName) > 0 {

		active := controllers.ToggleURL(linkName)

		var urlResp interfaces.ToggleURLResponse
		urlResp.Active = active

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(urlResp)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "URL Not Found")
	}
}

/*
 Update response from the form data
*/
func insertResponse(w *http.ResponseWriter, data FormResponse) {

	// if url is not empty,
	if len(data.Link) > 0 {

		active, id, _ := controllers.FindIfValidLink(data.Link)
		if active > 0 && id > 0 {

			Log.Info("Insert Form Response")
			// yes, insert the response ... Hmm check if the response contains unsolicated words
			// if yes, reply like Ayo it's not allowed here
			// if no, insert into DB
			_, err := controllers.InsertResponse(id, data.Text)

			if err != nil {
				http.Error((*w), "Required Missing Data", http.StatusUnprocessableEntity)
				return
			}

			test, _ := controllers.UserIdByLink(id)
			fmt.Print("test ", test)
			// go SendResponsesToClient(user_id)

			(*w).WriteHeader(http.StatusOK)
			fmt.Fprint((*w), "Success")
			return

		} else {
			http.Error((*w), "Invalid URL", http.StatusNotFound)
			// return no active error
		}
	} else {
		http.Error((*w), "URL Not Found", http.StatusNotFound)
	}
}

func Link(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS,POST")

	if r.Method == "GET" {

		params := r.URL.Query()

		link := params.Get("link")
		linkName := string(link)

		if len(linkName) > 0 {

			validLink(&w, linkName)
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "URL Not Found")
		}

	} else if r.Method == "POST" {
		// id or all

		body, _ := io.ReadAll(r.Body)
		var data FormResponse
		err := json.Unmarshal(body, &data) // map data into struct
		if err != nil {
			Log.Error("Error", zap.Error((err)))
		}

		if !utils.IsEmpty(data) {
			insertResponse(&w, data)
			return
		}
	}

}
