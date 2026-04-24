import {
  createBrowserRouter,
} from "react-router-dom";

import HomePage from "../pages/Homepage/Homepage";
import Login from "../pages/Login/Login";

function useRouter() {
    return createBrowserRouter([
            {
                path: "/",
                element: <HomePage/>,
            },
            {
                path: "/login",
                element: <Login/>
            }
            
        ]
    )
}

export default useRouter