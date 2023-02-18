import { Setter } from "solid-js"
import { updateUser } from "../helper/postUserData";
import store from "../store/store";
import "../styles/Responses.css"

export default function ShowConfirmationBtn(props: { class: string, setMode: Setter<string>, prevName: string }) {

  const { user, setUser } = store

  const handleClick = async () => {
    if (user().name.length > 1) {
      // add setError()
      const updated = await updateUser(user()?.name).catch(err => {})
      if(updated) {
        props.setMode("label");
      }

    }
  }

  const handleCancel = () => {
    let temp = user();
    temp.name = props.prevName;
    setUser((prev) => { return { ...prev, ...temp } })
    props.setMode("label");
  }

  return (
    <div class={props.class}>
      <button class="btn btn-success" onClick={handleClick}>Save</button>
      <button class="btn btn-warning" onClick={handleCancel}>Cancel</button>
    </div>)
}