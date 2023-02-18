import { Link } from "@solidjs/router";
import Navbar from "../components/Navbar";

export default function Error() {

  return (
    <main>
      <Navbar displayFor={-1} />
      <div class="error-container">
        <h1>
          404
        </h1>
        <h2>
          Page Not Found!
        </h2>
        Oops, The page you are looking for might have deleted or unavailable at the moment!
        <Link class="btn btn-blue" href="/">Go to Home Page</Link>
      </div>
    </main>
  )
}