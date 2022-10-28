CREATE TABLE "like_feature" (
  "from_account_id" bigint NOT NULL,
  "is_like" boolean NOT NULL DEFAULT false,
  "post_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "comment_feature" (
  "comment_id" bigserial NOT NULL,
  "from_account_id" bigint NOT NULL,
  "comment" varchar NOT NULL,
  "sum_like" bigint NOT NULL DEFAULT 0,
  "post_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "retweet_feature" (
  "from_account_id" bigint NOT NULL,
  "retweet" boolean NOT NULL DEFAULT false,
  "post_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "qoute_retweet_feature" (
  "from_account_id" bigint NOT NULL,
  "qoute_retweet" boolean NOT NULL DEFAULT false,
  "qoute" varchar NOT NULL,
  "post_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "post_feature" (
  "post_id" bigint PRIMARY KEY,
  "sum_comment" bigint NOT NULL DEFAULT 0,
  "sum_like" bigint NOT NULL DEFAULT 0,
  "sum_retweet" bigint NOT NULL DEFAULT 0,
  "sum_qoute_retweet" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "like_feature" ("created_at");
CREATE INDEX ON "comment_feature" ("created_at");
CREATE INDEX ON "retweet_feature" ("created_at");
CREATE INDEX ON "like_feature" ("created_at");

ALTER TABLE "like_feature" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("accounts_id");

ALTER TABLE "like_feature" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "comment_feature" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("accounts_id");

ALTER TABLE "comment_feature" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "retweet_feature" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("accounts_id");

ALTER TABLE "retweet_feature" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "qoute_retweet_feature" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("accounts_id");

ALTER TABLE "qoute_retweet_feature" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "post_feature" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");