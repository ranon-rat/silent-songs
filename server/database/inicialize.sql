-- its for sqlite3
CREATE TABLE publ
(
    id INTEGER PRIMARY KEY,
    titulo TEXT NOT NULL,
    mineatura TEXT NOT NULL,
    body TEXT NOT NULL
);
CREATE TABLE users(
    id INTEGER PRIMARY KEY,
    ip_encrypted VARCHAR(64)
)
