import React, { useContext, useState } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import { UserContext } from '../Contex/User-context';
import Login from './Modal/Auth/Login';


function Navbar() {
    const [showLogin, setShowLogin] = useState(false);
    const [state, dispatch] = useContext(UserContext)
    const navigate = useNavigate()

    function Logout() {
        dispatch({
            type: "AUTH_ERROR",
        });
        navigate("/");
        setShowLogin(true);
    }

    return (
        <div className='Navbar d-md-flex align-items-center'>
            <div className="container d-md-flex justify-content-between align-items-center" >
                <Link to="./"><img src="Image/logo.png" alt="Logo" width={150} /></Link>
                {state?.isLogin == true ? (
                    <button className='login' onClick={Logout}>Logout</button>
                ) : (
                    <button className='login' onClick={() => setShowLogin(true)}>Login</button>
                )}
            </div>

            <Login show={showLogin} setShow={setShowLogin} />
        </div >
    )
}

export default Navbar;