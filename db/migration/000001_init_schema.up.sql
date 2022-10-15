CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "account_type" boolean NOT NULL DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "post" (
  "id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "picture_description" varchar DEFAULT null,
  "picture_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "post" ("account_id");

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "post" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner");
-- This Code Above To Make Consistent Account Name.