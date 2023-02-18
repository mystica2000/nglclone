import { createEffect, createSignal, JSX, onMount, Setter } from "solid-js";
import "../styles/Dialog.css"
import CloseIcon from "../assets/icons/close";
import { deleteQuestion } from "../helper/deleteQuestion";
import { updateUser } from "../helper/postUserData";
import { deleteProfile } from "../helper/deleteProfile";
import store from "../store/store";

export default function DeleteForResponseDialog(props: { title:string, okText:string, setShowDialog: Setter<boolean>, id: number | string,dontshow:boolean }) {

  const [showAgain, setShowAgain] = createSignal<number>(0);
  const [error,setError] = createSignal('');
  const {user} = store;

  let dialog: any;

  let blur = `
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(255, 255, 255, 0.5);
  backdrop-filter: blur(7px);
  z-index:1;
  `

  onMount(() => {
    dialog.show();
  });

  const handleCancel = () => {
    console.log("is it owrking!")
    props.setShowDialog(false);
    dialog.close();
  }

  const handleDelete = async () => {
    try{
    console.log(props.okText)
    if(props.okText === "Yes, Delete") { //TODO: Can be improved!
      await deleteQuestion(props.id);
    } else {
      console.log("henlo")
      if(user().image.length>0) {
        console.log("henlo #2")
        await deleteProfile(true);
      }
      console.log(user())
    }

    } catch(err) {
    setError('Something went wrong,Try again later!')
    }
    handleCancel();
  }

  const showAgainEvent = () => {
    let temp = showAgain() == 0 ? 1 : 0
    setShowAgain(temp);

    localStorage.setItem("dontshow", temp.toString());
  }

  return (
    <>
      <div style={blur}></div>
      <dialog style={"min-width:250px;width:100%;max-width:500px;z-index:2;"} id="dialog" ref={dialog}>
        <header style={"display:grid;grid-template-columns:1fr 0fr;background-color:#8fc6ff;padding:4px;"}>
          <p style={"margin-top:3px;"}>{props.title}</p>
          <button style={"border:transparent;background:transparent;"} onClick={handleCancel}><CloseIcon /></button>
        </header>
        <div style={"padding:15px;"}>
          {props.dontshow && <label style={"display:inline-block;"}><input type="checkbox" name="checkbox" value={showAgain()} onInput={showAgainEvent} /> Don't show again</label>}
          <div class="space-between fill">
            <button class="btn btn-success" onClick={handleDelete}>{props.okText}</button>
            <button class="btn btn-danger" onClick={handleCancel}>Cancel</button>
          </div>
        </div>
      </dialog>
    </>
  )
}