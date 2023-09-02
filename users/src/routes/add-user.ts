import express, { Request, Response } from 'express';
import { body } from 'express-validator';
import {
  requireAuth,
  validateRequest,
  BadRequestError,
  currentUser,
} from '@le-ma/common';
import { User } from '../models/user';

const router = express.Router();

router.post(
  '/api/users',
  requireAuth,
  [
    // body('userId').not().isEmpty().withMessage('UID is required'),
    body('fullnames').not().isEmpty().withMessage('Fullnames is required'),
    body('idno').not().isEmpty().withMessage('Omang Number is required'),
    body('idcard').not().isEmpty().withMessage('Upload ID card is required'),
    body('verified')
      .not()
      .isEmpty()
      .withMessage('Verification status is required'),
    // body('email').not().isEmpty().withMessage('Email is required'),
    body('status').not().isEmpty().withMessage('Status is required'),
    body('role').not().isEmpty().withMessage('Role is required'),
  ],
  validateRequest,
  async (req: Request, res: Response) => {
    const { fullnames, idno, role, idcard, verified, status, email } = req.body;

    const existingUser = await User.findOne({ userId: req.currentUser!.id });

    if (existingUser) {
      throw new BadRequestError('Email registered to another user.');
    }

    const user = User.build({
      userId: req.currentUser!.id,
      fullnames,
      idno,
      email: req.currentUser!.email,
      role,
      idcard,
      verified,
      status,
    });
    await user.save();
    res.status(201).send(user);
  }
);

export { router as createUserRouter };
