package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"nglclone/controllers"
	"nglclone/utils"
	"strconv"

	// "nglclone/utils"
	. "nglclone/logger"

	"go.uber.org/zap"
)

func getAllResponsesByUser(w *http.ResponseWriter, id string) {

	responses := controllers.GetResponses(id)

	jsonData, err := json.Marshal(responses)
	if err != nil {
		http.Error((*w), err.Error(), http.StatusInternalServerError)
		return
	}

	(*w).WriteHeader(http.StatusOK)
	(*w).Write(jsonData)
}

func updateResponse(w *http.ResponseWriter, response_id int64) {
	// if url is not empty,
	if response_id > 0 {

		done, err := controllers.GetQuestionStateAndUpdate(response_id)

		if err != nil {
			http.Error((*w), err.Error(), http.StatusUnprocessableEntity)
			return
		}

		(*w).WriteHeader(http.StatusOK)
		fmt.Fprint((*w), done)

	} else {
		(*w).WriteHeader(http.StatusNotFound)
		fmt.Fprint((*w), "Question Not Found")
	}
}

func deleteResponse(w *http.ResponseWriter, id int64) {
	if id > 0 {

		_, err := controllers.DeleteResponse(id)

		if err != nil {
			(*w).WriteHeader(http.StatusBadRequest)
			fmt.Fprint((*w), "Try again later")
		}

		(*w).WriteHeader(http.StatusOK)
		fmt.Fprint((*w), "Deleted the Response")

	} else {
		(*w).WriteHeader(http.StatusNotFound)
		fmt.Fprint((*w), "URL Not Found")
	}
}

func deleteAllResponseByUserId(w *http.ResponseWriter, user_id string) {
	if !utils.IsEmpty(user_id) {

		_, err := controllers.DeleteResponseByUserID(user_id)

		if err != nil {
			(*w).WriteHeader(http.StatusBadRequest)
			fmt.Fprint((*w), "Try again later")
		}

		(*w).WriteHeader(http.StatusOK)
		fmt.Fprint((*w), "deleted all responses")

	} else {
		(*w).WriteHeader(http.StatusNotFound)
		fmt.Fprint((*w), "URL Not Found")
	}
}

func deleteAllRepliedResponseByUserId(w *http.ResponseWriter, user_id string) {
	if !utils.IsEmpty(user_id) {

		_, err := controllers.DeleteRepliedResponseByUserID(user_id)

		if err != nil {
			(*w).WriteHeader(http.StatusBadRequest)
			fmt.Fprint((*w), "Try again later")
		}

		(*w).WriteHeader(http.StatusOK)
		fmt.Fprint((*w), "deleted all responses")

	} else {
		(*w).WriteHeader(http.StatusNotFound)
		fmt.Fprint((*w), "URL Not Found")
	}
}

func Responses(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)

	Log.Info("ID", zap.Any("id", id))

	if r.Method == "GET" {

		getAllResponsesByUser(&w, id)
		return

	} else if r.Method == "PUT" {
		// id or all

		body, _ := io.ReadAll(r.Body)
		response_id, _ := strconv.ParseInt(string(body), 0, 64)

		if !utils.IsEmpty(response_id) {
			updateResponse(&w, response_id)
			return
		}

	} else if r.Method == "DELETE" {
		// id or all!
		body, _ := io.ReadAll(r.Body)
		text := string(body)

		if text == "all" {
			deleteAllResponseByUserId(&w, id)
			return
		} else if text == "replied" {
			deleteAllRepliedResponseByUserId(&w, id)
			return
		} else {

			response_id, _ := strconv.ParseInt(text, 0, 64)

			if !utils.IsEmpty(response_id) {
				deleteResponse(&w, response_id)
				return
			}

		}
	}

	(w).WriteHeader(http.StatusNotFound)
	(w).Write([]byte(`{"error": "Not Found"}`))
}

// func ResponsesSSE(w http.ResponseWriter, r *http.Request) {

// 	fmt.Print("test")

// 	id := r.Context().Value("id").(string)
// 	flusher, ok := w.(http.Flusher)

// 	if !ok {
// 		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "text/event-stream")
// 	w.Header().Set("Cache-Control", "no-cache")
// 	w.Header().Set("Connection", "keep-alive")

// 	var client *interfaces.Clients
// 	for i, c := range Clients {
// 		if c.ID == id {
// 			client = &Clients[i]
// 			break
// 		}
// 	}

// 	if client == nil {
// 		client = &interfaces.Clients{
// 			ID:      id,
// 			Channel: make(chan []interfaces.QuestionResponse),
// 		}
// 		Clients = append(Clients, *client)
// 	}

// 	go func() {
// 		for {
// 			select {
// 			case data := <-client.Channel:
// 				var buf bytes.Buffer
// 				enc := json.NewEncoder(&buf)
// 				enc.Encode(data)
// 				fmt.Printf("test %v", data)
// 				fmt.Fprintf(w, "data: hi\n\n")
// 				flusher.Flush()
// 			}
// 		}
// 	}()

// }

// func SendResponsesToClient(clientID string) {

// 	responses := controllers.GetResponses(clientID)

// 	for _, client := range Clients {
// 		if client.ID == clientID {
// 			client.Channel <- responses
// 			break
// 		}
// 	}
// }

// // when they login, create a new clientID for them

// // responses is event stream (GET)
// //    flush to client

// // insert responses
// //   call the client ID to send the responses (along new!)

// // to delete the channel id,
// //   logout, when session expired db call!
