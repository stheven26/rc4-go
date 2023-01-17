import React from 'react'
import Navbar from '../components/navbar';
import { localAuth } from '../helpers/localAuth';
import './about.css'
const About = () => {
    const auth = localAuth();
    if (!auth.status) {
        return <><Navbar login={false} /><p className='d-flex justify-content-center'>You're not logged in!</p></>
    }
    if (auth.status) {
        return <><Navbar login={true} />
        <div className="container-sm home">
             <div className='container w-full'>
                 <div className='row about'>
                    <h2>Aplikasi Pengamanan File Menggunakan Algoritma Kriptografi <i>Riverse Code 4</i> (RC4) Berbasis Web Pada Kopi Tyadatara </h2> <br /><br />
                    <div className="logo-univ"></div>
                    <p>
                        Nama: Iwan Dwi Mahendra
                    </p> <br />
                    <p>
                        NPM: 1611500198
                    </p>
                 </div>
             </div>
         </div></>
    }
    return <><Navbar login={false} /><p className='d-flex justify-content-center'>loading...</p></>
}

export default About