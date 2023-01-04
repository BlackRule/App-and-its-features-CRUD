import { GraphQLClient } from 'graphql-request'

export default new GraphQLClient(
  'http://localhost:80/query',
  {
    credentials: 'include',
  }
)
