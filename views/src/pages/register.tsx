import React, { SyntheticEvent, useState } from 'react'
import { BrowserRouter, Route, Link, useNavigate } from 'react-router-dom';
import Navbar from '../components/navbar';
import { localAuth } from '../helpers/localAuth';
import { toast } from 'react-toastify';
import "./register.css"

const Register = () => {
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const auth = localAuth();
    const navigate = useNavigate();

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();
        let regist = await fetch("http://localhost:8080/v1/user/register", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                username,
                email,
                password
            })
        }).then(res => res.json()).catch(err => err)
        if (regist.message.lastIndexOf('1062') > 0) {
            toast.error("Email has been used for another account!")
        } else {
            toast.success(regist.message)
            navigate("/")
        }
    }

    if (auth.status) return <><Navbar login={true} /><p className='d-flex justify-content-center'><span className='sr'>Telah Login.</span><Link to="/home">Beranda</Link></p></>
    return (
        <><Navbar login={false} />
        <div className='container-sm register'>
            <div className='d-flex justify-content-center rwRegister'>
                <div className="row content">

                    <form className='auth-wrapper' onSubmit={submit}>
                        <div className="logo"></div>
                        <h1 className="h3 mb-3 fw-normal">Sign Up</h1>
                        <input
                            className="form-control"
                            placeholder="username"
                            required
                            onChange={e => setUsername(e.target.value)}
                        />
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
                        <button className="w-100 btn btn-lg btn-primary" type="submit">
                            Submit
                        </button>
                    </form>
                </div>
            </div>
        </div>
        </>
    )
}

export default Register
