CREATE TABLE IF NOT EXISTS contractor_secretary(
	id SERIAL PRIMARY KEY,
	company_uuid UUID,
	full_name TEXT,
	nationality TEXT,
	appointment_date TIMESTAMPTZ,
	box_address TEXT,
	physical_address TEXT
);

CREATE TABLE IF NOT EXISTS contractor_shareholder(
	id SERIAL PRIMARY KEY,
	company_uuid UUID,
	full_name TEXT,
	nationality TEXT,
	appointment_date TIMESTAMPTZ,
	box_address TEXT,
	physical_address TEXT
);

CREATE TABLE IF NOT EXISTS contractor_director(
	id SERIAL PRIMARY KEY,
	company_uuid UUID,
	full_name TEXT,
	nationality TEXT,
	appointment_date TIMESTAMPTZ,
	box_address TEXT,
	physical_address TEXT
);
