import { format } from 'node:path/win32';
import React, { useEffect, useState, SyntheticEvent, useRef } from 'react'
import { json, useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';
import Navbar from '../components/navbar';
import { localAuth } from '../helpers/localAuth';
import './home.css'
import Login from './login';

const Home = () => {
    const [ username, setUsername ] = useState("User");
    const [dataUpload, setDataUpload] = useState('');
    let encryptMessage:any; let decryptMessage:any;
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
    const handleSubmit = async (e: SyntheticEvent) => {        
        const formData = new FormData();
        formData.append('file', dataUpload)
        e.preventDefault();
        await fetch("http://localhost:8080/v1/file/upload", {
            method: "POST",
            credentials: "include",
            body: formData
        })
        .then(async (res) => {
            const result = await res.json();
            if(result.message === 'Success'){ 
                encryptMessage = "Success upload"; decryptMessage = "Success upload"
                toast.success("Success upload document")
            } else {
                toast.error("Fail to upload document")
            }
        })
        .catch(err => {})
    }
    const handleEncrypt = async (e:SyntheticEvent) => {
        e.preventDefault();
        await fetch("http://localhost:8080/v1/file/encrypt", {
            method: "POST",
            credentials:"include"
        })
        .then(resp => resp.blob())
        .then(blob => {
            if(encryptMessage == "Success upload"){
                let url = window.URL.createObjectURL(blob);
                const a = document.createElement("a");
                a.style.display = "none";
                a.href = url;
                a.download = "encrypt.txt";
                document.body.appendChild(a);
                a.click();
                window.URL.revokeObjectURL(url);
                toast.success("Success encrypt document")
                encryptMessage = ""
            }else {
                toast.error("Fail to encrypt document")
            }
        })
        .catch(() => alert("oh no!"));
    }
    const handleDecrypt = async (e:SyntheticEvent) => {
        e.preventDefault();
        const kirim = await fetch("http://localhost:8080/v1/file/decrypt", {
            method: "POST",
            credentials:"include"
        })
        .then(resp => resp.blob())
        .then(blob => {
            if(decryptMessage == "Success upload"){
                let url = window.URL.createObjectURL(blob);
                const a = document.createElement("a");
                a.style.display = "none";
                a.href = url;
                a.download = "decrypt.pdf";
                document.body.appendChild(a);
                a.click();
                window.URL.revokeObjectURL(url);
                toast.success("Success decrypt document")
                decryptMessage = ""
            }else {
                toast.error("Fail to decrypt document")
            }
        })
        .catch(() => alert("oh no!"));
    }
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
                         <form className='upload'
                            >
                            <input type="file" name="file" required onChange={handleChange}/>
                            <input type="submit" value="Upload" onClick={handleSubmit}/>
                        </form>
                     </div>
                     <div className='flex-container col-sm alert alert-warning m-3'>
                        <form 
                            //   action="http://localhost:8080/v1/file/encrypt"
                            //   method="post"
                            >
                            <input type="submit" name="encrypt" value="Encrypt File" onClick={handleEncrypt}/>
                        </form>
                        <form 
                            // action="http://localhost:8080/v1/file/decrypt"
                            // method="post"
                            >
                            <input type="submit" name="decrypt" value="Decrypt File" onClick={handleDecrypt}/>
                            </form>
                     </div>
                 </div>
             </div>
         </div></>
    }
    return <><Navbar login={false} /><p className='d-flex justify-content-center'>loading...</p></>

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
