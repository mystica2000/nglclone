import { Footer } from "../components/Footer";
import Navbar from "../components/Navbar";
import { nav } from "../utils/constants";

export default function Login() {

  return (
    <>
    <Navbar displayFor={nav.INDEX}/>
    <main class="login-container">
    </main>
    <Footer />
    </>
  )
}