import {useQuery} from '@apollo/client'
import {APPS} from '../graphql/queries'

export const Apps = () => {
  const {data, loading, error} = useQuery(APPS)
  const f=()=>{
    if (loading) return <p>Loading apps...</p>
    if (error) return <p>{error.message}</p>
    else {
      return data?.apps.map((app)=>
        <div>{app.name}</div>
      )
    }
  }
  return (
    <>
      <div>Apps</div>
      {f()}
    </>
  )
}