import axios, { AxiosResponse } from "axios"
import { createSignal, JSX } from "solid-js"
import store from "../store/store";
import "../styles/Toggle.css"

export default function ToggleButton(props:{active:number}) {

  const { user } = store
  const [active,setActive] = createSignal<boolean>(props.active ? true : false);

  // Code for update active state!
  const handleToggleChange = async () => {
    return new Promise(()=> {

    setActive(active() == false ? true : false);
    axios.post<number>("/link/toggle",
     user()?.name,
     { withCredentials: true })
    .then(function (response: AxiosResponse) {
      // handle success
      if (response.status == 200) {
        setActive(response.data.active);
      }
    })
    .catch(function (error: unknown) {
      // handle error
      if (axios.isAxiosError(error)) {
        if (error.response) {
          if(error.response.status == 423) {
            window.location.href = "/expired"
          } else {
          setActive(active() == false ? true : false) // active()
          }
        }
      } else {
        console.log(error)
      }
    })
    })
  }

  return(
  <>
    <label class="toggle-button edit">
    <input type="checkbox" id="toggle-checkbox" checked={active()} onChange={handleToggleChange} />
    <div class="toggle-button-slider"></div>
    </label>
  </>
  )
}