import React from 'react'
import ReactDOM from 'react-dom/client'
import './index.css'
import {client} from './graphql/client'
import {createBrowserRouter, RouterProvider} from 'react-router-dom'
import {ApolloProvider} from '@apollo/client'
import {App} from './pages/App'
import {Apps} from './pages/Apps'
import Register from './pages/Register'
import Login from './pages/Login'
import {Features} from './pages/Features'
import Main from './pages/Main'

const router = createBrowserRouter([
  {
    children: [
      {
        element: <Register/> ,
        path: '/auth/register',
      },
      {
        element: <Login/> ,
        path: '/auth/login',
      },
      {
        element: <Apps/> ,
        path: '/apps',
      },
      {
        element: <App />,
        path: 'apps/:appId',
      },
      {
        element: <Features/> ,
        path: '/features',
      },
      {
        element: <Main/> ,
        path: '/',
      },
    ],
    element: <div>Hello world!</div>,
    path: '/',
  },

])

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
)

root.render(
  <React.StrictMode>
    <ApolloProvider client={client}>
      <RouterProvider router={router} />
    </ApolloProvider>
  </React.StrictMode>
)