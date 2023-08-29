import express from "express";
import "express-async-errors";
import { body } from "express-validator";
import { json } from "body-parser";
import cookieSession from "cookie-session";
import { currentUser, errorHandler, NotFoundError, validateRequest } from "@le-ma/common";

import { signInUsingEmailAndPassword } from "./routes/signin";
import { signUp } from "./routes/signup";

const app = express();

app.set("trust proxy", true);
app.use(json());
app.use(
  cookieSession({
    signed: false,
    secure: process.env.NODE_ENV !== "test",
  })
);

app.use(errorHandler);

app.post(
  "/api/user/signup",
  [
    body("email").isEmail().withMessage("email must be valid"),
    body("firstName").trim().notEmpty().withMessage("first name is required"),
    body("lastName").trim().notEmpty().withMessage("last name is required"),
    body("phoneNumber").trim().notEmpty().withMessage("first name is required"),
    body("password")
      .trim()
      .notEmpty()
      .withMessage("password required"),
  ],
  validateRequest,
  signUp
)

app.post(
  "/api/user/signin",
  [
    body("email").isEmail().withMessage("email must be valid"),
    body("password")
      .trim()
      .isLength({ min: 4, max: 20 })
      .withMessage("password must be between 4 and 20 characters"),
  ],
  validateRequest,
  signInUsingEmailAndPassword
)

app.post(
  "/api/user/signout",
  async (req, res) => {
    req.session = null;
    res.send({})
});


app.get("/api/user/currentuser", currentUser, (req, res) => {
  res.send({ currentUser: req.currentUser || null});
});

app.all("*", async () => {
  throw new NotFoundError();
});

export { app };
