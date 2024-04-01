-- You can use dbdiagram.io to create tables

CREATE TABLE IF NOT EXISTS blogs (
  id SERIAL UNIQUE,
  author TEXT,
  url TEXT NOT NULL,
  title TEXT NOT NULL,
  likes INT DEFAULT 0
  --user_id INT FOREIGN KEY
);

CREATE TABLE IF NOT EXISTS "users" (
  id SERIAL PRIMARY KEY,
  username TEXT,
  name TEXT,
  password TEXT
);

-- Define one-to-many relationship between user and blogs
ALTER TABLE blogs ADD user_id INTEGER REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL;