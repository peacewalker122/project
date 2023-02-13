ALTER TABLE account_notifs
    ADD COLUMN IF NOT EXISTS "username" VARCHAR(255);