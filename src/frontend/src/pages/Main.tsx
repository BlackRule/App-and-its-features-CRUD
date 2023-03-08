import React from 'react'
import {Link} from 'react-router-dom'

const Main = () => {
  // if (!localStorage.getItem('jwt_token')) {
  //   return <Redirect to="/auth/login"/>
  // }

  return (<>
    <Link to="/apps">Apps</Link>
    <Link to="/features">Features</Link>
  </>)
}

export default Main
