import express from 'express';
import 'express-async-errors';
import { json } from 'body-parser';
import cookieSession from 'cookie-session';
import { errorHandler, NotFoundError, currentUser } from '@le-ma/common';

import { createUserRouter } from './routes/add-user';
import { getUserRouter } from './routes/get-user';
import { indexUserRouter } from './routes';
import { updateUserRouter } from './routes/update-user';

const app = express();
app.set('trust proxy', true);
app.use(json());
app.use(
  cookieSession({
    signed: false,
    secure: process.env.NODE_ENV !== 'test',
  })
);

app.use(currentUser);



app.use(indexUserRouter);
app.use(createUserRouter);
app.use(getUserRouter);
app.use(updateUserRouter);

app.all('*', async (req, res) => {
  throw new NotFoundError();
});

app.use(errorHandler);

export { app };
