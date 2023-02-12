import React from "react";
import { Routes, Route } from "react-router-dom";
import Signup from "./Signup/Signup";
import Login from "./Login/Login";
import Compte from "./Compte/Compte";
import Home from "./Home/Home";
import CharacterCreator from './CharacterCreator/CharacterCreator';

function App() {
  return (
      <Routes>
        <Route path="/signup" element={<Signup />} />
        <Route path="/login" element={<Login />} />
        <Route path="/" element={<Home />} />
        <Route path="/compte" element={<Compte />} />
		    <Route path="/create" element={<CharacterCreator />} />
      </Routes>
  );
}

export default App;
