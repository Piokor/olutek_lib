CREATE TABLE books (
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