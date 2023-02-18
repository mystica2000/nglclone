import { createSignal, Show } from 'solid-js';
import "../styles/Home.css"
import EditInputIcon from '../assets/icons/editInput';
import ShowConfirmationBtn from './ShowConfirmationBtn';
import ToggleButton from './ToggleButton';
import CopyToClipBoard from './CopyToClipBoard';
import store from '../store/store';
import EditImage from './EditImage';
import DeleteForProfileDialog from '../dialogs/DeleteForProfileDialog';

export default function Profile() {

  const { user, setUser } = store;

  const [mode, setMode] = createSignal<string>("label");
  const [showDeleteDialog, setShowDeleteDialog] = createSignal<boolean>();

  let prevName: string = user().name;

  const handleUserChange = (e: any) => {

    const newValue = e.target.value.replace(/\s/g, '');

    let temp: any = user()
    temp.name = newValue;
    setUser((state) => { return { ...state, ...temp } })
  }

  const handleDeleteProfile = () => {
    setShowDeleteDialog(true);
  }

  return (
    <section class="profile">
      <Show when={user().email.length > 0} fallback={<div class="loader">Loading...</div>}>
        <EditImage />

        <div class="form">
          <div class="form-row">

            <div class="edit space-between" tabIndex={0}
              onKeyPress={(e: KeyboardEvent) => { if (e.key === "Enter") { setMode(mode() === "label" ? "input" : "label"); } }}
              onClick={() => { setMode(mode() === "label" ? "input" : "label"); }}>
              <label> Username: </label>
              <EditInputIcon />
            </div>

            <Show when={mode() === "input"} fallback={<label>{user()?.name}</label>}>
              <input type="text" name="username" id="" value={user()?.name} onInput={handleUserChange} class='profile-input' />
              <ShowConfirmationBtn class="space-between fill" setMode={setMode} prevName={prevName} />
            </Show>

          </div>

          <label for="email" class="form-row">
            Email: <br />
            {user()?.email}
          </label>

          <div class="form-row space-between">
            <label style={"margin-top: 5px;"}> Active: </label>
            <ToggleButton active={user()?.active || 0} />
          </div>


          <div class="form-row">

            <div class="edit space-between">
              <label> Link: </label>
              <CopyToClipBoard link={window.location.origin + "/" + user()?.name} />
            </div>
            <div style={"margin-top:5px;"}>
              {window.location.origin + "/" + user()?.name}
            </div>
          </div>

          <div class='form-row'>
            <button class="btn btn-danger" style={"margin:auto;display:flex;"} onClick={handleDeleteProfile}>Delete Profile</button>
          </div>

          {showDeleteDialog() && <DeleteForProfileDialog setShowDialog={setShowDeleteDialog} />}

          <br />
          <br />
        </div>
      </Show>
    </section>
  )
}