
CREATE TABLE IF NOT EXISTS `users` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL UNIQUE,
  `group_id` INT UNSIGNED,
  PRIMARY KEY `pk_id`(`id`),
  CONSTRAINT FOREIGN KEY (`group_id`) REFERENCES `groups`(`id`)
);
