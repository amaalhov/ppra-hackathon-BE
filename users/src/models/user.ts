import mongoose from 'mongoose';

interface UserAttrs {
  userId: string;
  fullnames: string;
  idno: string;
  email: string;
  role: string;
  idcard: string;
  verified: string;
  status: string;
}

interface UserDoc extends mongoose.Document {
  userId: string;
  fullnames: string;
  idno: string;
  email: string;
  role: string;
  idcard: string;
  verified: string;
  status: string;
}

interface UsersModel extends mongoose.Model<UserDoc> {
  build(attrs: UserAttrs): UserDoc;
}

const userSchema = new mongoose.Schema(
  {
    userId: {
      type: String,
      required: true,
    },
    fullnames: {
      type: String,
      required: true,
    },
    idno: {
      type: String,
      required: true,
    },
    email: {
      type: String,
      required: true,
    },
    role: {
      type: String,
      required: true,
    },
    idcard: {
      type: String,
      required: true,
    },
    verified: {
      type: String,
      required: true,
    },
    status: {
      type: String,
      required: true,
    },
  },
  {
    toJSON: {
      transform(doc, ret) {
        ret.id = ret._id;
        delete ret._id;
      },
    },
  }
);

userSchema.statics.build = (attrs: UserAttrs) => {
  return new User(attrs);
};

const User = mongoose.model<UserDoc, UsersModel>('User', userSchema);

export { User };
