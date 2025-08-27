INSERT INTO messages_archive
SELECT * FROM posts
WHERE createat < NOW() - INTERVAL '2 years';

DELETE FROM posts
WHERE createat < NOW() - INTERVAL '2 years';