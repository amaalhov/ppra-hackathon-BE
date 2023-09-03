import express, { Request, Response } from 'express';
import { body } from 'express-validator';
import { requireAuth, validateRequest } from '@le-ma/common';
import { Application } from '../models/applications';

const router = express.Router();

router.post(
  '/api/application',
  requireAuth,
  // [
  //   body('title').not().isEmpty().withMessage('Title is required'),
  //   body('price')
  //     .isFloat({ gt: 0 })
  //     .withMessage('Price must be greater than 0'),
  // ],
  validateRequest,
  async (req: Request, res: Response) => {
    // const { descipline, code, description, subcodes } = req.body;

    const applicationData = req.body;

    const application = new Application(applicationData);

    // const application = Application.build({

    // });
    await application.save();

    res.status(201).send(application);
  }
);

export { router as createApplicationRouter };

// userId: string;
//   orderId: string;
//   descipline: string;
//   code: string;
//   description: string;
//   subcodes: object;
