CREATE TABLE "account_notif"(
    "notif_id" uuid PRIMARY KEY,
    "account_id" bigint NOT NULL,
    "notif_type" varchar NOT NULL,
    "deleted" boolean NOT NULL DEFAULT false,
    "notif_title" varchar NOT NULL,
    "notif_content" varchar NOT NULL,
    "notif_time" timestamptz DEFAULT (now()),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "account_notif_read"(
    "notif_id" uuid PRIMARY KEY,
    "account_id" bigint NOT NULL,
    "read_at" timestamptz
);