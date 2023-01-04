import requestClient from '../client'
import { refresh_token } from '../queries'
import React, {useState, createContext, useContext, useEffect, ReactNode} from 'react'
import {GraphQLClient} from 'graphql-request'

//fixme remove {} as GraphQLClient
export const ClientContext = createContext({} as GraphQLClient)

export const useClient = () => useContext(ClientContext)

export const ClientWrapper = ({ children }: { children: ReactNode }) => {
  const [client] = useState(requestClient)

  return (
    <ClientContext.Provider value={client}>{children}</ClientContext.Provider>
  )
}
type Customer={
  email: string,
  isFetched : true,
  name: string,
  password: string
}| null
type CustomerContextType={
  customer:Customer
  setCustomer:  React.Dispatch<React.SetStateAction<Customer>>
}
//fixme remove {} as CustomerContextType
export const CustomerContext = createContext({} as CustomerContextType)

export const useCustomer = () => useContext(CustomerContext)



type ProviderProps = {
  children: ReactNode;
};




export const CustomerWrapper: React.FC<ProviderProps> = ({ children }) => {
  const client = useClient()
  const [customer, setCustomer] = useState<Customer>(null)
  const [working, setWorking] = useState(true)


  const refreshToken = () => {
    client
      .request(refresh_token)
      .then(({ refresh_token: { customer, token, expires_in } }) => {
        client.setHeader('authorization', `Bearer ${token}`)

        setTimeout(() => {
          refreshToken()
        }, (expires_in * 1000) - 500)

        setCustomer({...customer, isFetched: true })
      })
      .catch(console.log)
      .finally(() => {
        setWorking(false)
      })
  }

  useEffect(() => {
    refreshToken()
    // eslint-disable-next-line
  }, []);


  return (
    <CustomerContext.Provider value={{ customer , setCustomer }}>
      {working ? null : children}
    </CustomerContext.Provider>
  )
}
