import React from 'react'
import {BrowserRouter, Route} from 'react-router-dom'


import Register from '../../pages/Register'
import Login from '../../pages/Login'
import Dashboard from '../../pages/Dashboard'
import {ApolloProvider} from '@apollo/client'
import {client} from '../../graphql/client'

function App() {
  return (
    <ApolloProvider client={client}>
      <BrowserRouter>
        <Route path="/auth/register" component={Register}/>
        <Route path="/auth/login" component={Login}/>
        <Route path="/" exact component={Dashboard}/>
      </BrowserRouter>
    </ApolloProvider>
  )
}

export default App
