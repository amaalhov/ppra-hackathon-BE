CREATE TABLE IF NOT EXISTS contractor_project (
	id SERIAL PRIMARY KEY,
	company_uuid UUID,
	name TEXT,
	description TEXT,
	client_name TEXT,
	client_representative TEXT,
	client_contact_number TEXT
);
