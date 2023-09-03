import express, { Request, Response, application } from 'express';
import { Application } from '../models/applications';

const router = express.Router();

router.get('/api/application', async (req: Request, res: Response) => {
  const tickets = await Application.find({});

  res.send(application);
});

export { router as indexApplicationRouter };
