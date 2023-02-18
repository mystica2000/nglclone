import axios, { AxiosResponse } from "axios"
import { UserResponse } from "../utils/types"

export const getResponses = async (): Promise<UserResponse[]> => {
  return new Promise((resolve, reject) => {

    axios.get<UserResponse[]>("/response", { withCredentials: true })
      .then(function (response: AxiosResponse) {
        // handle success
        if (response.status == 200) {
          resolve(response.data)
        }
      })
      .catch(function (error: unknown) {
        // handle error
        if (axios.isAxiosError(error)) {
          if (error.response) {
            if(error.response.status == 423) {
              window.location.href = "/expired"
            }
          } else if (error.request) {
            reject('error in request!')
          }
        } else {
          console.log(error)
        }
      })

  })
}