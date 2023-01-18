import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom';
import { setLocalAuth } from '../helpers/localAuth';
import Navbar from '../components/navbar';
import { toast } from 'react-toastify';

const Logout = () => {
    const [ logout, setLogout ] = useState("loging out...")
    const [ login, setLogin ] = useState(true)
    let navigate = useNavigate();

    useEffect(() => {
        (
            async () => {
                const response = await fetch("http://localhost:8080/v1/user/logout", {
                    method: "POST",
                    headers: { "Content-Type": "application/json", "Cookie": document.cookie },
                    credentials: "include",
                }).then(res => res.json()).catch(err => {
                    return { json: { name: "" } }
                })
                const content = await response
                if (content.message == "Success Logout") {
                    setLocalAuth({status: false})
                    navigate('/', { replace: false })
                }else {
                    toast.error("Failed Logout")
                }
            }
        )();
    });
    return <><Navbar login={false} /><p className='d-flex justify-content-center'>Logging out...</p></>
}

export default Logout
