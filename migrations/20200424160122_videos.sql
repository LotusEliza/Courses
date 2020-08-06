-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';

create table if not exists videos (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_course` int(11) NOT NULL,
  `title` varchar(100) NOT NULL DEFAULT '',
  `name` varchar(100) NOT NULL DEFAULT '',
  `sequence` int(11) NOT NULL,
  `poster` varchar(100) NOT NULL,
  FOREIGN KEY (id_course) REFERENCES courses(id),
  PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
