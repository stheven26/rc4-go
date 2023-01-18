import React, { useEffect, useState, SyntheticEvent, useRef } from 'react'
import { useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';
import Navbar from '../components/navbar';
import { localAuth } from '../helpers/localAuth';
import './home.css'

const Home = () => {
    const [ username, setUsername ] = useState("User");
    const [user, setUser] = useState("");
    const [doc, setDoc] = useState("")
    const [dataUpload, setDataUpload] = useState('');
    let encryptMessage:any; let decryptMessage:any;
    let navigate = useNavigate();
    const auth = localAuth();
    useEffect(() => auth.status == 0 ? navigate("/") : undefined, [])
    useEffect(() => {
        (async () => {
            let user = await getUserInfo()
            setUsername(user.data.user ? user.data.user.username : "User")
        })() 
    }, [])
    useEffect(() => {
        (async () => {
            let user = await getAllUser()
            setUser(user.data)            
        })() 
    }, [])
    useEffect(() => {
        (async () => {
            let doc = await getTotalDoc()
            setDoc(doc.data)            
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
                a.download = "decrypt";
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
        <div className="container-sm home">
             <p className='d-flex justify-content-center' id='hello'>
                 HI <b>{username}!</b>
             </p>
             <div className='container w-full'>
                <div className="row">
                    <div className='col-sm alert alert-info m-3 info'>
                         Total User: {user}
                     </div>
                     <div className='col-sm alert alert-info m-3 info'>
                         Upload Dokumen: {doc}
                     </div>
                </div>
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
                            >
                            <input type="submit" name="encrypt" value="Encrypt File" onClick={handleEncrypt}/>
                        </form>
                        <form 
                            >
                            <input type="submit" name="decrypt" value="Decrypt File" onClick={handleDecrypt}/>
                            </form>
                     </div>
                 </div>
                 <div className='row about-home'>
                    <h2>Kopi Tydatara</h2> <br /><br />
                    <p>
                    Kopi Tydatara adalah UMKM dibidang kuliner terutama pada minuman yang terbuat dari kopi tydatara ini terbentuk pada tahun 2015 sampai saat ini, Tyadatara kopi juga dapat membantu masyarakat untuk mengenal lebih banyak tentang kopi, dari cara penyajian maupun jenis jenis biji kopi dari manca negara. <br />
                    </p>
                    <h3>Visi</h3><br />
                    <p>
                    Visi dari kedai kopi ini adalah menjadi kedai kopi yang menawarkan suasana, kondisi, tempat, menu yang bervariatif dengan cita rasa dan kualitas kopi Tyadatara yang baik dan dapat memenuhi selera para pengunjung dan pelanggannya sehingga dapat meningkatkan loyalitas pelanggan dan pantas menjadi ikon kuliner di kota Tanggerang , serta dapat mengembangkan bisnis kedai kopi ini ke berbagai daerah untuk mengenalkan dan membudayakan minum kopi produk kopi Tyadatara serta membuka lowongan pekerjaan di tengah masyarakat. <br />
                    </p>
                    <h3>Misi</h3><br />
                    <ol>
                        <li>Menyediakan coffee yang berkualitas.</li>
                        <li>Menyediakan tempat yang nyaman untuk berkumpul dan bersantai.</li>
                        <li>Menempatkan pelanggan sebagai prioritas.</li>
                        <li>Memberikan pelayanan yang prima dan unggul dalam penyajian.</li>
                        <li>Memotivasi karyawan dalam meraih mimpi.</li>
                    </ol>
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
    async function getAllUser() {
        const response = await fetch("http://localhost:8080/v1/user/all", {
                        method: "GET",
                        headers: { "Content-Type": "application/json"},
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
    async function getTotalDoc() {
        const response = await fetch("http://localhost:8080/v1/file", {
                        method: "GET",
                        headers: { "Content-Type": "application/json"},
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
