import { z } from "zod";

const schema = z.object({});

const data = {};

export const env = schema.parse(data);
