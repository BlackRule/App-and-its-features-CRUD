import React, {useState} from 'react'
import {Link, useNavigate} from 'react-router-dom'
import {useMutation} from '@apollo/client'
import {LOGIN} from '../graphql/queries'


const Login = () => {
  const navigate = useNavigate()
  const [login, {data, loading, error}] = useMutation(LOGIN)
  const [submitted, setSubmitted] = useState(false)

  if (submitted) {
    if (loading) return <p>Submitting...</p>
    if (error) return <p>Submission error! {error.message}</p>
    else {
      localStorage.setItem('jwt_token', data?.login?.token ?? '')
      navigate('/')
    }
  }
  const onSubmit: React.FormEventHandler<HTMLFormElement> = (event) => {
    event.preventDefault()
    if (!(event.target instanceof HTMLFormElement)) return
    const [email, password] = event.target.elements
    if (!(email instanceof HTMLInputElement) || !(password instanceof HTMLInputElement)) return
    login({
      variables: {
        email: email.value,
        password: password.value,
      }
    })
    setSubmitted(true)
  }
  return (
    <div className="">
      <form className="" onSubmit={onSubmit}>
        <h1 className="">Please sign in</h1>
        <label htmlFor="inputEmail" className="">
          Email address
        </label>
        <input
          type="email"
          id="inputEmail"
          className=""
          placeholder="Email address"
          required
          autoFocus
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
          Sign in
        </button>
        <p className="">
          No account ? <Link to="/auth/register">Create one here</Link>
        </p>
      </form>
    </div>
  )
}

export default Login
