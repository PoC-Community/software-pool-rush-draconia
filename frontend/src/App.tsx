import React, { useState } from "react";
import { Routes, Route } from "react-router-dom";
import Signup from "./Signup/Signup";
import Login from "./Login/Login";
import Compte from "./Compte/Compte";
import Home from "./Home/Home";
import CharacterCreator from './CharacterCreator/CharacterCreator';

function App() {
  const [user, setUser] = useState("");
  return (
      <Routes>
        <Route path="/signup" element={<Signup user={user}/>} />
        <Route path="/login" element={<Login user={user} setUser={setUser}/>} />
        <Route path="/" element={<Home />} />
        <Route path="/compte" element={<Compte />} />
		    <Route path="/create" element={<CharacterCreator />} />
      </Routes>
  );
}

export default App;
