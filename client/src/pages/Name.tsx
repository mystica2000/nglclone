import { createEffect, createSignal, Show } from "solid-js"
import EditImage from "../components/EditImage";
import { Footer } from "../components/Footer";
import Navbar from "../components/Navbar";
import { getUserData } from "../helper/getUserData";
import { updateUser } from "../helper/postUserData";
import store from "../store/store";
/**
 * this page will be displayed as long as user name is empty
 * After login for the *first time* (mostly),
 *   - get the name from the user before home page.
 *   - if they only changed the profile picture, again it will be displayed
*/
export default function Name() {

  const { user } = store;

  createEffect(async () => {
    await getUserData().catch(err => { console.log(err) });
  }, [])

  const [name, setName] = createSignal('');
  const [status, setStatus] = createSignal(false);
  const [errString, setErrString] = createSignal("");

  const handleUserChange = (e: any) => {
    const newValue = e.target.value.replace(/\s/g, '');
    setName(newValue);
  }

  const handlePost = async () => {
    setName(name().trim());
    if (name().length < 2) {
      setStatus(true);
      setErrString("Username is Required");
      return;
    }
    try {
      const updated = await updateUser(name(), undefined); // update the name only, not image here!
      if (updated) {
        window.location.href = "http://localhost:3000/home"
      }
    } catch (e) {
      console.error(e)
    }
  }

  return (
    <>
      <Navbar displayFor={-1} />

      <Show when={user().email.length > 0} fallback={<div class="loader">Loading...</div>}>

        <main class="name-container">
          <div class=" name-container-2">
            <h1>Just one more step...😉</h1>
            <EditImage />

            <div style={"display:grid;grid-template-rows:1fr;gap:1em;"}>
              <label for="name" style={"text-align:left;"}>
                Enter your Username:
              </label>
              <input type="text" id="name" value={name()} onKeyUp={handleUserChange} class="input-text" />
              <Show when={status()}>
                <p class="error-msg" style={"text-align:center;"}>{errString()}</p>
              </Show>
              <button class="btn btn-success profile-input" onClick={handlePost}>Next</button>
            </div>
          </div>
        </main>

      </Show>

      <Footer />
    </>
  )
}