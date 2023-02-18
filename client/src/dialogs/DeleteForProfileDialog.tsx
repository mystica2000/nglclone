import { createEffect, createSignal, JSX, onMount, Setter } from "solid-js";
import "../styles/Dialog.css"
import CloseIcon from "../assets/icons/close";
import { deleteProfile } from "../helper/deleteProfile";
import store from "../store/store";
import AfterDeleteDialog from "./AfterDeleteDialog";

export default function DeleteForProfileDialog(props: { setShowDialog: Setter<boolean>}) {

  const [text, setText] = createSignal<string>("");
  const {user} = store;
  const [showError,setShowError] = createSignal(false);
  const [deleted,setDeleted] = createSignal(false);

  let dialog: any;

  let blur = `
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(255, 255, 255, 0.5);
  backdrop-filter: blur(7px);
  `

  onMount(() => {
    dialog.show();
  });

  const handleCancel = () => {
    props.setShowDialog(false);
    dialog.close();
  }

  const handleChangeText = (e:any) => {
    setShowError(false);
    setText(e.target.value);
  }

  const handleDelete = async () => {
    if (text() === user().name) {
      try{
      setDeleted(await deleteProfile(false)); // not only image, whole!
        // info dialog for deleted user and then on ok, redirect to home page!
      } catch(e) {
        // error! something went wrong..
      }
    } else {
      setShowError(true);
    }
  }

  return (
    <>
      <div style={blur}></div>
      <dialog style={"min-width:250px;width:100%;max-width:500px;"} id="dialog1" ref={dialog}>
        <header style={"display:grid;grid-template-columns:1fr 0fr;background-color:#8fc6ff;padding:4px;"}>
          <p style={"margin-top:3px;font-size:20px;"}>Delete Profile</p>
          <button style={"border:transparent;background:transparent;"} onClick={handleCancel}><CloseIcon /></button>
        </header>
        <div style={"padding:15px;font-size:15px;"}>
          <div>
            <label>Confirm Deleting Profile by, Type your username:
              <input type="text" value={text()} onInput={handleChangeText}/>
            </label>
            {showError() && <div class="error-msg">Invalid Username</div>}
          </div>
          <div class="space-between fill">
            <button class="btn btn-success" onClick={handleDelete}>Yes, Delete Profile</button>
            <button class="btn btn-danger" onClick={handleCancel}>Cancel</button>
          </div>
        </div>
      </dialog>
      { deleted() && <AfterDeleteDialog />}
    </>
  )
}