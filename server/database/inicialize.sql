-- its for sqlite3
CREATE TABLE publ
(
    id INTEGER PRIMARY KEY,
    titulo TEXT NOT NULL,
    miniature TEXT NOT NULL,
    body TEXT NOT NULL
);
CREATE TABLE comment(
    id INTEGER PRIMARY KEY,
    id_publ INTEGER,
    username VARCHAR(32),
    body TEXT
);
CREATE TABLE users(
    id INTEGER PRIMARY KEY,
    ip_encrypted VARCHAR(64)
)
