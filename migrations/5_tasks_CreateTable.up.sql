
CREATE TABLE IF NOT EXISTS `tasks` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `description` TEXT NOT NULL,
  `creation_date` DATETIME NOT NULL,
  `in_work_time` VARCHAR(255),
  `author_id` INT UNSIGNED NOT NULL,
  `developer_id` INT UNSIGNED,
  `customer_id` INT UNSIGNED,
  `status_id` INT UNSIGNED NOT NULL,
  `category_id` INT UNSIGNED NOT NULL,
  `project_id` INT UNSIGNED,
  `priority` TINYINT DEFAULT(0),
  `exec_order` INT UNSIGNED DEFAULT(0),
  PRIMARY KEY `pk_id`(`id`),
  CONSTRAINT FOREIGN KEY (`author_id`) REFERENCES `users`(`id`),
  CONSTRAINT FOREIGN KEY (`developer_id`) REFERENCES `users`(`id`),
  CONSTRAINT FOREIGN KEY (`customer_id`) REFERENCES `users`(`id`),
  CONSTRAINT FOREIGN KEY (`category_id`) REFERENCES `categories`(`id`),
  CONSTRAINT FOREIGN KEY (`project_id`) REFERENCES `projects`(`id`)
);
