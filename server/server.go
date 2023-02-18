package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"

	"nglclone/controllers"
	. "nglclone/logger"
	"nglclone/routes"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func cronJob(wg *sync.WaitGroup) {

	Log.Info("Cron Job Started")

	c := cron.New()
	c.AddFunc("@every 10m", func() {
		Log.Info("Cron Deleting the Session Data of the Users")
		controllers.DeleteTokenByCRON()
	})
	c.Start()
}

func startHTTPServer(wg *sync.WaitGroup) {

	// Define a middleware function to add headers to the response
	authMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			(w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
			(w).Header().Set("Access-Control-Allow-Credentials", "true")
			(w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			(w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			w.Header().Set("Content-Type", "application/json")

			if r.Method == "OPTIONS" {
				w.WriteHeader(204)
				return
			}

			cookie, err := r.Cookie("session")
			if err != nil {
				switch {
				case errors.Is(err, http.ErrNoCookie):
					http.Error(w, "cookie not found", http.StatusLocked)
				default:
					log.Println(err)
					http.Error(w, "server error", http.StatusInternalServerError)
				}
				return
			}

			// check if cookie is present in the DB
			id, err := controllers.FindToken(cookie.Value)

			if err != nil {
				Log.Info("Error ", zap.Error(err))
			}

			if id != "" {
				ctx := context.WithValue(r.Context(), "id", id)
				r = r.WithContext(ctx)
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "session expired", http.StatusLocked)
			}
		})
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/login", routes.GoogleLogin)
	mux.HandleFunc("/login/callback", routes.GoogleCallBack)
	mux.Handle("/logout", authMiddleware(http.HandlerFunc(routes.GoogleLogout)))

	mux.Handle("/user", authMiddleware(http.HandlerFunc(routes.User)))
	mux.Handle("/response", authMiddleware(http.HandlerFunc(routes.Responses)))
	mux.Handle("/link/toggle", authMiddleware(http.HandlerFunc(routes.ToggleURL)))

	mux.HandleFunc("/link", routes.Link)

	/**
	User - add,update(name),delete,read
	Link - add,update(active,edition)
	**/

	Log.Info("Server running on :: http://localhost:8080")
	if oops := http.ListenAndServe(":8080", mux); oops != nil {
		log.Fatal(oops)
	}
}

func main() {

	InitLogger()
	var wg sync.WaitGroup
	wg.Add(2)

	go cronJob(&wg)
	go startHTTPServer(&wg)

	wg.Wait()
}
