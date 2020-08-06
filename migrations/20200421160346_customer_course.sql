-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';

create table if not exists customer_course (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_customer` int(11) NOT NULL,
  `id_course` int(11) NOT NULL,
--   `completed` boolean,
  PRIMARY KEY (`id`),
  FOREIGN KEY (id_customer) REFERENCES customers(id),
  FOREIGN KEY (id_course) REFERENCES courses(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
