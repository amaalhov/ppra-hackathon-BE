CREATE TABLE IF NOT EXISTS contractor_bank_detail (
	id SERIAL PRIMARY KEY,
	company_uuid UUID,
	bank_name TEXT,
	branch TEXT,
	branch_address TEXT,
	account_number TEXT,
	account_type TEXT
);
