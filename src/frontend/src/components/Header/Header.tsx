import React from 'react'
import {Link} from 'react-router-dom'

const logout = () => {
  localStorage.removeItem('jwt_token')
  window.location.reload()
}

export function Header() {

  return <header>
    {
      (localStorage.getItem('jwt_token'))?
        <button
          onClick={logout}
          className=""
        >Logout</button>
        :
        <Link to = "/auth/login">Login</Link>
    }
  </header>
}

