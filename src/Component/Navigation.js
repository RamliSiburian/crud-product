import { useContext } from "react";
import { Navigate, Outlet } from "react-router-dom";
import { UserContext } from "../Contex/User-context";

export const PrivateRoute = () => {
    const [state, dispatch] = useContext(UserContext)

    const Logins = state.isLogin ? <Outlet /> : <Navigate to="/" />
    return Logins;
} 