import http from "./httpService";
 
const baseUrls = process.env.REACT_APP_SERVER_URL;
const apiEndpoint =   `${baseUrls}/api/register`;

export function register(user) {
    return http.post(apiEndpoint, {
        email: user.email,
        password: user.password,
        firstName: user.firstName,
        lastName: user.lastName,
    });
}
