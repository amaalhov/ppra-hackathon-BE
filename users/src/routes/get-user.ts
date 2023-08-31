import express, { Request, Response } from 'express';
import { NotFoundError } from '@le-ma/common';
import { User } from '../models/users';

const router = express.Router();

router.get('/api/users/:id', async (req: Request, res: Response) => {
  const user = await User.findById(req.params.id);

  if (!user) {
    throw new NotFoundError();
  }

  res.send(user);
});

export { router as getUserRouter };
