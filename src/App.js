import './App.css';
import Navbar from './Component/Navbar';
import Home from './Pages/Home';
import { Routes, Route, useNavigate } from 'react-router-dom';
import { PrivateRoute } from './Component/Navigation';
import DashboardLogin from './Pages/Dashboard-login';
import { UserContext } from './Contex/User-context';
import { useContext, useEffect, useState } from 'react';
import { API, setAuthToken } from './Config/Api';

function App() {
  const [state, dispatch] = useContext(UserContext);
  const [isLoading, setIsLoading] = useState(true)
  const navigate = useNavigate()

  const checkUser = async () => {
    if (localStorage.token) {
      setAuthToken(localStorage.token);
    }
    try {

      const response = await API.get("/check-auth");
      let payload = response.data.data;
      payload.token = localStorage.token;

      dispatch({
        type: "USER_SUCCESS",
        payload,
      });
      if (response.data.code === 200) {
        setIsLoading(false)
        navigate("/Dashboard")
      }
    } catch (error) {
      if (error.response.data.code === 401) {
        // setIsLoading(false)
        navigate("/")
      }
    }
  };

  useEffect(() => {
    checkUser()
  }, []);

  return (
    <>
      <Navbar />
      <Routes>
        <Route exact path='/' element={<Home />} />
        <Route exact path='/' element={<PrivateRoute />}>
          <Route exact path='/Dashboard' element={<DashboardLogin />} />
        </Route>
      </Routes>

    </>
  );
}

export default App;