-- name: GetContacts :many
SELECT id, name, lastname, email, phone, create_at, updated_at FROM contact;
    
-- name: GetContactByID :one
SELECT id, name, lastname, email, phone, create_at, updated_at FROM contact
WHERE id = $1;

-- name: CreateContact :one
INSERT INTO contact (id, name, lastname, email, phone) values ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateContact :one
UPDATE contact SET name = $2, lastname = $3, email = $4, phone = $5, updated_at=now()
WHERE id = $1 and updated_at=$6
RETURNING *;

-- name: DeleteContact :exec
DELETE FROM contact WHERE id = $1;
