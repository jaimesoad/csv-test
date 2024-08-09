-- name: GetAllUsers :many
SELECT * FROM Users;

-- name: NewUser :exec
INSERT INTO Users (
    Username,
    Identifier,
    One_time_password,
    Recovery_code,
    First_name,
    Last_name,
    Department,
    Location,
    Created
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);