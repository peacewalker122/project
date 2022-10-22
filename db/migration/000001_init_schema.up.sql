CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "accounts" (
  "accounts_id" bigserial PRIMARY KEY,
  "owner" varchar NOT NULL,
  "is_private" boolean NOT NULL DEFAULT true,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
-- considered to add picture id in post.
CREATE TABLE "post" (
  "post_id" bigserial PRIMARY KEY,
  "account_id" bigint NOT NULL,
  "picture_description" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "entries_id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "post_id" bigint NOT NULL,
  "type_entries" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "accounts" ("owner");

CREATE INDEX ON "post" ("account_id");

CREATE INDEX ON "entries" ("post_id");

ALTER TABLE "accounts" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

ALTER TABLE "post" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("accounts_id");

ALTER TABLE "accounts" ADD CONSTRAINT "owner_currency_key" UNIQUE ("owner");
-- This Code Above To Make Consistent Account Name.

ALTER TABLE "entries" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("accounts_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");