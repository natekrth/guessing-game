import React, { useState, createContext } from 'react';
import { Route, Routes, Navigate } from "react-router-dom";

import Home from "./pages/home/Home";
import Login from './pages/login/Login';
import Register from './pages/register/Register';

const UserStateContext = createContext();

function App() {
    const [userState, setUserState] = useState(null); // Initial state is null

    // Function to handle login and set userState with token
    const handleLogin = (token) => {
        setUserState(token);
    };

    console.log(userState);

    return (
        <UserStateContext.Provider value={{ userState, handleLogin }}>
            <Routes>
                <Route path="/" element={<Login handleLogin={handleLogin} />} />
                <Route path="/register" element={<Register />} />
                <Route path="/home" element={<Home />} />
                <Route path="*" element={<Navigate to="/" />} />
            </Routes>
        </UserStateContext.Provider>
    );
}

export { UserStateContext };
export default App;
