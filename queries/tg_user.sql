
-- name: CreateUser :one
INSERT INTO users (
  telegram_id, first_name, last_name, username, language, user_type
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id;


-- name: GetUserByTelegramID :one
SELECT * FROM users
WHERE telegram_id = $1 LIMIT 1;

-- name: CheckUserExist :one
SELECT EXISTS(
  SELECT 1 FROM users
  WHERE telegram_id = $1
) AS exist;

-- name: UpdateUserType :exec
UPDATE users
SET user_type = $2
WHERE telegram_id = $1
RETURNING id;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC;

