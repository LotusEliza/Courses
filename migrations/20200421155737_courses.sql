-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';

create table if not exists courses (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `description` varchar(100) NOT NULL DEFAULT '',
  `price` decimal(10.0) NOT NULL,
  `duration` varchar(100) NOT NULL DEFAULT '',
  `program` varchar(100) NOT NULL DEFAULT '',
  `img` varchar(100) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
