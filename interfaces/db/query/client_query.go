package query

const ClientTableName = "cp_client"
const ClientTableColumns = "id, created_at, first_name, last_name, company_name"

const ClientCreateTable = `
CREATE TABLE IF NOT EXISTS cp_client (
	id 				uuid PRIMARY KEY,
	created_at 		timestamptz NOT NULL,
	first_name 		text NOT NULL,
	last_name 		text NOT NULL,
	company_name 	text NOT NULL
)`

const ClientFindRowByID = `
SELECT
	id, created_at, first_name, last_name, company_name
FROM
	cp_client
WHERE
	id = $1
LIMIT 1
`

const ClientDeleteRowByID = `
DELETE
FROM
	cp_client
WHERE
	id = $1
`

const ClientInsertRow = `
INSERT INTO cp_client (
	id, created_at, first_name, last_name, company_name
)
VALUES (
	$1, $2, $3, $4, $5
)
RETURNING
	id, created_at, first_name, last_name, company_name
`
