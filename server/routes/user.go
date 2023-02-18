package routes

import (
	"encoding/json"
	"io"
	//"io"
	"net/http"
	"nglclone/controllers"
	interfaces "nglclone/interface"

	//	interfaces "nglclone/interface"
	"nglclone/utils"

	. "nglclone/logger"

	"go.uber.org/zap"
	"golang.org/x/net/context"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var cld *cloudinary.Cloudinary
var ctx context.Context

const MAX_UPLOAD_SIZE = 1 << 20

func init() {
	var err error
	cld, err = cloudinary.NewFromURL(V.GetString("CLOUDINARY_URL"))

	ctx = context.Background()

	if err != nil {
		Log.Error("Failed to intialize Cloudinary: ", zap.Error(err))
	}

	Log.Info("Got :", zap.String("api", V.GetString("CLOUDINARY_URL")))
}

func getUserInfo(w *http.ResponseWriter, id string) {

	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	(*w).Header().Set("Content-Type", "application/json")

	userInfo := controllers.FindById(id)

	if !utils.IsEmpty(userInfo) {
		(*w).WriteHeader(http.StatusOK)
		json.NewEncoder(*w).Encode(userInfo)
	} else {
		(*w).WriteHeader(http.StatusNotFound)
		(*w).Write([]byte(`{"error": "Not Found"}`))
	}
}

func updateUser(w *http.ResponseWriter, id string, user interfaces.UserForm) {
	if utils.IsProtectedName(user.Name) {
		(*w).WriteHeader(http.StatusOK)
		(*w).Write([]byte(`username not valid`))
		return
	}

	res, error := controllers.UpdateUser(id, user)
	if error != nil {
		http.Error(*w, error.Error(), http.StatusForbidden)
	} else {
		(*w).WriteHeader(http.StatusOK)
		json.NewEncoder(*w).Encode(res)
	}

}

func updateUserHelper(w *http.ResponseWriter, r *http.Request, k string, user interfaces.UserForm, id string) {
	file, fileHeader, err := r.FormFile(k)

	Log.Info("", zap.Int64("size", fileHeader.Size))

	if err != nil {
		Log.Error("Error :", zap.Error(err))
		return
	}

	// check before update that if it is already present using public id and user id!

	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{})

	if err != nil {
		Log.Error("Error: ", zap.Error(err))
		return
	}

	user.Image = result.SecureURL
	user.PublicID = result.PublicID

	Log.Info("result ", zap.Any("Result ", result.PublicID))
	defer file.Close()
	updateUser(w, id, user)
}

func deleteUserProfile(w *http.ResponseWriter, user_id string) bool {
	publicID, err := controllers.CheckPublicID(user_id)
	Log.Info("Test", zap.Any("e", publicID))
	if err != nil {
		Log.Error("Error :", zap.Error(err))
		(*w).WriteHeader(http.StatusNotFound)
		(*w).Write([]byte(`{"error": "Not Found"}`))
		return false
	}
	if !utils.IsEmpty(publicID) {
		res, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: publicID})

		Log.Info("Test", zap.Any("e", res))

		if err != nil {
			Log.Error("Error :", zap.Error(err))
			return false
		}
	}
	return true
}

func deleteUser(w *http.ResponseWriter, id string) {

	if deleteUserProfile(w, id) {
		_, error := controllers.DeleteUser(id)
		if error != nil {
			http.Error(*w, error.Error(), http.StatusForbidden)
		} else {
			(*w).WriteHeader(http.StatusOK)
			(*w).Write([]byte(`Deleted`))
		}
	} else {
		return
	}

}

func removeUserImage(w *http.ResponseWriter, id string) {

	if deleteUserProfile(w, id) {
		Log.Info("Profile pic deleted!")
		_, error := controllers.UpdateUserImage(id)
		if error != nil {
			http.Error(*w, error.Error(), http.StatusForbidden)
		} else {
			(*w).WriteHeader(http.StatusOK)
			(*w).Write([]byte(`deleted`))
		}
	} else {
		(*w).WriteHeader(http.StatusInternalServerError)
		(*w).Write([]byte(`tryagain`))
		return
	}
}

func User(w http.ResponseWriter, r *http.Request) {

	id := r.Context().Value("id").(string)

	Log.Info("ID", zap.String("id", id))

	if r.Method == "GET" {
		getUserInfo(&w, id)
		return

	} else if r.Method == "PUT" {

		r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
		err := r.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if err != nil {
			http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
			return
		}

		if err != nil {
			Log.Error("Error :", zap.Error(err))

			(w).WriteHeader(http.StatusNotFound)
			(w).Write([]byte(`{"error": "Not Found"}`))
			return
		}

		Log.Info("Name : ", zap.Any("data", r.Form.Get("name")))

		var user interfaces.UserForm

		user.Name = r.Form.Get("name")

		multiForm := r.MultipartForm

		for k, _ := range multiForm.File {

			if deleteUserProfile(&w, id) {
				updateUserHelper(&w, r, k, user, id)
			}
			return
		}

		user.Image = ""
		user.PublicID = ""
		updateUser(&w, id, user)
		return

	} else if r.Method == "DELETE" {

		body, _ := io.ReadAll(r.Body)
		if string(body) == "image" {
			removeUserImage(&w, id)
		} else {
			deleteUser(&w, id)
		}

		return
	}
}
