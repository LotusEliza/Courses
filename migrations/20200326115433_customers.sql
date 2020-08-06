-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';

create table if not exists customers (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '',
  `email` varchar(100) NOT NULL DEFAULT '',
  `tel` varchar(100) NOT NULL DEFAULT '',
  `date_create` varchar(100) NOT NULL DEFAULT '',
  `password` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
