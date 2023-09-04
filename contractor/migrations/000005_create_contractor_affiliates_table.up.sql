CREATE TABLE IF NOT EXISTS contractor_affiliate (
	id SERIAL PRIMARY KEY,  
	company_uuid UUID,
	full_name TEXT,
	address TEXT,
	attachment_url TEXT
);
