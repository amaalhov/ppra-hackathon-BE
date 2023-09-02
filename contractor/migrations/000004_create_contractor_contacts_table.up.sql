  CREATE TABLE IF NOT EXISTS contractor_contact (
	id SERIAL PRIMARY KEY,
	full_name TEXT,
	first_name TEXT,
	middle_name TEXT,
	last_name TEXT,
	nationality TEXT,
	gov_id TEXT,
	gov_id_type TEXT,
	contact_type TEXT,
	date_of_birth TIMESTAMPTZ,
	email TEXT,
	cellphone TEXT,
	telephone TEXT,
	business_number TEXT
);
