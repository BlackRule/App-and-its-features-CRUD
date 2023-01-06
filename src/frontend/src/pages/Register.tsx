import React, {useState} from 'react'
import {Link, Redirect, useHistory} from 'react-router-dom'

import {useMutation} from "@apollo/client"
import { gql } from '../__generated__/gql'

const REGISTER = gql(`
    mutation register(
        $email: String!
        $password: String!
    ) {
        register(
            input: { email: $email, password: $password }
        ) {
            token
            user {
                id
                email
            }
        }
    }
`)


const Register = () => {
    const history = useHistory()
    const [register, {data, loading, error}] = useMutation(REGISTER);
    const [submitted,setSubmitted]=useState(false)
    if(submitted) {
        if (loading) return <p>Submitting...</p>
        if (error) return <p>Submission error! {error.message}</p>
        else {
            localStorage.setItem("jwt_token", data?.register?.token??"")
            history.push('/')
        }
    }
    if (localStorage.getItem("jwt_token") !== null) return <Redirect to="/"/>
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
        });
        setSubmitted(true)
    }

    return <form onSubmit={onSubmit} className="form-signin">
        <h1 className="h3 mb-3 font-weight-normal text-center">
            Please register a new account
        </h1>
        <label htmlFor="inputEmail" className="sr-only">
            Email address
        </label>
        <input
            type="email"
            id="inputEmail"
            className="form-control middle-field"
            placeholder="Email address"
            required
        />
        <label htmlFor="inputPassword" className="sr-only">
            Password
        </label>
        <input
            type="password"
            id="inputPassword"
            className="form-control mb-3 bottom-field"
            placeholder="Password"
            required
        />
        <button className="btn btn-lg btn-primary btn-block" type="submit">
            Create Account
        </button>
        <p className="mt-3 mb-3 text-muted text-center">
            Already registered ? <Link to="/auth/signin">Sign in here</Link>
        </p>
    </form>
}

export default Register
