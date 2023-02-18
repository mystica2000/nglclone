import { Link } from "@solidjs/router"

export default function SessionExpired() {
  const loginURL = import.meta.env.VITE_API_URL + `/login`

  return (
    <div class="error-container">
      Oops, your session expired. please login again to continue.
      <Link href={loginURL}>
        login again!
      </Link>
    </div>
  )
}