

CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE IF NOT EXISTS "Users" (
        "mail"  TEXT NOT NULL UNIQUE,
        "username"      TEXT NOT NULL UNIQUE,
        "psw"   TEXT NOT NULL,
        "id"    INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,
        "code"  TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS "Session" (
        "UserId"        INTEGER NOT NULL,
        "Uuid"  TEXT NOT NULL UNIQUE,
        "Time"  timestamp
);
CREATE TABLE IF NOT EXISTS "ThreadStats" (
        "FromUserID"    INTEGER NOT NULL,
        "ToThreadID"    INTEGER NOT NULL,
        "Value" INTEGER
);
CREATE TABLE IF NOT EXISTS "Images" (
        "Path"  TEXT NOT NULL UNIQUE,
        "IsUser"        INTEGER NOT NULL,
        "ID"    INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS "Thread" (
        "Title" TEXT,
        "UserID"        INTEGER NOT NULL,
        "Likes" INTEGER,
        "Dislikes"      INTEGER,
        "ThreadID"      INTEGER PRIMARY KEY AUTOINCREMENT,
        "ToThreadID"    TEXT,
        "Date"  TEXT NOT NULL,
        "Content"       TEXT NOT NULL,
        "Category"      TEXT NOT NULL,
        "Username"      TEXT,
        "Image" TEXT NOT NULL
);