package token

const queryBase = `
SELECT
token.id,
token.user_id,
token.name,
token.hashed_token,
token.created_at,
token.updated_at
FROM token
`

const queryByToken = queryBase + `
WHERE token.hashed_token = $1
`

const queryByUser = queryBase + `
WHERE token.user_id = $1
`

const stmtInsert = `
INSERT INTO token (
user_id,
hashed_token,
name,
created_at,
updated_at
) VALUES (
$1,
$2,
$3,
$4,
$5
)
`

const stmtInsertPostgres = `
INSERT INTO token (
user_id,
hashed_token,
name,
created_at,
updated_at
) VALUES (
$1,
$2,
$3,
$4,
$5
) RETURNING id
`

const stmtDelete = `
DELETE FROM token WHERE id = $1
`
