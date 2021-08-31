
CREATE TABLE IF NOT EXISTS `task_track` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `task_id` INT UNSIGNED NOT NULL,
  `user_id` INT UNSIGNED NOT NULL,
  `start_time` DATETIME,
  `end_time` DATETIME,
  `difference` INT UNSIGNED,
  PRIMARY KEY `pk_id`(`id`),
  CONSTRAINT FOREIGN KEY (`task_id`) REFERENCES `tasks`(`id`),
  CONSTRAINT FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
)
