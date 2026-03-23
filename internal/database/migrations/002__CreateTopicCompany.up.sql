CREATE TABLE IF NOT EXISTS topic (
	id SERIAL PRIMARY KEY,
	company_id INT NOT NULL,
	name VARCHAR(255) NOT NULL,
	partitions INT NOT NULL DEFAULT 1,
	CONSTRAINT fk_company_id FOREIGN KEY (company_id) REFERENCES company(id)
);
