import React from "react";
import Joi from "joi-browser";
import { Navigate, useLocation, useNavigate } from "react-router-dom";
import Form from "../common/form";
import auth from "../services/authService";

export function withRouter(Children) {
    return (props) => {
        const navigate = useNavigate();
        const location = useLocation();
        try {
            const { from, teacherData } = location.state;
            return (
                <Children
                    {...props}
                    from={from}
                    teacherData={teacherData}
                    navigate={navigate}
                />
            );
        } catch (ex) {
            return <Children {...props} />;
        }
    };
}

class LoginForm extends Form {
    state = {
        data: { username: "", password: "" },
        errors: {},
    };

    schema = {
        username: Joi.string().required().email().label("Username"),
        password: Joi.string().required().label("Password"),
    };

    doSubmit = async () => {
        try {
            const { data } = this.state;
            await auth.login(data.username, data.password);

            if (this.props) {
                const { from, teacherData } = this.props;
                if (from && teacherData) {
                    this.props.navigate(from, {
                        state: { teacherData: teacherData },
                    });
                    this.props.navigate(0);
                } else if (from) {
                    this.props.navigate(from);
                    this.props.navigate(0);
                } else {
                    window.location = "/";
                }
            } else {
                window.location = "/";
            }
        } catch (ex) {
            if (ex.response && ex.response.status === 401) {
                const errors = { ...this.state.errors };
                errors.username = ex.response.data;
                this.setState({ errors });
            }
        }
    };

    render() {
        if (auth.getCurrentUser()) return <Navigate to="/" />;

        return (
            <div className="container-fluid min-vh-100">
                <div className="row mt-5" />
                <div className="row">
                    <div className="col-sm-8 bg-success text-white">
                        <h1>Login</h1>
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
                                    {this.renderInput(
                                        "username",
                                        "Username",
                                        "text",
                                        true
                                    )}
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
                                    {this.renderButton("Login")}
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

export default withRouter(LoginForm);