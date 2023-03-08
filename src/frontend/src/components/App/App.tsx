import React from 'react'
import {BrowserRouter, Route} from 'react-router-dom'


import Register from '../../pages/Register'
import Login from '../../pages/Login'
import Dashboard from '../../pages/Dashboard'
import {ApolloProvider} from '@apollo/client'
import {client} from '../../graphql/client'
import {Apps} from '../../pages/Apps'
import {Features} from '../../pages/Features'

function App() {
  return (
    <ApolloProvider client={client}>
      <BrowserRouter>
        <Route path="/auth/register" component={Register}/>
        <Route path="/auth/login" component={Login}/>
        <Route path="/apps" exact component={Apps}/>
        <Route path="/features" exact component={Features}/>
        <Route path="/" exact component={Dashboard}/>
      </BrowserRouter>
    </ApolloProvider>
  )
}

export default App
