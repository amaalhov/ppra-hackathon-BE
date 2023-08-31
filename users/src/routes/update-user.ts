import express, { Request, Response } from 'express';
import { body } from 'express-validator';
import {
  validateRequest,
  NotFoundError,
  requireAuth,
  NotAuthorizedError,
} from '@le-ma/common';
import { User } from '../models/user';

const router = express.Router();

router.put(
  '/api/users/:id',
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
    const user = await User.findById(req.currentUser?.id);

    console.log('user', user);

    if (!user) {
      throw new NotFoundError();
    }

    if (user.userId !== req.currentUser!.id) {
      throw new NotAuthorizedError();
    }

    user.set({
      fullnames: req.body.fullnames,
      idno: req.body.idno,
      role: req.body.role,
      idcard: req.body.idcard,
      verified: req.body.verified,
      status: req.body.status,
    });
    await user.save();
    // res.status(201);
    res.send(user);
  }
);

export { router as updateUserRouter };
