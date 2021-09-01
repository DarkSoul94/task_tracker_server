package task

/**
 *
 * @api {GET} /task_tracker/tasks Получение списка задач
 * @apiName GetTasksList
 * @apiGroup 02 Задачи
 * @apiVersion  0.0.1
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 *
 * @apiSuccess (Success 200) {Uint64} 		id		 					ИД задачи
 * @apiSuccess (Success 200) {String} 		name		 				Название задачи
 * @apiSuccess (Success 200) {String} 		description		 	Описание задачи
 * @apiSuccess (Success 200) {String} 		creation_date		Дата создания задачи
 * @apiSuccess (Success 200) {User} 			author		 			Автор задачи
 * @apiSuccess (Success 200) {Uint64} 		author.id		 		ИД автора задачи
 * @apiSuccess (Success 200) {String} 		author.name		 	Имя автора
 * @apiSuccess (Success 200) {User} 			developer		 		Разработчик
 * @apiSuccess (Success 200) {Uint64} 		developer.id		ИД разработчика
 * @apiSuccess (Success 200) {String} 		developer.name	Имя разработчика
 * @apiSuccess (Success 200) {User} 			customer		 		Заказчик задачи
 * @apiSuccess (Success 200) {Uint64} 		customer.id		 	ИД заказчика
 * @apiSuccess (Success 200) {String} 		customer.name		Имя заказчика
 * @apiSuccess (Success 200) {Category} 	category		 		Категория задачи
 * @apiSuccess (Success 200) {Uint64} 		category.id		 	ИД категории
 * @apiSuccess (Success 200) {String} 		category.name		Название категории
 * @apiSuccess (Success 200) {Project} 		project		 			Проект к которому относится задача
 * @apiSuccess (Success 200) {Uint64} 		project.id		 	ИД проекта
 * @apiSuccess (Success 200) {String} 		project.name		Название проекта
 * @apiSuccess (Success 200) {Bool} 			priority		 		Приоритет задачи (`true` - экстренный, `false` - стандартный )
 * @apiSuccess (Success 200) {Uint64} 		exec_order		 	Порядок выполнения задачи
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *     {
  "status": "success",
  "data": [
    {
      "id": 1,
      "name": "test",
      "description": "test",
      "creation_date": "2021-08-31 11:21:13 +0300 EEST",
      "author": {
        "id": 1,
        "name": "Евгений Николаевич Табаков"
      },
      "developer": {
        "id": 2,
        "name": "Вячеслав Викторович Тищенко"
      },
      "customer": {
        "id": 1,
        "name": "Евгений Николаевич Табаков"
      },
      "category": {
        "id": 1,
        "name": "test"
      },
      "project": null,
      "priority": true,
      "exec_order": 0
    },
    {
      "id": 2,
      "name": "banner",
      "description": "banner",
      "creation_date": "2021-08-31 11:21:28 +0300 EEST",
      "author": {
        "id": 1,
        "name": "Евгений Николаевич Табаков"
      },
      "developer": null,
      "customer": null,
      "category": {
        "id": 1,
        "name": "test"
      },
      "project": null,
      "priority": false,
      "exec_order": 0
    }
  ]
}
 * }
 *
 *
*/
