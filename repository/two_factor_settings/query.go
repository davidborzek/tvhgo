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
WHERE two_factor_settings.user_id = ?
`

const stmtInsert = `
INSERT INTO two_factor_settings (
user_id,
secret,
enabled,
created_at,
updated_at
) VALUES (
?, ?, ?, ?, ?
)
`

const stmtDelete = `
DELETE FROM two_factor_settings WHERE user_id = ?
`

const stmtUpdateEnabled = `
UPDATE two_factor_settings SET
enabled = ?
WHERE user_id = ?
`
