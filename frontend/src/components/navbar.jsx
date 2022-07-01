import React from "react";
import "bootstrap";
import "bootstrap/dist/css/bootstrap.css";
import "./navbar.css";
import { NavLink } from "react-router-dom";
import Logo from "../images/logo.png";

const NavBar = ({ user }) => {
   return (
      <nav className="navbar navbar-expand-sm navbar-dark bg-dark sticky-top">
         <div className="container-fluid">
            <NavLink className="navbar-brand" to="/">
               <img
                  width="50px"
                  height="40px"
                  src={Logo}
                  className="d-inline-block align-text-center"
                  alt=""
               />
               <span className="text-primary">Voucher Marketplaces</span>
            </NavLink>
            {!user && (
               <React.Fragment>
                  <ul className="list-unstyled">

                  <li className="nav-item dropdown">
                     <NavLink
                        className="nav-link dropdown-toggle bg-dark text-light"
                        to="#"
                        role="button"
                        data-bs-toggle="dropdown"
                     >
                        Login
                     </NavLink>
                     <ul className="dropdown-menu me-5 bg-dark">
                        <li>
                           <NavLink
                              className="dropdown-item bg-dark text-light"
                              to="/login"
                           >
                              Login
                           </NavLink>
                        </li>
                        <li>
                           <NavLink
                              className="dropdown-item bg-dark text-light"
                              to="/register"
                           >
                              Sign-up
                           </NavLink>
                        </li>
                     </ul>
                  </li>
                  </ul>
               </React.Fragment>
            )}
            {user && (
               <React.Fragment>
                  <ul className="list-unstyled">
                  <li className="nav-item dropdown">
                     <NavLink
                        className="nav-link dropdown-toggle bg-dark text-light"
                        to="#"
                        role="button"
                        data-bs-toggle="dropdown"
                     >
                        {user.Email}
                     </NavLink>
                     <ul className="dropdown-menu me-5 bg-dark">
                        <li>
                           <NavLink
                              className="dropdown-item bg-dark text-light"
                              to="/order-history"
                           >
                              Order History
                           </NavLink>
                        </li>
                        <li>
                           <NavLink
                              className="dropdown-item bg-dark text-light"
                              to="/logout"
                           >
                              Logout
                           </NavLink>
                        </li>
                     </ul>
                  </li>
                  </ul>
               </React.Fragment>
            )}
         </div>
      </nav>
   );
};

export default NavBar;
