import React from 'react'
import { ClientWrapper, CustomerWrapper } from '../../store'
import { Route, BrowserRouter } from 'react-router-dom'


import Login from '../../pages/Login'
import Register from '../../pages/Register'
import Dashboard from '../../pages/Dashboard'

function App() {
  return (
    <ClientWrapper>
      <CustomerWrapper>
        <BrowserRouter>
          <Route path="/auth/signin" component={Login} />
          <Route path="/auth/register" component={Register} />
          <Route path="/" exact component={Dashboard} />
        </BrowserRouter>
      </CustomerWrapper>
    </ClientWrapper>
  )
}

export default App
