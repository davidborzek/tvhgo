package user

// Count users query.
const queryCount = `
SELECT COUNT(*) FROM "user"
`

// Select user query base.
const queryBase = `
SELECT
"user".id,
"user".username,
"user".password_hash,
"user".email,
"user".display_name,
"user".is_admin,
"user".created_at,
"user".updated_at,
two_factor_settings.enabled
FROM "user"
LEFT JOIN two_factor_settings ON "user".id = two_factor_settings.user_id
`

// Select user by id query.
const queryById = queryBase + `
WHERE
"user".id = $1
`

// Select user by username query
const queryByUsername = queryBase + `
WHERE
"user".username = $1
`

// Select user by email query
const queryByEmail = queryBase + `
WHERE
"user".email = $1
`

// Insert user statement.
const stmtInsert = `
INSERT INTO "user" (
username,
password_hash,
email,
display_name,
is_admin,
created_at,
updated_at
) VALUES (
$1, $2, $3, $4, $5, $6, $7
)
`

// Insert user statement.
const stmtInsertPostgres = `
INSERT INTO "user" (
username,
password_hash,
email,
display_name,
is_admin,
created_at,
updated_at
) VALUES (
$1, $2, $3, $4, $5, $6, $7
) RETURNING id
`

// Delete user statement.
const stmtDelete = `
DELETE FROM "user"
WHERE "user".id = $1
`

// Update user statement.
const stmtUpdate = `
UPDATE "user" SET
username = $1,
password_hash = $2,
email = $3,
display_name = $4,
is_admin = $5,
updated_at = $6
WHERE id = $7
`
