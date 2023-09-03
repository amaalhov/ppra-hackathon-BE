CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS contractor (
	uid uuid DEFAULT uuid_generate_V4(),
	name TEXT,
	business_type TEXT,
	ownership_type TEXT,
	cipa_uin TEXT,
	national_id TEXT,
	is_registered_with_cipa BOOL,
	registration_date TIMESTAMPTZ
);
