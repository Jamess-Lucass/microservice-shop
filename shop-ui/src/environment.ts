import { z } from "zod";

const schema = z.object({
  CATALOG_SERVICE_BASE_URL: z.string().url().optional().or(z.literal("")),
  BASKET_SERVICE_BASE_URL: z.string().url().optional().or(z.literal("")),
  GOOGLE_CLIENT_ID: z.string().optional().or(z.literal("")),
  GOOGLE_CLIENT_SECRET: z.string().optional().or(z.literal("")),
});

const data = {
  CATALOG_SERVICE_BASE_URL: process.env.NEXT_PUBLIC_CATALOG_SERVICE_BASE_URL,
  BASKET_SERVICE_BASE_URL: process.env.NEXT_PUBLIC_BASKET_SERVICE_BASE_URL,
  GOOGLE_CLIENT_ID: process.env.GOOGLE_CLIENT_ID,
  GOOGLE_CLIENT_SECRET: process.env.GOOGLE_CLIENT_SECRET,
};

export const env = schema.parse(data);
