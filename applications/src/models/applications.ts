import mongoose from 'mongoose';
import { updateIfCurrentPlugin } from 'mongoose-update-if-current';

const subcodeSchema = new mongoose.Schema({
  subcode: {
    type: String,
    required: true,
  },
  description: {
    type: String,
    required: true,
  },
});

interface ApplicationAttrs {
  userId: string;
  orderId: string;
  descipline: string;
  code: string;
  description: string;
  subcodes: Array<{
    subcode: string;
    description: string;
  }>;
}

interface ApplicationDoc extends mongoose.Document {
  userId: string;
  orderId: string;
  descipline: string;
  code: string;
  description: string;
  subcodes: Array<{
    subcode: string;
    description: string;
  }>;
}

interface ApplicationModel extends mongoose.Model<ApplicationDoc> {
  build(attrs: ApplicationAttrs): ApplicationDoc;
}

const applicationSchema = new mongoose.Schema(
  {
    userId: {
      type: String,
      required: true,
    },
    orderId: {
      type: String,
      required: true,
    },
    descipline: {
      type: String,
      required: true,
    },
    code: {
      type: String,
      required: true,
    },
    description: {
      type: String,
      required: true,
    },
    subcodes: [subcodeSchema],
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
applicationSchema.set('versionKey', 'version');
applicationSchema.plugin(updateIfCurrentPlugin);

applicationSchema.statics.build = (attrs: ApplicationAttrs) => {
  return new Application(attrs);
};

const Application = mongoose.model<ApplicationDoc, ApplicationModel>(
  'Application',
  applicationSchema
);

export { Application };

// [
// {
//   "Discipline":"Service",
//   "Code": "120",
//   "Description": "ICT Technical Support Services",
//   "Sub-codes": [
//   { "Sub-code": "01", "Description": "Systems Development Services and maintenance services" },
//   { "Sub-code": "02", "Description": "Server Management and maintenance Services" },
//   { "Sub-code": "03", "Description": "Data center maintenance & hosting facilities" },
//   { "Sub-code": "04", "Description": "Desktop Management and maintenance Services" },
//   { "Sub-code": "05", "Description": "Network Management and maintenance Services" },
//   { "Sub-code": "06", "Description": "ICT Security Management and maintenance Services" },
//   { "Sub-code": "07", "Description": "Internet Services" },
//   { "Sub-code": "08", "Description": "ICT Risk Management Services" },
//   { "Sub-code": "09", "Description": "Imaging, data capture and migration services" }
//   ]
// },
// {
//   "Discipline":"Service",
//   "Code": "120",
//   "Description": "ICT Technical Support Services",
//   "Sub-codes": [
//   { "Sub-code": "01", "Description": "Systems Development Services and maintenance services" },
//   { "Sub-code": "02", "Description": "Server Management and maintenance Services" },
//   { "Sub-code": "03", "Description": "Data center maintenance & hosting facilities" },
//   { "Sub-code": "04", "Description": "Desktop Management and maintenance Services" },
//   { "Sub-code": "05", "Description": "Network Management and maintenance Services" },
//   { "Sub-code": "06", "Description": "ICT Security Management and maintenance Services" },
//   { "Sub-code": "07", "Description": "Internet Services" },
//   { "Sub-code": "08", "Description": "ICT Risk Management Services" },
//   { "Sub-code": "09", "Description": "Imaging, data capture and migration services" }
//   ]
// },
//  ]
