import React, { useEffect, useState } from "react";
import "bootstrap/dist/css/bootstrap.css";
import "./homepage.css";
import { Card } from 'react-bootstrap';
import http from "../services/httpService";
import { NavLink, useParams } from "react-router-dom";

const baseUrls = process.env.REACT_APP_SERVER_URL
const apiEndpoint = `${baseUrls}/api`;

const CategoryPage = () => {
    const [products, setProducts] = useState([])
    const params = useParams();

    const categoryID = Number(params.id);
    useEffect(() => {
        http.get(`${apiEndpoint}/category/${categoryID}/products`).then(res => {
            const newData = [...res.data.data];
            setProducts(newData);
        });
    }, [categoryID]);

    return (
        <div className="container-fluid container-xxl">
            <div className="row my-4 justify-content-center">
                {products?.map((item, index) => {
                    return (
                        <div className="col-md-3 col-xxl-3 justify-content-center my-4" key={index}>
                            <Card>
                                <NavLink to={`/product/${item.ID}`}>
                                    <Card.Img variant="top" src={baseUrls + '/' + item.ImagePath} />
                                </NavLink>
                                <Card.Body>
                                    <Card.Title>
                                        <NavLink className="text-black text-decoration-none" to={`/product/${item.ID}`}>{item.Name}</NavLink>
                                    </Card.Title>
                                    <Card.Text>
                                        {item.Description}
                                    </Card.Text>
                                    <div className=" flex justify-content-end">
                                        <h5><span className="badge p-2 bg-primary">Rp. {item.Price}</span></h5>
                                    </div>
                                </Card.Body>
                            </Card>
                        </div>
                    )
                }
                )}
            </div>
        </div>
    );

}

export default CategoryPage;
