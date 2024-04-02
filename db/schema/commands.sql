-- You can use dbdiagram.io to create tables

CREATE TABLE IF NOT EXISTS blogs (
  id SERIAL PRIMARY KEY,
  author TEXT,
  url TEXT NOT NULL,
  title TEXT NOT NULL,
  likes INTEGER DEFAULT 0,
  -- Define one-to-many relationship between user and blogs
  user_id INTEGER REFERENCES users(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS "users" (
  id SERIAL PRIMARY KEY,
  username TEXT,
  name TEXT,
  password TEXT
);