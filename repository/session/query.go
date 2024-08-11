package session

// Base select query
const queryBase = `
SELECT
"session".id,
"session".user_id,
"session".hashed_token,
"session".client_ip,
"session".user_agent,
"session".created_at,
"session".last_used_at,
"session".rotated_at
FROM "session"
`

// Select session by token and expiration
const queryByToken = queryBase + `
WHERE
"session".hashed_token = $1
`

// Select sessions by user id
const queryByUserID = queryBase + `
WHERE
"session".user_id = $1
ORDER BY "session".last_used_at DESC
`

// Insert session statement
const stmtInsert = `
INSERT INTO "session" (
user_id,
hashed_token,
client_ip,
user_agent,
created_at,
last_used_at,
rotated_at
) VALUES (
$1, $2, $3, $4, $5, $6, $7
)
`

const stmtInsertPostgres = `
INSERT INTO "session" (
user_id,
hashed_token,
client_ip,
user_agent,
created_at,
last_used_at,
rotated_at
) VALUES (
$1, $2, $3, $4, $5, $6, $7
) RETURNING id
`

// Update user statement.
const stmtUpdate = `
UPDATE "session" SET
hashed_token = $1,
client_ip = $2,
user_agent = $3,
last_used_at = $4,
rotated_at = $5
WHERE id = $6
`

// Delete session statement
const stmtDelete = `
DELETE FROM "session"
WHERE "session".id=$1 AND
"session".user_id=$2
`

// Delete expired sessions statement
const stmtDeleteExpired = `
DELETE FROM "session"
WHERE "session"."created_at" < $1
OR "session"."last_used_at" < $2
`
