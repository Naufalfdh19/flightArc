import useRouter from './router/Router'
import { RouterProvider } from 'react-router-dom'

function App() {
  const router = useRouter();

  return (
    <RouterProvider router={router}/>
  )
}

export default App
