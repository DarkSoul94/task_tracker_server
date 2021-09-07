
CREATE TABLE IF NOT EXISTS `user_groups` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `permissions` JSON NOT NULL,
  PRIMARY KEY `pk_id`(`id`)
);

INSERT INTO `user_groups` SET
`name` = "Regular user",
`permissions` = "{}";

INSERT INTO `user_groups` SET
`name` = "Admin",
`permissions` = '
{
  \"sub_perm\": {
    \"task\": {
      \"sub_perm\": {
        \"get\": {
          \"actions_list\": [\"author\", \"dev\", \"customer\"]
        }
      }, 
        \"actions_list\": [\"create\", \"update\"]
    }, 
    \"settings\": {
      \"sub_perm\": {
        \"user\": {
          \"actions_list\": [\"update\"]
        }, 
        \"group\": {
          \"actions_list\": [\"update\", \"create\"]
        }
      }
    }
  }
}';

INSERT INTO `user_groups` SET
`name` = "PM",
`permissions` = "{}";

INSERT INTO `user_groups` SET
`name` = "Developer",
`permissions` = "{}";