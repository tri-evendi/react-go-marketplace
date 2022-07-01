import React, { useEffect, useState } from 'react'
import { Button } from 'react-bootstrap';
import { useNavigate, useParams } from 'react-router-dom';
import auth from '../services/authService';
import http from '../services/httpService';

const baseUrls = process.env.REACT_APP_SERVER_URL
const apiEndpoint = `${baseUrls}/api`;

const ProductDetailPage = () => {
  const [detail, setDetail] = useState({});
  const [user, setUser] = useState(null);
  const [cart, setCart] = useState(JSON.parse(window.localStorage.getItem('cart')));
  const navigate = useNavigate();

  const params = useParams();
  const productID = Number(params.id);

  useEffect(() => {
    localStorage.getItem('cart', JSON.stringify(cart));
    const users = auth.getCurrentUser();
    setUser(users);
  }, [cart]);
  
  useEffect(() => {
    http.get(`${apiEndpoint}/product/${productID}`).then(res => {
      const newData = res.data.data;
      // console.log(newData)
      setDetail(newData);
    });
  }, [productID]);

  async function fetchData() {
    localStorage.setItem('cart', JSON.stringify(cart));
  }

  useEffect(() => {
    localStorage.setItem('cart', JSON.stringify(cart));
    fetchData();
  })

  const handleOnclick = (e) => { 
    e.preventDefault();
    console.log('You clicked submit.');
    if (user) {
      if (!cart.find(cartItem => cartItem.ID === detail.ID)) {
        setCart([...cart, detail]);
        fetchData();
        window.location.reload().then(() => {
          navigate('/carts');
        });
      }
    } else {
      navigate('/login');
    }
  }

  return (
    <div className='container my-4'>
      {detail && (
        <div className='row p-4'>
          <div className='col-md-6'>
            <div className='card'>
              <img src={baseUrls + '/' + detail.ImagePath} alt='product' />
            </div>
          </div>
          <div className='col-md-6'>
            <h1>{detail.Name}</h1>
            <h3>Rp. {detail.Price}</h3>
            <div className='my-5'>
              <p>{detail.Description}</p>
            </div>
          </div>
        </div>
      )}
      <div className='flex justify-content-end'>
        <Button className="btn btn-primary" onClick={handleOnclick}> Masukkan Keranjang</Button>
      </div>
    </div>
  )
}

export default ProductDetailPage