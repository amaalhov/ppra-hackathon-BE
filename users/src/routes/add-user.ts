import express, { Request, Response } from 'express';
import { body } from 'express-validator';
import { requireAuth, validateRequest } from '@le-ma/common';
import { User } from '../models/users';

const router = express.Router();

router.post(
  '/api/users/add',
  requireAuth,
  [
    body('userId').not().isEmpty().withMessage('UID is required'),
    body('fullnames').not().isEmpty().withMessage('Fullnames is required'),
    body('idno').not().isEmpty().withMessage('Omang Number is required'),
    body('idcard').not().isEmpty().withMessage('Upload ID card is required'),
    body('verified')
      .not()
      .isEmpty()
      .withMessage('Verification status is required'),
    body('email').not().isEmpty().withMessage('Email is required'),
    body('status').not().isEmpty().withMessage('Status is required'),
    body('role').not().isEmpty().withMessage('Role is required'),
  ],
  validateRequest,
  async (req: Request, res: Response) => {
    const { fullnames, idno, email, role, idcard, verified, status } = req.body;
    const user = User.build({
      userId: req.currentUser!.id,
      fullnames,
      idno,
      email,
      role,
      idcard,
      verified,
      status,
    });
    await user.save();
    res.sendStatus(201).send(user);
  }
);

export { router as createUserRouter };
