import React from 'react'
import Navbar from '../components/navbar';
import { localAuth } from '../helpers/localAuth';
import './help.css'
const Help = () => {
    const auth = localAuth();
    if (!auth.status) {
        return <><Navbar login={false} /><p className='d-flex justify-content-center'>You're not logged in!</p></>
    }
    if (auth.status) {
        return <><Navbar login={true} />
        <div className="container-sm home">
             <div className='container w-full'>
                 <div className='row help'>
                    <h2>Cara Pengunaan: </h2> <br /><br />
                    <ol>
                        <li>Klik choosen file, lalu pilih file yang ingin di upload.</li>
                        <li>Format file yang dapat di upload adalah pdf, doc, txt dan xlsx dengan besar kapasitas 10 mb.</li>
                        <li>Setelah memilih file, klik upload.</li>
                        <li>Setelah berhasil upload file, maka dapat mengakses layanan enskripsi dan deskripsi.</li>
                        <li>Klik Encrypt File untuk mengenkripsi file yang telah di upload.</li>
                        <li>Klik Decrypt File untuk mendekripsi file yang telah di upload.</li>
                        <li>File yang telah di Encrypt/Decrypt akan otomatis terdownload.</li>
                    </ol>
                 </div>
             </div>
         </div></>
    }
    return <><Navbar login={false} /><p className='d-flex justify-content-center'>loading...</p></>
}

export default Help