import {gql} from './__generated__'

// If we want to store these in .qraphql files https://www.apollographql.com/docs/react/integrations/webpack/
// then cra eject or rewiring needed

export const REGISTER = gql(`
  mutation register($email: String! $password: String!) {
    register(input: { email: $email, password: $password }) {
      token
      user {
        id
        email
      }
    }
  }
`)

export const LOGIN = gql(`
  mutation login($email: String! $password: String!) {
    login(input: { email: $email, password: $password }) {
      token
      user {
        id
        email
      }
    }
  }
`)
