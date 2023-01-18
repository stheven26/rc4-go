import React, { useEffect, useState } from 'react';
import { BrowserRouter, Route, Routes, useNavigate } from 'react-router-dom';
import './App.css';
import { setLocalAuth } from './helpers/localAuth';
import Home from './pages/home';
import Login from './pages/login';
import Logout from './pages/logout';
import Register from './pages/register';
import { ToastContainer } from 'react-toastify';

import 'react-toastify/dist/ReactToastify.min.css';
import Help from './pages/help';
import About from './pages/about';

function App() {

  useEffect(() => { authenticate() }, [name]);

  return (
    <div className="App">
      <ToastContainer />
      <BrowserRouter>
        <main>
          <Routes>
            <Route path="/home" element={<Home />} />
            <Route path="/" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/logout" element={<Logout />} />
            <Route path='/help' element={<Help/>} />
            <Route path='/about' element={<About/>} />
          </Routes>
        </main>
      </BrowserRouter>
    </div >
  );

  async function authenticate() {
    await fetch("http://localhost:8080/v1/user", {
                    method: "GET",
                    headers: { "Content-Type": "application/json", "Cookie": document.cookie },
                    credentials: "include",
        })
        .then(res => res.json())
        .then(res => {
          if (res.message == "authorized") setLocalAuth({status: true})
          else setLocalAuth({status: false})
        })
        .catch(() => setLocalAuth({status: false}))
  }

}

export default App;
