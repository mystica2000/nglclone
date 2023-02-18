import { Switch, Match } from "solid-js"
import { logoURL, nav } from "../utils/constants"
import "../styles/Navbar.css";
import LogoutIcon from "../assets/icons/logout";

export default function Navbar(prop: { displayFor: number }) {
  const loginURL = import.meta.env.VITE_API_URL + `/login`
  const logoutURL = import.meta.env.VITE_API_URL + `/logout`

  return (
    <nav class="nav">
      <img src="src/assets/icons/logo.png" alt="anon logo" width="250px" height="100px" class="img" />
      <Switch>
        <Match when={prop.displayFor === nav.INDEX}>
          <a href={loginURL} class="btn">login</a>
        </Match>
        <Match when={prop.displayFor === nav.HOME}>
          <div class="nav-home">
            <a href={logoutURL} class="btn">Logout <LogoutIcon /></a>
          </div>
        </Match>
      </Switch>
    </nav>
  )
}