CREATE TABLE IF NOT EXISTS "account_notifs"(
    "id" uuid PRIMARY KEY,
    "account_id" bigint,
    "notif_type" varchar NOT NULL,
    "is_read" boolean NOT NULL DEFAULT false,
    "notif_title" varchar,
    "notif_content" varchar,
    "notif_time" timestamptz DEFAULT (now()),
    "created_at" timestamptz NOT NULL DEFAULT (now())
);