import React from 'react'
import { Link } from 'react-router-dom'
import { localAuth } from '../helpers/localAuth'

const Navbar = (props: { login: boolean }) => {


    let notLogedIn = props.login 
        ?   <ul className="navbar-nav me-auto mb-2 mb-md-0">
                <li className="nav-item">
                    <Link to="/logout" className="nav-link active">
                        Logout
                    </Link>
                </li>
            </ul>
        :   <ul className="navbar-nav me-auto mb-2 mb-md-0">
                <li className="nav-item">
                    <Link to="/login" className="nav-link active">
                        Login
                    </Link>
                </li>
                <li className="nav-item">
                    <Link to="/register" className="nav-link active">
                        Register
                    </Link>
                </li>
            </ul>
    return (
        <nav className="navbar navbar-expand-md navbar-dark bg-dark mb-4">
            <div className="container-fluid">
                <Link to="/" className="navbar-brand">
                    Home
                </Link>
                <div className="collapse navbar-collapse">
                    {
                        notLogedIn
                    }
                </div>
            </div>
        </nav>
    )
}

export default Navbar
