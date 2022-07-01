import React, { useContext, useEffect, useState } from "react";
import "bootstrap";
import "bootstrap/dist/css/bootstrap.css";
import "./navbar.css";
import { NavLink } from "react-router-dom";
import http from "../services/httpService";
// Contexts
import CartContext from '../contexts/CartContext';
const apiEndpoint = process.env.REACT_APP_SERVER_URL + "/api";

const MenuBar = ({ user }) => {
   const [category, setCategory] = useState([])
   const { cart } = useContext(CartContext);

   useEffect(() => {
      http.get(`${apiEndpoint}/categories`).then(res => {
         const newData = [...res.data.data];
         setCategory(newData);
      });
   }, []);

   useEffect(() => {
      localStorage.getItem('cart', JSON.stringify(cart));
   },[cart]);

   return (
      <nav className="navbar navbar-expand-sm bg-light m-2 p-3">
         <div className="container-fluid">
            <div>
               {category?.map((item, index) => {
                  return (
                     <NavLink className="p-2 text-black text-decoration-none" to={`/category/${item.ID}/product`} key={index}>{item.Name}</NavLink>
                  )
               })}
            </div>
            <div>
               <NavLink to="/carts" reloadDocument className="btn btn-sm position-relative">
                  <i className="pi pi-shopping-cart mr-2"></i>
                  <span className="position-absolute top-0 start-200 translate-middle badge rounded-pill bg-danger">
                     {cart?.length}
                     <span className="visually-hidden">unread messages</span>
                  </span>
               </NavLink>
            </div>
         </div>
      </nav>
   );
};

export default MenuBar;
