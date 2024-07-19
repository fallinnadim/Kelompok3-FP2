-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS photos (
		id SERIAL PRIMARY KEY,
		title VARCHAR(100) NOT NULL,
		caption VARCHAR(255) NOT NULL,
		photo_url VARCHAR(255) NOT NULL,
		user_id SERIAL NOT NULL,
        created_at DATE NOT NULL,
        updated_at DATE NOT NULL,
		CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id)
	);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE photos;
-- +goose StatementEnd