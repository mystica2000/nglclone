import axios, { AxiosResponse } from "axios"
import store from "../store/store"

export const deleteProfile = async (imageOnly:boolean): Promise<boolean> => {

  return new Promise(async (resolve, reject) => {

    const {user,setUser} = store;


    await axios.delete("/user", { data:(imageOnly == true ? "image":""), withCredentials: true })
      .then(function (res: AxiosResponse) {
        // handle success
        if (res.status == 200) {
          let temp: any = user()
          temp.image = "";
          setUser((state) => { return { ...state, ...temp } })


          resolve(true);
        } else {
          resolve(false);
        }
      })
      .catch(function (error: unknown) {
        // handle error
        if (axios.isAxiosError(error)) {
          console.log(error.response)
          if (error.response) {
            if(error.response.status == 423) {
              window.location.href = "/expired"
            } else {
              reject('something went wrong! try again later.');
            }
          } else if (error.request) {
            reject('network ')
          }
        } else {
          reject()
        }
      })
  })
}