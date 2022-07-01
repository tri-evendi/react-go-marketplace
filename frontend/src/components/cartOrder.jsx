import React, { useContext, useEffect, useState } from 'react'
import { Button } from 'react-bootstrap';
// Contexts
import CartContext from '../contexts/CartContext';
const baseUrls = process.env.REACT_APP_SERVER_URL

function CartOrderPage() {
    const { cart } = useContext(CartContext);
    const [order, setOrder] = useState([]);

    useEffect(() => {
        setOrder(cart);
    }, [cart]);


    const handleSubmit = (e) => {
        e.preventDefault();
        const url = `${baseUrls}/api/order/bulk-create`;
        const options = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(order)
        };
        fetch(url, options)
            .then(res => res.json())
            .then(data => {
                console.log(data);
                setOrder([]);
            }
            )
            .catch(err => console.log(err));
    }


    const getCartTotal = () => {
        return cart
            .reduce((acc, value) => {
                return acc + value.Price;
            }, 0)
            .toFixed(2);
    };

    useEffect(() => {
        localStorage.getItem('cart', JSON.stringify(cart));
    }, [cart]);

    const removeItem = ID => {
        let items = JSON.parse(localStorage.getItem("cart"));
        items = items.filter((item) => item.ID !== ID);
        localStorage.setItem("cart", JSON.stringify(items));
        if (items.length === 0) {
            localStorage.removeItem("cart");
            window.location.reload();
        }
    };

    console.log(cart);
    console.log("order",order);
    return (
        <div className="container my-5">
            <h1>Cart</h1>
            <div className='row my-5'>
                <div className='col-6'>
                    {cart.map((item, index) => (
                        <div className='card mb-4' key={index}>
                            <div className="card-body p-2 row">
                                <div className='col-4 image-cart'>
                                    <img width="150px" src={baseUrls + '/' + item.ImagePath} alt={`${item.Name}`} />
                                </div>
                                <div className='col-8'>
                                    <h4>{item.Name}</h4>
                                    <h5>Rp {item.Price}</h5>
                                    <Button className='btn btn-danger btn-sm' onClick={() => removeItem(item.ID)}>
                                        Remove from cart
                                    </Button>
                                </div>
                            </div>
                        </div>
                    ))}
                </div>
                <div className='col-6 text-center'>
                    <h4 className='flex justify-content-end'>Total</h4>
                    <div className="flex justify-content-end">
                        <div className='total'>
                            <h4> Rp{getCartTotal()}</h4>
                        </div>
                    </div>
                    <div className="flex justify-content-end">
                        <div>
                            <Button className='btn btn-primary' onClick={handleSubmit}>Checkout</Button>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default CartOrderPage