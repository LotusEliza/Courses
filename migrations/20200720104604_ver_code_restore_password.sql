-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

create table if not exists ver_code_restore_password (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_customer` int(11) NOT NULL,
  `code` int(11) NOT NULL,
  `expiration_time` varchar(100) NOT NULL,
  FOREIGN KEY (id_customer) REFERENCES customers(id),
  PRIMARY KEY (`id`)
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
