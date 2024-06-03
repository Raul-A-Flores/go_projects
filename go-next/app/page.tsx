async function getData() {
  const res = await fetch('http://localhost:8080/api/todos')
  console.log(res)

  

  if (!res.ok) {
    // This will activate the closest `error.js` Error Boundary
    throw new Error('Failed to fetch data')
  }
  console.log(res.json)
  return res.json()
}
 
export default async function Page() {

  const data = await getData()
 
  return <main>
    {data}
    HELLO
  </main>
}