import { Footer } from "../components/Footer";
import Navbar from "../components/Navbar";

export default function Redirect() {

  //  After submit, send another message?
  //  Or Sign up :

  const handleRedirect = () => {
    let url = localStorage.getItem("url"); // contains previous url
    window.location.href = url ? url : "http://localhost:3000"; // if no previous url, redirect to home page instead
  }

  const handleHomePage = () => {
    window.location.href = "http://localhost:3000" // sign up!
  }

  return (
    <>
      {/* display nav bar with only logo */}
      <Navbar displayFor={-1} />

      <main class="name-container">
        <div class=" name-container-2">
          <h1>Your Response has been submitted 🎉</h1>
          <div style={"display:grid;grid-template-rows:1fr;gap:1em;justify-content: center;margin-top:3px;"}>
            <button class="btn btn-blue" onClick={handleRedirect}>Send another!</button>
            <button class="btn btn-success" onClick={handleHomePage}>Sign up</button>
          </div>
        </div>
      </main>

      <Footer />
    </>
  )
}