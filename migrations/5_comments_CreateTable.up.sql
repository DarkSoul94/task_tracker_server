
CREATE TABLE IF NOT EXISTS `comments` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `text` TEXT NOT NULL,
  `create_time` DATETIME,
  `author_id` INT UNSIGNED,
  `task_id` INT UNSIGNED,
  PRIMARY KEY `pk_id`(`id`),
  CONSTRAINT FOREIGN KEY (`author_id`) REFERENCES `users`(`id`),
  CONSTRAINT FOREIGN KEY (`task_id`) REFERENCES `tasks`(`id`)
);
