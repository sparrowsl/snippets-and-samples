import { redirect, useNavigate, useParams } from "@solidjs/router"

export default function RedirectURL() {
  const params = useParams()
  async function redirectTo() {
    const res = await fetch(`http://localhost:5000/urls/${params.id}`)

    if (!res.ok) {
      window.location.pathname = "/"
    }

    const { url } = await res.json()
    window.location.replace(url.long_url)
  }
  redirectTo()

  return <></>
}
