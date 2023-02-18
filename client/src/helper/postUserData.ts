import axios, { AxiosResponse } from "axios";
import store from "../store/store";

export const updateUser = (name?: string, imageFile?: File) => {

  const {user,setUser} = store;

  return new Promise((resolve,reject)=> {
    const formData = new FormData() // multi-part form since we might send the image too
    let flag = false; // for the form data
    if (name) {
      formData.append("name", name);
      flag = true
    }

    if (imageFile) {
      formData.append("image", imageFile);
      flag = true;
    }

    if(flag) {

      axios.put("/user", formData, {  withCredentials: true,  headers: { "Content-Type": "multipart/form-data" }, }).then((response: AxiosResponse) => {
        if (response.status == 200) {
          if(formData.get("image")) {
            // updae the image
            let temp: any = user()
            temp.image = response.data.image;
            setUser((state) => { return { ...state, ...temp } })
          }
          resolve(true);
        }
      }).catch((err: unknown) => {
        if (axios.isAxiosError(err)) { // handle error
          if (err.response) {
            if(err.response.status == 423) {
              window.location.href = "/expired"
            } else if(err.response.status == 400) {
              alert(err.response.data);
              reject();
            }
          } else {
            console.log(err);
          }
        }
      });

    } else {
      resolve(true);
    }
  })
}