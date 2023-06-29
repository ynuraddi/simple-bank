-- name: CreateUser :one
insert into users (
    username,
    hashed_password,
    full_name,
    email
) values (
    $1, $2, $3, $4
) returning *;

-- name: GetUser :one
select * from users
where username = $1 limit 1;

-- name: UpdateUser :one
UPDATE users
SET
    hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    email = COALESCE(sqlc.narg(email), email)
WHERE
    username = sqlc.arg(username)
RETURNING *;