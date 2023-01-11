import React, { useEffect, useState, SyntheticEvent } from 'react'
import { json, useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';
import Navbar from '../components/navbar';
import { localAuth } from '../helpers/localAuth';
import './home.css'
import Login from './login';

const Home = () => {
    const [ username, setUsername ] = useState("User")
    const [dataUpload, setDataUpload] = useState('')

    let navigate = useNavigate();
    const auth = localAuth();
    useEffect(() => auth.status == 0 ? navigate("/login") : undefined, [])
    useEffect(() => {
        (async () => {
            let user = await getUserInfo()
            setUsername(user.data.user ? user.data.user.username : "User")
        })() 
    }, [])

    const handleChange = (e:any) => {
        setDataUpload(e.target.files[0])
    }

    console.log(dataUpload);
    

    const handleSubmit = async (e: SyntheticEvent) => {        
        const formData = new FormData();

        formData.append('file', dataUpload)

        e.preventDefault();
        const kirim = await fetch("http://localhost:8080/v1/file/upload", {
            method: "POST",
            credentials: "include",
            body: formData
        })
        .then((res) => 
            {
                res.json()
            }
        )
        .then((result) => {
            console.log('Success:', result);
        })
        .catch(err => {})
        // if (kirim && kirim.message == "Success") {
        //     toast.success("Success upload document")
        // } else {
        //     toast.error("Fail to upload document")
        // }
    }
    // const encrypt = async(e: SyntheticEvent) => {
    //     const kirim = await fetch("http://localhost:8080/v1/file/encrypt", {
    //         method: "POST",
    //         headers: {"Content-Type": "multipart/form-data", "Cookie": document.cookie },
    //         credentials: "include",
    //     }).then(res => res.json()).catch(err => err)
    //     if (kirim && kirim.message == "Success") {
    //         toast.success("Success upload document")
    //     } else {
    //         toast.error("Fail to upload document")
    //     }
    // }
    if (!auth.status) {
        return <><Navbar login={false} /><p className='d-flex justify-content-center'>You're not logged in!</p></>
    }
    if (auth.status) {
        return <><Navbar login={true} />
        <div className="container-sm">
             <p className='d-flex justify-content-center' id='hello'>
                 HI <b>{username}!</b>
             </p>
             <div className='container w-full'>
                 <div className='row'>
                     <div className='col-sm alert alert-info m-3'>
                         <form
                            // encType="multipart/form-data"
                            // action="http://localhost:8080/v1/file/upload"
                            // method="post"
                            // onSubmit={submit}
                            >
                            <input type="file" name="file" required onChange={handleChange}/>
                            <input type="submit" value="Upload" onClick={handleSubmit}/>
                        </form>
                     </div>
                     <div className='flex-container col-sm alert alert-warning m-3'>
                        <form 
                            action="http://localhost:8080/v1/file/encrypt"
                            method="post"
                            // onSubmit={encrypt}
                            >
                            <input type="submit" name="encrypt" value="Encrypt File"/>
                        </form>
                        <form 
                            action="http://localhost:8080/v1/file/decrypt"
                            method="post"
                            >
                            <input type="submit" name="decrypt" value="Decrypt File"/>
                            </form>
                     </div>
                 </div>
             </div>
         </div></>
    }
    return <><Navbar login={false} /><p className='d-flex justify-content-center'>loading...</p></>

    // async function upload() {
    //     const response = await fetch("http://localhost:8080/v1/file/upload", {
    //         method: "POST",
    //         headers: { "Content-Type": "application/json", "Cookie": document.cookie },
    //         credentials: "include",
    //     }).then(res => res.json()).catch(err => err)
    //     return response
    // }

    async function getUserInfo() {
        const response = await fetch("http://localhost:8080/v1/user", {
                        method: "GET",
                        headers: { "Content-Type": "application/json", "Cookie": document.cookie },
                        credentials: "include",
            }).then(res => res.json()).catch(err => err)
        if (response.available) {
            if (response.message == "Success"){
                toast.success(response.message)
            }else {
                toast.error(response.error)
            }
        }
        return response
    }
}


export default Home
