const development = !process.env.NODE_ENV || process.env.NODE_ENV === 'development';
const api_endpoint = "_api";
const apiUrlPro = "/";
const apiUrlDev = "http://localhost/";

export const config = {
    development,
    homePath: development ? "" : "",
    apiUrl: (development ? apiUrlDev : apiUrlPro) + api_endpoint,
};