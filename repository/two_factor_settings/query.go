package twofactorsettings

const queryBase = `
SELECT
two_factor_settings.user_id,
two_factor_settings.secret,
two_factor_settings.enabled,
two_factor_settings.created_at,
two_factor_settings.updated_at
FROM two_factor_settings
`

const queryByUserID = queryBase + `
WHERE two_factor_settings.user_id = $1
`

const stmtInsert = `
INSERT INTO two_factor_settings (
user_id,
secret,
enabled,
created_at,
updated_at
) VALUES (
$1, $2, $3, $4, $5
)
`
const stmtUpdate = `
UPDATE two_factor_settings SET
secret = $1,
enabled = $2,
updated_at = $3
WHERE user_id = $4
`

const stmtDelete = `
DELETE FROM two_factor_settings WHERE user_id = $1
`
