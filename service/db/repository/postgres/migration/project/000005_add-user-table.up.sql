CREATE TABLE tokens (
    email varchar(255) NOT NULL,
	access_token varchar(255) NOT NULL,
	refresh_token varchar(255),
	token_type varchar(255),
	expiry timestamp with time zone,
	raw jsonb
);

ALTER TABLE tokens
ADD FOREIGN KEY (email) REFERENCES users(email);