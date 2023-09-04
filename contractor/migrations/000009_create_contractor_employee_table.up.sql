CREATE TABLE IF NOT EXISTS contractor_employee_stats (
	id SERIAL PRIMARY KEY,
	company_uuid UUID,
	total_number_of_citizens INTEGER,
	total_number_of_non_citizens INTEGER,
	total_employees INTEGER
);

CREATE TABLE IF NOT EXISTS contractor_employee (
	id SERIAL PRIMARY KEY,
	first_name TEXT,
	middle_name TEXT,
	last_name TEXT,
	date_of_birth TIMESTAMPTZ,
	gender CHAR,
	employee_stats_id INT,
	FOREIGN KEY(employee_stats_id)
		REFERENCES contractor_employee_stats(id)
		ON DELETE CASCADE
);
