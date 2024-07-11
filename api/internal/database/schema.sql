CREATE TABLE books (
    id SERIAL UNIQUE,
    api_id CHAR(12) NOT NULL UNIQUE,
    title VARCHAR(512) NOT NULL,
	authors VARCHAR(256)[],
    subtitle VARCHAR(512),
    page_count SMALLINT,
	publish_date VARCHAR(64),
	description TEXT,
    publisher VARCHAR(256),
	language VARCHAR(128),
	small_thumbnail VARCHAR(512),
	thumbnail VARCHAR(512)
)

CREATE TABLE users (
    id SERIAL UNIQUE,
    username VARCHAR(128) NOT NULL UNIQUE,
    password CHAR(64) NOT NULL
)

CREATE TABLE session (
    id UUID UNIQUE,
    user_id INTEGER REFERENCES users(id) NOT NULL UNIQUE,
    expires_at TIMESTAMP NOT NULL, 
    data TEXT
)