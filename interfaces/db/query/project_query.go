package query

const ProjectTableName = "cp_project"
const ProjectTableColumns = "id, client_id, created_at, title, description"

const ProjectCreateTable = `
CREATE TABLE IF NOT EXISTS cp_project (
	id 			uuid NOT NULL,
	client_id 	uuid REFERENCES cp_client(id) ON DELETE RESTRICT,
	created_at 	timestamptz NOT NULL,
	title 		text NOT NULL,
	description text NOT NULL,
PRIMARY KEY (id,client_id)
)`

const ProjectFindRowByID = `
SELECT
	id, client_id, created_at, title, description
FROM
	cp_project
WHERE
	id = $1
LIMIT 1
`

const ProjectDeleteRowByID = `
DELETE
FROM
	cp_project
WHERE
	id = $1
`

const ProjectInsertRow = `
INSERT INTO cp_project (
	id, client_id, created_at, title, description
)
VALUES (
	$1, $2, $3, $4, $5
)
RETURNING
	id, client_id, created_at, title, description
`

const ProjectFindRowsByClientID = `
SELECT
	id, client_id, created_at, title, description
FROM
	cp_project
WHERE
	client_id = $1
`

const ProjectFindLastRowForClientByCreatedAt = `
SELECT
	id, client_id, created_at, title, description
FROM
	cp_project
WHERE
	client_id = $1
ORDER BY created_at DESC
LIMIT 1
`

const ProjectFindFirstRowForClientByCreatedAt = `
SELECT
	id, client_id, created_at, title, description
FROM
	cp_project
WHERE
	client_id = $1
ORDER BY created_at ASC
LIMIT 1
`

const ProjectFindFirstRowsForClientByCreatedAt = `
SELECT
	id, client_id, created_at, title, description
FROM
	cp_project
WHERE
	client_id = $1 AND
ORDER BY created_at ASC
LIMIT $2
`

const ProjectFindFirstAfterRowsForClientByCreatedAt = `
SELECT
	id, client_id, created_at, title, description
FROM
	cp_project
WHERE
	client_id = $1 AND
	created_at > $2::timestamptz
ORDER BY created_at ASC
LIMIT $3
`

const ProjectFindLastRowsForClientByCreatedAt = `
SELECT
	id, client_id, created_at, title, description
FROM
	cp_project
WHERE
	client_id = $1 AND
ORDER BY created_at DESC
LIMIT $2
`

const ProjectFindLastBeforeRowsForClientByCreatedAt = `
SELECT
	id, client_id, created_at, title, description
FROM
	cp_project
WHERE
	client_id = $1 AND
	created_at < $2::timestamptz
ORDER BY created_at DESC
LIMIT $3
`
