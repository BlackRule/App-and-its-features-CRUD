import React from 'react'
import {Link, Redirect} from 'react-router-dom'
import {Header} from '../components/Header/Header'

const Dashboard = () => {
  // if (!localStorage.getItem('jwt_token')) {
  //   return <Redirect to="/auth/login"/>
  // }

  return (<>
    <Header/>
    <Link to="/apps">Apps</Link>
    <Link to="/features">Features</Link>
  </>)
}

export default Dashboard
