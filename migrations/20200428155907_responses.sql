-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';

create table if not exists responses (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_question` int(11) NOT NULL,
  `text` varchar(100) NOT NULL DEFAULT '',
  `correct` boolean,
  FOREIGN KEY (id_question) REFERENCES questions(id),
  PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
