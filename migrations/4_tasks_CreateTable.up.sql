
CREATE TABLE IF NOT EXISTS `tasks` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `description` TEXT,
  `create_time` DATETIME,
  `in_work_time` VARCHAR(255),
  `author_id` INT UNSIGNED,
  `developer` INT UNSIGNED,
  `status_id` INT UNSIGNED,
  PRIMARY KEY `pk_id`(`id`),
  CONSTRAINT FOREIGN KEY (`author_id`) REFERENCES `users`(`id`),
  CONSTRAINT FOREIGN KEY (`developer`) REFERENCES `users`(`id`),
  CONSTRAINT FOREIGN KEY (`status_id`) REFERENCES `task_status`(`id`)
);
