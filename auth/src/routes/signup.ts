import express, { Request, Response } from "express";
import { body, validationResult } from "express-validator";

const router = express.Router();

router.post(
  "/api/users/signup",
  [
    body("email").isEmail().withMessage("Email is not valid"),
    body("password")
      .trim()
      .isLength({ min: 4, max: 20 })
      .withMessage("Password must be 4 or more characters long"),
  ],
  (req: Request, res: Response) => {
    const errors = validationResult(req);

    if (!errors.isEmpty()) {
      return res.status(400).send(errors.array());
    }
    //   res.send("sign up");
    const { email, password } = req.body;

    console.log("Creating a user...");

    res.send({});

    //   if(!email || typeof email !== 'string' ) {

    //   }
  }
);

export { router as signupRouter };
