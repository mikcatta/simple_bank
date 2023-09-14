-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
    username,
    email,
    secret_code
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateVerifyEmail :one
UPDATE verify_emails 
SET 
    is_used = True
WHERE 
    id = @id And
    secret_code = @secret_code And
    is_used = false And
    epired_at > now()
RETURNING *;
