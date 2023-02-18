package routes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"nglclone/controllers"
	interfaces "nglclone/interface"
	"time"

	. "nglclone/logger"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleInfo struct {
	Id      string `json:"sub"`
	Picture string `json:"picture"`
	Email   string `json:"email"`
	Verfied bool   `json:"email_verified"`
}

var V *viper.Viper
var Clients []interfaces.Clients

func init() {

	if Log == nil {
		InitLogger()
	}

	V = viper.New()
	V.SetConfigName("env")
	V.SetConfigType("json")
	V.AddConfigPath(".")

	err := V.ReadInConfig()

	if err != nil {
		Log.Error("Error reading Config file", zap.Error(err))
	} else {
		Log.Info("Viper Config File Read!", zap.Any("viper", V.Get("clientID")))
	}
}

func GoogleLogin(w http.ResponseWriter, r *http.Request) {

	Log.Info("/login")

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	conf := &oauth2.Config{
		ClientID:    V.Get("clientID").(string),
		RedirectURL: "http://localhost:8080/login/callback", // URL it sends the results to, (ie) "Authorization_code"
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	// Redirect user to Google's consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("new")
	Log.Info("Visit the URL for the Auth:", zap.String("url", url))

	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}

func GoogleCallBack(w http.ResponseWriter, r *http.Request) {
	// change it to frontend url when deploying
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the auth token from the query
	code := r.URL.Query()["code"][0]

	conf := &oauth2.Config{
		ClientID:     V.Get("clientID").(string),
		ClientSecret: V.Get("clientSecret").(string),
		RedirectURL:  "http://localhost:8080/login/callback", // URL it sends the results to, (ie) "Authorization_code"
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	// // Handle the exchange code to initiate a transport.

	// Send POST request containing "authorization_code" to the token Endpoint to get the access token..
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}

	// tok contains access token
	url1 := "https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + tok.AccessToken
	client := conf.Client(oauth2.NoContext, tok)
	resp, _ := client.Get(url1)

	str, _ := ioutil.ReadAll(resp.Body)

	// Prints Email Address since we asked for the email address in the scope before :)
	Log.Info("user", zap.String("profile", string(str)))

	user := new(GoogleInfo)

	err = json.Unmarshal(str, &user)
	if err != nil {
		Log.Error("Error: ", zap.Error(err))
	}

	defer resp.Body.Close()

	// if user is verified
	if user.Verfied {

		temp := interfaces.User{Id: user.Id, Name: "", Email: user.Email}

		// get the id and name from the email if present
		found, isNamePresent := controllers.Find(temp.Email)

		/**
		- Find By Id
		- Create User
		**/

		// if id is present
		if len(found) > 0 {

			// create cookie for the user
			createCookie(w, found)

			// if name is present then direct to home page
			// else, redirect to names page to get the username
			if isNamePresent {
				http.Redirect(w, r, "http://localhost:3000/home", http.StatusPermanentRedirect)
				return
			} else {
				http.Redirect(w, r, "http://localhost:3000/name", http.StatusPermanentRedirect)
				return
			}

		} else {

			// if id is not present, create a new user!
			rows, _ := controllers.CreateUser(temp)

			if rows > 0 {
				createCookie(w, temp.Id)
				http.Redirect(w, r, "http://localhost:3000/name", http.StatusPermanentRedirect)
				return
			} else {
				http.Redirect(w, r, "http://localhost:3000", http.StatusInternalServerError)
				return
			}

		}
	} else {
		http.Redirect(w, r, "http://localhost:3000", http.StatusInternalServerError)
		return
	}
}

func GoogleLogout(w http.ResponseWriter, r *http.Request) {

	Log.Info("/logout")

	id := r.Context().Value("id").(string)

	Log.Info("ID", zap.String("id", id))

	// delete the session token from the DB!
	if controllers.DeleteToken(id) {
		Log.Info("logout successful")

		cookie := http.Cookie{
			Name:   "session",
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}

		http.SetCookie(w, &cookie)
		// send logout message

		http.Redirect(w, r, "http://localhost:3000/", http.StatusTemporaryRedirect)
		return
	} else {
		Log.Error("Server Error when logging out, try again")
		// send error if cannot able to delete the token off the DB
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
}

func createCookie(w http.ResponseWriter, id string) {

	expiration := time.Now().Add(time.Minute * 4)

	token, err := controllers.InsertToken(id)

	if err != nil {
		Log.Error("Unable to create a token")
		return
	}

	cookie := http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  expiration,
	}

	http.SetCookie(w, &cookie)
}
