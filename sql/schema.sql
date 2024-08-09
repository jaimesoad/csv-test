DROP TABLE IF EXISTS Users;

CREATE TABLE Users (
    Username TEXT NOT NULL,
    Identifier INTEGER PRIMARY KEY,
    One_time_password TEXT NOT NULL,
    Recovery_code TEXT NOT NULL,
    First_name TEXT NOT NULL,
    Last_name TEXT NOT NULL,
    Department TEXT NOT NULL,
    Location TEXT NOT NULL,
    Created DATETIME NOT NULL
);