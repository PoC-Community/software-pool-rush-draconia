import { RequestHandler } from "express";
import { getUser } from "~/dependencies/getUser";
require('dotenv').config();
import { User } from "~/types/Users";
import { promises as fs } from 'fs';

export const register: RequestHandler = async (req, res) => {
  const user = await getUser(req.body.email, null);

  if (!req.body.email)
    return res.status(401).send("please provide an email");
  if (user)
    return res.status(401).send("user already existing");

  if (req.body.password !== req.body.confirmationPassword)
    return res.status(401).send("password !== confirmation password");

  let users: User[];
  try {
    users = JSON.parse(await fs.readFile("database/users.json", "utf-8"));
  } catch (error) {
    console.log("database/users.json is not correct.\n", error);
    return res.status(500).send("error on database");
  }


  let newUser: User = {
    id: users.length,
    email: req.body.email,
    name: req.body.name,
    type: "user",
    password: req.body.password
  };
  users.push(newUser);

  let json_text = JSON.stringify(users);
  try {
    fs.writeFile("database/users.json", json_text, "utf-8");
  } catch (error) {
    console.log("database/users.json is not writable.\n", error);
    return res.status(500).send("error on database");
  }

  res.status(200).send("user registered");
}
