package session

// Base select query
const queryBase = `
SELECT
session.id,
session.user_id,
session.hashed_token,
session.client_ip,
session.user_agent,
session.created_at,
session.last_used_at
FROM session
`

// Select session by token and expiration
const queryByToken = queryBase + `
WHERE
session.hashed_token = ?
`

// Insert session statement
const stmtInsert = `
INSERT INTO session (
user_id,
hashed_token,
client_ip,
user_agent,
created_at,
last_used_at
) VALUES (
?, ?, ?, ?, ?, ?
)
`

// Update user statement.
const stmtUpdate = `
UPDATE session SET
hashed_token = ?,
client_ip = ?,
user_agent = ?,
last_used_at = ?
WHERE id = ?
`

// Delete session statement
const stmtDelete = `
DELETE FROM session
WHERE session.id=?
`
