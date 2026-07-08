-- +goose Up
-- +goose StatementBegin
ALTER TABLE `messages` MODIFY COLUMN `content` longtext NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `messages` MODIFY COLUMN `content` text NOT NULL;
-- +goose StatementEnd
