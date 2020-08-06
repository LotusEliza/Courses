-- +goose Up
-- +goose StatementBegin
-- SELECT 'up SQL query';

create table if not exists questions (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `id_video` int(11) NOT NULL,
  `text` varchar(100) NOT NULL DEFAULT '',
  `sequence` int(11) NOT NULL,
  FOREIGN KEY (id_video) REFERENCES videos(id),
  PRIMARY KEY (`id`)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
