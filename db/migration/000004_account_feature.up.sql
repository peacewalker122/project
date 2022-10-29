CREATE TABLE "accounts_follow" (
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "follow" boolean not NULL DEFAULT false,
  "follow_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "accounts_queue" (
  "from_account_id" bigint NOT NULL,
  "queue" boolean not NULL DEFAULT false,
  "to_account_id" bigint NOT NULL,
  "queue_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "accounts_follow" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("accounts_id");

ALTER TABLE "accounts_follow" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("accounts_id");

ALTER TABLE "accounts_queue" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("accounts_id");

ALTER TABLE "accounts_queue" ADD FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("accounts_id");