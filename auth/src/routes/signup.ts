import { Request, Response } from "express";
import jwt from "jsonwebtoken";
import { BadRequestError } from "@le-ma/common";

import { User } from "../models/user";

export async function signUp(req: Request, res: Response) {
  const { email, password } = req.body;

  const existingUser = await User.findOne({ email });

  if (existingUser) {
    throw new BadRequestError("Email in use");
  }

  const user = User.build({ email, password });
  await user.save();

  // Generate JWT
  const userJwt = jwt.sign(
    {
      id: user.id,
      email: user.email,
    },
    process.env.JWT_KEY!
  );

  // Store it on session object
  req.session = {
    jwt: userJwt,
  };

  res.status(201).send(user);
}
