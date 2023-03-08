import React, {useState} from 'react'
import {Link, useNavigate} from 'react-router-dom'

import {useMutation} from '@apollo/client'
import {REGISTER} from '../graphql/queries'


const Register = () => {
  const navigate = useNavigate()
  const [register, {data, loading, error}] = useMutation(REGISTER)
  const [submitted, setSubmitted] = useState(false)
  if (submitted) {
    if (loading) return <p>Submitting...</p>
    if (error) return <p>Submission error! {error.message}</p>
    else {
      localStorage.setItem('jwt_token', data?.register?.token ?? '')
      navigate('/')
    }
  }
  if (localStorage.getItem('jwt_token') !== null) navigate('/')
  const onSubmit: React.FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault()
    if (!(event.target instanceof HTMLFormElement)) return
    const [email, password] = event.target.elements
    if (!(email instanceof HTMLInputElement) || !(password instanceof HTMLInputElement)) return
    register({
      variables: {
        email: email.value,
        password: password.value,
      }
    })
    setSubmitted(true)
  }

  return <form onSubmit={onSubmit} className="form-signin">
    <h1 className="">
      Please register a new account
    </h1>
    <label htmlFor="inputEmail" className="">
      Email address
    </label>
    <input
      type="email"
      id="inputEmail"
      className=""
      placeholder="Email address"
      required
    />
    <label htmlFor="inputPassword" className="">
      Password
    </label>
    <input
      type="password"
      id="inputPassword"
      className=""
      placeholder="Password"
      required
    />
    <button className="" type="submit">
      Create Account
    </button>
    <p className="">
      Already registered ? <Link to="/auth/login">Logn in here</Link>
    </p>
  </form>
}

export default Register
