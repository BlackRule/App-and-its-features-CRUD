import React from 'react'
import {Redirect} from 'react-router-dom'

const Dashboard = () => {
  const logout = () => {
    localStorage.removeItem('jwt_token')
    window.location.reload()
  }

  if (!localStorage.getItem('jwt_token')) {
    return <Redirect to="/auth/login"/>
  }

  return (
    <header>
      <button
        onClick={() => logout()}
        className=""
      >
        Logout
      </button>
    </header>
  )
}

export default Dashboard
