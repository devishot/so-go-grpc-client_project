package repository

const ProjectTableName = "cp_project"
const ProjectFields = "id, client_id, created_at, title, description"

const ProjectCreateTable = `
CREATE TABLE IF NOT EXISTS cp_project (
	id 			uuid NOT NULL,
	client_id 	uuid REFERENCES cp_client(id) ON DELETE RESTRICT,
	created_at 	timestamptz NOT NULL,
	title text 	NOT NULL,
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
	id, description
`
