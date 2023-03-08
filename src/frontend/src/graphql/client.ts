import {ApolloClient, createHttpLink, InMemoryCache} from '@apollo/client'
import {setContext} from '@apollo/client/link/context'

/*
* This code automatically reads and sets Authorization header after localStorage.setItem('jwt_token', jwt_token) in SignIn mutation.
* */

const apolloHttpLink = createHttpLink({
  uri: 'http://localhost:80/query'
})

const apolloAuthContext = setContext(async (_, {headers}) => {
  const jwt_token = localStorage.getItem('jwt_token')
  return {
    headers: {
      ...headers,
      Authorization: jwt_token ? `Bearer ${jwt_token}` : ''
    },
  }
})

export const client = new ApolloClient({
  cache: new InMemoryCache(),
  link: apolloAuthContext.concat(apolloHttpLink),
})


