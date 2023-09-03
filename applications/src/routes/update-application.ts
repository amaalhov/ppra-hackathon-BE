import express, { Request, Response } from 'express';
import { body } from 'express-validator';
import {
  validateRequest,
  NotFoundError,
  requireAuth,
  NotAuthorizedError,
  BadRequestError,
} from '@le-ma/common';
import { Application } from '../models/applications';

const router = express.Router();

router.put(
  '/api/application/:id',
  requireAuth,
  [
    body('title').not().isEmpty().withMessage('Title is required'),
    body('price')
      .isFloat({ gt: 0 })
      .withMessage('Price must be provided and must be greater than 0'),
  ],
  validateRequest,
  async (req: Request, res: Response) => {
    const application = await Application.findById(req.params.id);

    if (!application) {
      throw new NotFoundError();
    }

    if (application.orderId) {
      throw new BadRequestError('Cannot edit a reserved ticket');
    }

    if (application.userId !== req.currentUser!.id) {
      throw new NotAuthorizedError();
    }

    application.set({
      title: req.body.title,
      price: req.body.price,
    });
    await application.save();

    res.send(application);
  }
);

export { router as updateApplicationRouter };
