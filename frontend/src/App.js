import React, { useEffect, useState } from "react";
import "./App.css";
import "bootstrap";
import "bootstrap/dist/css/bootstrap.css";
import 'primereact/resources/themes/saga-blue/theme.css';
import 'primereact/resources/primereact.min.css';
import 'primeicons/primeicons.css';
import 'primeflex/primeflex.css';
import NavBar from "./components/navbar";
import HomePage from "./components/homepage";
import Footer from "./components/footer";
import RegisterForm from "./components/registerForm";
import LoginForm from "./components/loginForm";
import Logout from "./components/logout";
import NotFound from "./components/notFound";
import { Route, Routes } from "react-router-dom";
import auth from "./services/authService";
import MenuBar from "./components/menubar";
import CategoryPage from "./components/category";
import ProductDetailPage from "./components/productDetail";
import CardContext from './contexts/CartContext';
import CartOrderPage from "./components/cartOrder";


const App = () => {
    const initialCart = () => JSON.parse(localStorage.getItem('cart')) || [];

    const [cart] = useState(initialCart);
    const [user, setUser] = useState(null);

    useEffect(() => {
        localStorage.setItem('cart', JSON.stringify(cart));
        const users = auth.getCurrentUser();
        setUser(users);
    }, [cart]);

    console.log('carts',cart);
    return (
        <React.Fragment>
                <CardContext.Provider value={{ cart }}>
                    <NavBar user={user} />
                    <MenuBar />
                    <div className="container-fluid">
                        <Routes>
                            <Route path="/" element={<HomePage />} />
                            <Route path="/register" element={<RegisterForm />} />
                            <Route path="/login" element={<LoginForm />} />
                            <Route path="/logout" element={<Logout />} />
                            <Route path="/category/:id/product" element={<CategoryPage />} />
                            <Route path="/product/:id" element={<ProductDetailPage />} />
                            <Route path="/carts" element={<CartOrderPage />} />
                            <Route path="/not-found" element={<NotFound />} />
                            <Route path="*" element={<NotFound />} />
                        </Routes>
                    </div>
                    <Footer />
                </CardContext.Provider>
        </React.Fragment>
    );

}

export default App;
