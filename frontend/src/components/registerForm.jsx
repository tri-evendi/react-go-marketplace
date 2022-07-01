import React from "react";
import Joi from "joi-browser";
import Form from "../common/form";
import * as userService from "../services/userService";
import auth from "../services/authService";

class RegisterForm extends Form {
    state = {
        data: { email: "", password: "", firstName: "", lastName: "" },
        errors: {},
    };

    schema = {
        email: Joi.string().required().email().label("Email"),
        password: Joi.string().required().min(5).label("Password"),
        firstName: Joi.string().required().label("First Name"),
        lastName: Joi.string().required().label("Last Name"),
    };

    doSubmit = async () => {
        try {
            const response = await userService.register(this.state.data);
            auth.loginWithJwt(response.headers["x-auth-token"]);
            window.location = "/login";
        } catch (ex) {
            if (ex.response && ex.response.status === 401) {
                const errors = { ...this.state.errors };
                errors.username = ex.response.data;
                this.setState({ errors });
            }
        }
    };

    render() {
        return (
            <div className="container-fluid min-vh-100">
                <div className="row mt-5" />
                <div className="row">
                    <div className="col-sm-8 bg-success text-white">
                        <h1>Register</h1>
                    </div>
                    <div className="col-sm-8" />
                </div>
                <div className="row" />
                <div className="row">
                    <div className="col-sm-8 bg-secondary">
                        <form onSubmit={this.handleSubmit}>
                            <div className="row mt-4" />
                            <div className="row">
                                <div className="col-sm-8">
                                    {this.renderInput("email", "Email")}
                                </div>
                                <div className="col-sm-8" />
                            </div>
                            <div className="row mt-3" />
                            <div className="row">
                                <div className="col-sm-8">
                                    {this.renderInput(
                                        "password",
                                        "Password",
                                        "password"
                                    )}
                                </div>
                                <div className="col-sm-8" />
                            </div>
                            <div className="row mt-3" />
                            <div className="row">
                                <div className="col-sm-8">
                                    {this.renderInput(
                                        "firstName",
                                        "First Name"
                                    )}
                                </div>
                                <div className="col-sm-8" />
                            </div>
                            <div className="row mt-3" />
                            <div className="row">
                                <div className="col-sm-8">
                                    {this.renderInput("lastName", "Last Name")}
                                </div>
                                <div className="col-sm-8" />
                            </div>
                            <div className="row mt-3" />
                            <div className="row">
                                <div className="col-sm-8">
                                    {this.renderButton("Register")}
                                </div>
                                <div className="col-sm-8" />
                            </div>
                            <div className="row mt-3" />
                        </form>
                    </div>
                    <div className="col-sm-8" />
                </div>
            </div>
        );
    }
}

export default RegisterForm;
