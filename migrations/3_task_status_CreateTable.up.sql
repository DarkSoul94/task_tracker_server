
CREATE TABLE IF NOT EXISTS `task_status` (
  `id` INT UNSIGNED NOT NULL,
  `name` VARCHAR(255) NOT NULL,
  PRIMARY KEY `pk_id`(`id`)
);

INSERT INTO task_status SET
id = 1,
name = "Новая";

INSERT INTO task_status SET
id = 2,
name = "В доработке";

INSERT INTO task_status SET
id = 3,
name = "В работе";

INSERT INTO task_status SET
id = 4,
name = "Отложена";

INSERT INTO task_status SET
id = 5,
name = "Выполнена";

INSERT INTO task_status SET
id = 6,
name = "Отклонена";