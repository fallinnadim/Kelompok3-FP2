-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS social_medias (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50) NOT NULL,
		social_media_url VARCHAR(50) NOT NULL,
		user_id SERIAL NOT NULL,
		created_at DATE NOT NULL,
        updated_at DATE NOT NULL,
		CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id)
	);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE social_medias;
-- +goose StatementEnd
