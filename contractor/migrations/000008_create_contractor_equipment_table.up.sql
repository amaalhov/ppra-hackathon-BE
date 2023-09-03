CREATE TABLE IF NOT EXISTS contractor_equipment (
	id SERIAL PRIMARY KEY,
	value_of_assets DOUBLE PRECISION,
	value_of_equipment DOUBLE PRECISION,
	paid_up_capital DOUBLE PRECISION
);

CREATE TABLE IF NOT EXISTS contractor_vehicle(
	id SERIAL PRIMARY KEY,
	registered_owner TEXT,
	ownership TEXT,
	registration_number TEXT,
	date_of_registration TIMESTAMPTZ,
	vehicle_model TEXT,
	FOREIGN KEY(equipment_id)
		REFERENCES contractor_equiment(id)
		ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS contractor_plant(
	id SERIAL PRIMARY KEY,
	registered_owner TEXT,
	ownership TEXT,
	description TEXT,
	registration_number TEXT,
	date_of_purchase TIMESTAMPTZ,
	FOREIGN KEY(equipment_id)
		REFERENCES contractor_equiment(id)
		ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS contractor_property(
	id SERIAL PRIMARY KEY,
	ownership TEXT,
	present_value DOUBLE PRECISION,
	attachment_url TEXT,
	locality TEXT,
	FOREIGN KEY(equipment_id)
		REFERENCES contractor_equiment(id)
		ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS contractor_office_equipment(
	id SERIAL PRIMARY KEY,
	office_equipment TEXT,
	present_value DOUBLE PRECISION
	attachment_url TEXT,
	FOREIGN KEY(equipment_id)
		REFERENCES contractor_equiment(id)
		ON DELETE CASCADE
);
