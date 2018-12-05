package repository

const ProjectCreateTable = `
CREATE TABLE IF NOT EXISTS cp_project (
	id 			uuid NOT NULL,
	client_id 	uuid REFERENCES cp_client(id) ON DELETE RESTRICT,
	created_at 	timestamptz NOT NULL,
	title text 	NOT NULL,
	description text NOT NULL,
PRIMARY KEY (id,client_id)
)`

const ProjectFindByID = `
SELECT
	id, client_id, created_at, title, description
FROM
	cp_project
WHERE
	id = $1

`
