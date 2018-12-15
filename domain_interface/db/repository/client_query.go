package repository

const ClientCreateTable = `
CREATE TABLE IF NOT EXISTS cp_client (
	id 				uuid PRIMARY KEY,
	created_at 		timestamptz NOT NULL,
	first_name 		text NOT NULL,
	last_name 		text NOT NULL,
	company_name 	text NOT NULL,
)`
