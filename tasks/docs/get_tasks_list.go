package docs

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
 * @apiSuccess (Success 200) {String} 		status		 					  Статус ответа на запрос
 * @apiSuccess (Success 200) {Data} 		  data		 					    Объект с данными по запросу
 * @apiSuccess (Success 200) {Uint64} 		data.id		 					  ИД задачи
 * @apiSuccess (Success 200) {String} 		data.name		 				  Название задачи
 * @apiSuccess (Success 200) {String} 		data.description		 	Описание задачи
 * @apiSuccess (Success 200) {String} 		data.creation_date		Дата создания задачи
 * @apiSuccess (Success 200) {String} 		data.in_work_time 		Время проведенное в работе в `секундах`
 * @apiSuccess (Success 200) {Status} 		data.status        		Статус задачи
 * @apiSuccess (Success 200) {Uint64} 		data.status.id        ИД статуса
 * @apiSuccess (Success 200) {String} 		data.status.name      Название статус
 * @apiSuccess (Success 200) {User} 			data.author		 			  Автор задачи
 * @apiSuccess (Success 200) {Uint64} 		data.author.id		 		ИД автора задачи
 * @apiSuccess (Success 200) {String} 		data.author.name		 	Имя автора
 * @apiSuccess (Success 200) {User} 			data.developer		 		Разработчик
 * @apiSuccess (Success 200) {Uint64} 		data.developer.id		  ИД разработчика
 * @apiSuccess (Success 200) {String} 		data.developer.name	  Имя разработчика
 * @apiSuccess (Success 200) {User} 			data.customer		 		  Заказчик задачи
 * @apiSuccess (Success 200) {Uint64} 		data.customer.id		 	ИД заказчика
 * @apiSuccess (Success 200) {String} 		data.customer.name		Имя заказчика
 * @apiSuccess (Success 200) {Category} 	data.category		 		  Категория задачи
 * @apiSuccess (Success 200) {Uint64} 		data.category.id		 	ИД категории
 * @apiSuccess (Success 200) {String} 		data.category.name		Название категории
 * @apiSuccess (Success 200) {Project} 		data.project		 			Проект к которому относится задача
 * @apiSuccess (Success 200) {Uint64} 		data.project.id		 	  ИД проекта
 * @apiSuccess (Success 200) {String} 		data.project.name		  Название проекта
 * @apiSuccess (Success 200) {Bool} 			data.priority		 		  Приоритет задачи (`true` - экстренный, `false` - стандартный )
 * @apiSuccess (Success 200) {Uint64} 		data.exec_order		 	  Порядок выполнения задачи
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
