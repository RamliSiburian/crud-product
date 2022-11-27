import React, { useContext, useState } from 'react'
import { Alert, Button, Form, Modal } from 'react-bootstrap';
import { useMutation } from 'react-query';
import { useNavigate } from 'react-router-dom';
import { API } from '../../../Config/Api';
import { UserContext } from '../../../Contex/User-context'



function Login({ show, setShow }) {
    const handleClose = () => setShow(false)
    const [state, dispatch] = useContext(UserContext)
    const [message, setMessage] = useState(null);
    const navigate = useNavigate()

    const [formLogin, setFormLogin] = useState({
        username: "",
        password: "",
    })

    const handleOnChange = (e) => {
        setFormLogin({
            ...formLogin,
            [e.target.name]: e.target.value,
        })
    }

    const handleOnSubmit = useMutation(async (e) => {
        try {
            e.preventDefault()

            const dataLogin = await API.post("/Login", formLogin)
            dispatch({
                type: "LOGIN_SUCCESS",
                payload: dataLogin.data.data,
            })

            const alert = (
                <Alert variant="success" className="py-1">
                    Login berhasil
                </Alert>
            )
            setMessage(alert);

            const timer = setTimeout(navigates, 1000);

            function navigates() {
                setShow(false)
                navigate("/Dashboard")
                setMessage(null)

                setFormLogin({
                    username: "",
                    password: "",
                })
            }


        } catch (error) {
            const alert = (
                <Alert variant="danger" className="py-1">
                    Email atau password salah
                </Alert>
            )
            setMessage(alert);
        }
    })

    return (
        <Modal show={show} onHide={handleClose} size="sm" centered>
            <Modal.Header closeButton>
                <Modal.Title className='modal-title fs-1 fw-bold' style={{ color: "#ff4b00" }}>Login</Modal.Title>
            </Modal.Header>
            <Modal.Body>
                {message && message}
                <Form onSubmit={(e) => handleOnSubmit.mutate(e)}>
                    <Form.Group className="mb-3" controlId="formBasicUsername">
                        <Form.Label>Username</Form.Label>
                        <Form.Control
                            type="text"
                            name="username"
                            placeholder="Username"
                            onChange={handleOnChange}
                            value={formLogin.username}
                        />
                    </Form.Group>
                    <Form.Group className="mb-3" controlId="formBasicPassword">
                        <Form.Label>Password</Form.Label>
                        <Form.Control
                            type="password"
                            name="password"
                            placeholder="Password"
                            onChange={handleOnChange}
                            value={formLogin.password}
                        />
                    </Form.Group>
                    <Button type='submit' className='btn text-white fw-bold link w-100 border-0' style={{ background: "#ff4b00" }}>Login</Button>
                </Form>
            </Modal.Body>
        </Modal>
    )
}

export default Login;