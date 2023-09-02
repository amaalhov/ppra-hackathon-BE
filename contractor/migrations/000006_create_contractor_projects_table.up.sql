CREATE TABLE IF NOT EXISTS contractor_project (
	id SERIAL PRIMARY KEY,
	name TEXT,
	description TEXT,
	client_name TEXT,
	client_representative,
	client_contact_number
);
