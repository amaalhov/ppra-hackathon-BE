import express from 'express';
import 'express-async-errors';
import { json } from 'body-parser';
import cookieSession from 'cookie-session';
import { errorHandler, NotFoundError, currentUser } from '@le-ma/common';
import { createApplicationRouter } from './routes/create-application';
import { showApplicationRouter } from './routes/get-application';
import { indexApplicationRouter } from './routes/index';
import { updateApplicationRouter } from './routes/update-application';

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

app.use(createApplicationRouter);
app.use(showApplicationRouter);
app.use(indexApplicationRouter);
app.use(updateApplicationRouter);

app.all('*', async (req, res) => {
  throw new NotFoundError();
});

app.use(errorHandler);

export { app };
