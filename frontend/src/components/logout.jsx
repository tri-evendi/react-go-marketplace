import { Component } from "react";
import auth from "../services/authService";

class Logout extends Component {
    componentDidMount() {
        //localStorage.removeItem("currentUser");
        auth.logout();
        window.location = "/";
    }

    render() {
        return null;
    }
}

export default Logout;
