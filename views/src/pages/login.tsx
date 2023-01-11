import React, { SyntheticEvent, useEffect, useState } from 'react'
import { Link } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';
import Navbar from '../components/navbar';
import { localAuth, setLocalAuth } from '../helpers/localAuth';
import { toast } from 'react-toastify';
import "./auth.css"

const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [redirect, setRedirect] = useState(false);
    const [ error, setError ] = useState(false)
    let navigate = useNavigate();
    const auth = localAuth()

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();
        const kirim = await fetch("http://localhost:8080/v1/user/login", {
            method: "POST",
            headers: { "Content-Type": "application/json", "Cookie": document.cookie },
            credentials: "include",
            body: JSON.stringify({
                email,
                password
            })
        })
        .then(res => res.json())
        .catch(err => {})
        if (kirim && kirim.message == "Success") {
            toast.success("Login berhasil!")
            setLocalAuth({status: true})
            navigate("/", {replace: true})
        } else {
            toast.error(kirim ? kirim.message : "Email Incorrect!")
            setError(true)
        }
    }
    if (auth.status) return <><Navbar login={true} /><p className='d-flex justify-content-center'><span className='sr'>Telah Login.</span><Link to="/">Beranda</Link></p></>
    if (!auth.status || error) return <><Navbar login={false} />
                            <div className="container-sm">
                                <div className='d-flex justify-content-center'>
                                    <form className="auth-wrapper" onSubmit={submit}>
                                        <h1 className="h3 mb-3 fw-normal">Please login</h1>
                                        <input
                                            type="email"
                                            className="form-control"
                                            placeholder="email@example.com"
                                            required
                                            onChange={e => setEmail(e.target.value)}
                                            />
                                        <input
                                            type="password"
                                            className="form-control"
                                            placeholder="password"
                                            required
                                            onChange={e => setPassword(e.target.value)}
                                            />
                                        <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
                                    </form>
                                </div>
                            </div></>

    return <><p className='d-flex justify-content-center'>loading...</p></> 
}

export default Login
