import axios, { AxiosResponse } from "axios"
import store from "../store/store"
import { User } from "../utils/types"

export const getUserData = async() => {

  return new Promise<User>(async (resolve,reject)=> {

  const {setUser} = store

  await axios.get<User>("/user", {withCredentials: true})
  .then(function (response: AxiosResponse) {
    // handle success
    if (response.status == 200) {
      setUser(response.data);
      resolve(response.data);
    }
  })
  .catch(function (error: unknown) {
    // handle error
    if(axios.isAxiosError(error)) {
      console.log(error.response)
      if (error.response) {
        if(error.response.status == 423) {
          window.location.href = "/expired"
        }
      } else if (error.request) {
        reject()
      }
    } else {
      reject()
    }
  })

  })
}