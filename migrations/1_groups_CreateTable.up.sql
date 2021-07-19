
CREATE TABLE IF NOT EXISTS `groups` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `group_name` VARCHAR(255) NOT NULL,
  PRIMARY KEY `pk_id`(`id`)
);

INSERT INTO `groups` SET
group_name = "Admin";

INSERT INTO `groups` SET
group_name = "PM";

INSERT INTO `groups` SET
group_name = "Developer";