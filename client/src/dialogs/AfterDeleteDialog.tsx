import { createEffect, createSignal, JSX, onMount, Setter } from "solid-js";
import "../styles/Dialog.css"
import CloseIcon from "../assets/icons/close";

export default function AfterDeleteDialog() {

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

  const handleDelete = async () => {
    window.location.href = "/"
  }

  return (
    <>
      <div style={blur}></div>
      <dialog style={"min-width:250px;width:100%;max-width:500px;"} id="dialog" ref={dialog}>
        <header style={"display:grid;grid-template-columns:1fr 0fr;background-color:#8fc6ff;padding:4px;"}>
          <p style={"margin-top:3px;"}>Your Account has been deleted!</p>
          <button style={"border:transparent;background:transparent;"} onClick={handleDelete}><CloseIcon /></button>
        </header>
        <div style={"padding:15px;"}>
          <div class="space-between fill">
            <button class="btn btn-success" onClick={handleDelete}>Ok</button>
          </div>
        </div>
      </dialog>
    </>
  )
}