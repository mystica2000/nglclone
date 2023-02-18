import axios, { AxiosResponse } from "axios"
import store from "../store/store"
import { User, UserResponse } from "../utils/types"

export const deleteQuestion = async (id: number | string) => {
  const { responses, setResponses } = store

  return new Promise<User>(async (resolve, reject) => {

    await axios.delete("/response", { data: id, withCredentials: true })
      .then(function (res: AxiosResponse) {
        // handle success
        let updatedResponses;
        if (res.status == 200) {

          if (id == "all") {
            updatedResponses = [] as UserResponse[];
          } else if(id == "replied") {
            updatedResponses = responses.filter((responses) => responses.done != 1)
          } else {
            updatedResponses = responses.filter((response) => response.id !== id);
          }

          setResponses(updatedResponses);
          resolve(res.data);
        }
      })
      .catch(function (error: unknown) {
        // handle error
        if (axios.isAxiosError(error)) {
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