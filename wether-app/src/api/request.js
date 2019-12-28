import fetchWithToaster from "../utilities/connectFetchWithToaster";

export const fetchLogin = (username, password) => (
    fetchWithToaster("http://localhost:4000/login",{
        method: "POST",
        credentials: "include",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
    })
);

export const fetchRegistration = (username, password) => (
    fetchWithToaster("http://localhost:4000/register",{
        method: "POST",
        credentials: "include",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
    })
);


export const fetchCurrentWeather = (username, password) => (
    fetchWithToaster("http://localhost:4000/api/getweather",{
        method: "POST",
        credentials: "include",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({ username, password }),
    })
);