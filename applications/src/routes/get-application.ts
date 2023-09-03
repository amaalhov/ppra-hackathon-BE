import express, { Request, Response } from 'express';
import { NotFoundError } from '@le-ma/common';
import { Application } from '../models/applications';

const router = express.Router();

router.get('/api/application/:id', async (req: Request, res: Response) => {
  const application = await Application.findById(req.params.id);

  if (!application) {
    throw new NotFoundError();
  }

  res.send(application);
});

export { router as showApplicationRouter };
