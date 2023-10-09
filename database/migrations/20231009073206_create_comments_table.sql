-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS comments (
		id SERIAL PRIMARY KEY,
        user_id SERIAL NOT NULL,
        photo_id SERIAL NOT NULL,
		message VARCHAR(255) NOT NULL,		
        created_at DATE NOT NULL,
        updated_at DATE NOT NULL,
		CONSTRAINT fk_user_id FOREIGN KEY(user_id) REFERENCES users(id),
        CONSTRAINT fk_photo_id FOREIGN KEY(photo_id) REFERENCES photos(id)
	);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE comments;
-- +goose StatementEnd
