CREATE TABLE users (
    id                      SERIAL PRIMARY KEY,
    username                VARCHAR(100) UNIQUE,
    password                VARCHAR(1000)
);
CREATE TABLE habit_lists (
    id                      SERIAL PRIMARY KEY,
    owner_id                INT,
    habit_name              VARCHAR(100),
    icon_url                TEXT,
    color                   VARCHAR(30),
    default_repeat_count    INT,
    FOREIGN KEY(owner_id)   REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE habits (
    id                          SERIAL PRIMARY KEY,
    owner_id                    INT,
    habit_list_id               INT,
    habit_name                  VARCHAR(100),
    date_created                DATE,
    comment                     TEXT,
    target_repeat_count         INT,
    repeat_count                INT,
    FOREIGN KEY(owner_id)       REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(habit_list_id)  REFERENCES habit_lists(id) ON DELETE CASCADE ON UPDATE CASCADE
);

