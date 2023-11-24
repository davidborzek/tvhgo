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
WHERE token.hashed_token = @hashed_token
`

const queryByUser = queryBase + `
WHERE token.user_id = @user_id
`

const stmtInsert = `
INSERT INTO token (
user_id,
hashed_token,
name,
created_at,
updated_at
) VALUES (
@user_id,
@hashed_token,
@name,
@created_at,
@updated_at
)
`

const stmtDelete = `
DELETE FROM token WHERE id = @id
`
