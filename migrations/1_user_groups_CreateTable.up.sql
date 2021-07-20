
CREATE TABLE IF NOT EXISTS `user_groups` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `group_name` VARCHAR(255) NOT NULL,
  PRIMARY KEY `pk_id`(`id`)
);

INSERT INTO `user_groups` SET
group_name = "Regular user";

INSERT INTO `user_groups` SET
group_name = "Admin";

INSERT INTO `user_groups` SET
group_name = "PM";

INSERT INTO `user_groups` SET
group_name = "Developer";