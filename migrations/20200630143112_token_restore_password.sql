-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';

create table if not exists token_restore_password (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_customer` int(11) NOT NULL,
  `token` varchar(100) NOT NULL,
  `expiration_time` varchar(100) NOT NULL,
  FOREIGN KEY (id_customer) REFERENCES customers(id),
  PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
