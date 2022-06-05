CREATE TABLE users (
    id                      SERIAL PRIMARY KEY,
    username                VARCHAR(100),
    password                VARCHAR(1000)
);
CREATE TABLE habit_lists (
    id                      SERIAL PRIMARY KEY,
    owner_id                INT,
    habit_name              VARCHAR(100) UNIQUE,
    icon_url                TEXT,
    color                   VARCHAR(30),
    default_repeat_count    INT,
    FOREIGN KEY(owner_id)   REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE habits (
    id                      SERIAL PRIMARY KEY,
    owner_id                INT,
    habit_name              VARCHAR(100),
    date_created            DATE,
    comment                 TEXT,
    target_repeat_count     INT,
    repeat_count            INT,
    FOREIGN KEY(owner_id)   REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(habit_name) REFERENCES habit_lists(habit_name) ON DELETE CASCADE
);

