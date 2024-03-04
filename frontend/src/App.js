import React, { useState, createContext } from 'react';
import { BrowserRouter, Route, Routes } from "react-router-dom";

import Home from "./pages/home/Home";
import Login from './pages/login/Login';

const UserStateContext = createContext();

function App() {
    return (
        <Routes>
            <Route path="/" element={<Home/>}></Route>
            <Route path="/login" element={<Login/>}></Route>
        </Routes>
    );
}

export default App;
