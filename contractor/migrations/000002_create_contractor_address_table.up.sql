CREATE TABLE IF NOT EXISTS contractor_address (
	id SERIAL PRIMARY KEY,
	company_uuid UUID,
	country TEXT,
	district_name TEXT,
	town TEXT,
	plot_number TEXT,
	street TEXT
)
