CREATE TABLE "like_feature" (
  "from_account_id" bigint NOT NULL,
  "is_like" boolean NOT NULL DEFAULT false,
  "post_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "comment_feature" (
  "comment_id" uuid NOT NULL,
  "from_account_id" bigint NOT NULL,
  "comment" varchar NOT NULL,
  "sum_like" bigint NOT NULL DEFAULT 0,
  "post_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "retweet_feature" (
  "from_account_id" bigint NOT NULL,
  "retweet" boolean NOT NULL DEFAULT false,
  "post_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "qoute_retweet_feature" (
  "from_account_id" bigint NOT NULL,
  "qoute_retweet" boolean NOT NULL DEFAULT false,
  "qoute" varchar NOT NULL,
  "post_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "post_feature" (
  "post_id" uuid PRIMARY KEY,
  "sum_comment" bigint NOT NULL DEFAULT 0,
  "sum_like" bigint NOT NULL DEFAULT 0,
  "sum_retweet" bigint NOT NULL DEFAULT 0,
  "sum_qoute_retweet" bigint NOT NULL DEFAULT 0,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "entries" (
  "entries_id" bigserial PRIMARY KEY,
  "from_account_id" bigint NOT NULL,
  "to_account_id" bigint NOT NULL,
  "post_id" uuid NOT NULL,
  "type_entries" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON like_feature ("created_at" DESC);
CREATE INDEX ON "comment_feature" ("comment_id");
CREATE INDEX ON "retweet_feature" ("created_at" DESC);
CREATE INDEX ON "like_feature" ("created_at" DESC);
CREATE INDEX ON "entries" ("post_id");

ALTER TABLE "like_feature" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("accounts_id");

ALTER TABLE "like_feature" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "comment_feature" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("accounts_id");

ALTER TABLE "comment_feature" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "retweet_feature" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("accounts_id");

ALTER TABLE "retweet_feature" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "qoute_retweet_feature" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("accounts_id");

ALTER TABLE "qoute_retweet_feature" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "post_feature" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("accounts_id");

ALTER TABLE "entries" ADD FOREIGN KEY ("post_id") REFERENCES "post" ("post_id");