import {Header} from '../components/Header/Header'
import {useQuery} from '@apollo/client'
import {FEATURES} from '../graphql/queries'

export const Features = () => {
  const {data, loading, error} = useQuery(FEATURES)
  const f=()=>{
    if (loading) return <p>Loading features...</p>
    if (error) return <p>{error.message}</p>
    else {
      return data?.features.map((feature)=>
        <div>{feature.name}</div>
      )
    }
  }
  return (
    <>
      <Header/>
      <div>Features</div>
      {f()}
    </>
  )
}