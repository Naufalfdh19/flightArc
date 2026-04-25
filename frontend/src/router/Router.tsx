import {
  createBrowserRouter,
} from "react-router-dom";

import HomePage from "../pages/Homepage/Homepage";
import Login from "../pages/Login/Login";
import Register from "../pages/Register/Register";

function useRouter() {
    return createBrowserRouter([
            {
                path: "/",
                element: <HomePage/>,
            },
            {
                path: "/login",
                element: <Login/>
            },
            {
                path: "/register",
                element: <Register/>
            }
        ]
    )
}

export default useRouter