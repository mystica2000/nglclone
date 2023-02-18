import { createEffect, createSignal, Show } from "solid-js"
import { JSX } from "solid-js";
import { Link, useLocation } from "@solidjs/router";
import "../styles/Login.css"
import axios, { AxiosResponse } from "axios";
import "../styles/Home.css"

import "../styles/Responses.css"
import Navbar from "../components/Navbar";
import Error from "./Error";
import { Footer } from "../components/Footer";
import { defaultProfilePicURL } from "../utils/constants";

export default function Form() {

  const [text, setText] = createSignal('');
  const location = useLocation();
  const [link, setLink] = createSignal<string>('');
  const [valid, setValid] = createSignal(false);
  const [src, setSrc] = createSignal('');

  createEffect(() => {
    let link: string = location.pathname.substring(1, location.pathname.length); // link name (username)
    setLink(link);
    /** checks whether the link is valid or not,if not fallback to error */
    axios.get<string>("/link", { params: { link: link } })
      .then(function (response: AxiosResponse) {
        if (response.status == 200) {
          setValid(true);
          setSrc(response.data || defaultProfilePicURL);
        }
      })
      .catch(function (error: unknown) {
        if (axios.isAxiosError(error)) {
          if (error.response) {
            setValid(false);
          }
        } else {
          console.log(error)
        }
      })

  }, [])

  const handleOnChange: JSX.EventHandler<HTMLTextAreaElement, InputEvent> = (e) => {
    setText(e?.currentTarget.value)
  }

  const handlePost = () => {
    if (text()) {
      axios.post("/link", JSON.stringify({ link: location.pathname.substring(1, location.pathname.length), question: text() }), { withCredentials: false })
        .then(function (response: AxiosResponse) {
          // handle success
          if (response.status == 200) {
            if (response.data) {
              localStorage.setItem("url", window.location.href); // store the current url in localstorage so that we can use it in redirect page!
              window.location.href = "/redirect"; // after successful post, redirect page!
            }
          }
        })
        .catch(function (error: unknown) {
          // handle error
          if (axios.isAxiosError(error)) {
            if (error.response) {
              window.location.href = "/"
            }
          }
          console.log(error)
        });
    }
  }

  return (
    <>
      <Show when={valid()} fallback={<Error />}>
        <Navbar displayFor={-1} />
        <main class="name-container">
          <div class=" name-container-2">
            <div class="header">
              <img src={src()} alt="profile picture" class="small-img" />
              <h2>{link()}</h2>
            </div>
            <div class="sub-header">Send me anonymous messages! </div>
            <div class="main">
              <textarea onInput={handleOnChange} cols="40" rows="5" class="textarea"></textarea></div>
            <div class="footer"><button class="btn btn-success" onClick={handlePost}>Send</button><button class="btn btn-danger">Cancel</button></div>
          </div>
        </main>
        <Footer />
      </Show>
    </>
  )
}