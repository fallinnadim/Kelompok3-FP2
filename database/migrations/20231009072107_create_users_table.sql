-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(50) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL,
		age SMALLSERIAL NOT NULL,
		created_at DATE NOT NULL,
		updated_at DATE NOT NULL
	);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
