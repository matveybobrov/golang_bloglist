CREATE TABLE blogs (
  id SERIAL UNIQUE,
  author TEXT,
  url TEXT NOT NULL,
  title TEXT NOT NULL,
  likes INT DEFAULT 0
);

INSERT INTO blogs
(author, url, title) VALUES
('John Doe', 'goolge.com', 'Hello world'),
('Ivan Dorn', 'yandex.ru', 'Goodbye world');