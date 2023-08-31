import express, { Request, Response } from 'express';
import { NotFoundError, currentUser } from '@le-ma/common';
import { User } from '../models/user';

const router = express.Router();

router.get('/api/users/:id', async (req: Request, res: Response) => {
  //   const user = await User.findById(req.params.id);
  const user = await User.findById(req.currentUser?.id);

  if (!user) {
    throw new NotFoundError();
  }

  res.send(user);
});

export { router as getUserRouter };
