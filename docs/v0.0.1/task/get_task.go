package task

/**
 *
 * @api {GET} /task_tracker/tasks/task Получение задачи
 * @apiName GetTask
 * @apiGroup 02 Задачи
 * @apiVersion  0.0.1
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {Uint64} id ИД задачи которую нужно получить
 *
 * @apiExample  Example usage:
 * http://localhost:8585/task_tracker/tasks/task?id=1
 *
 * @apiSuccess (Success 200) {String} 		status 					Статус ответа на запрос
 * @apiSuccess (Success 200) {Data} 		data 					Объект с данными по запросу
 * @apiSuccess (Success 200) {Uint64} 		data.id					ИД задачи
 * @apiSuccess (Success 200) {String} 		data.name 				Название задачи
 * @apiSuccess (Success 200) {String} 		data.description 		Описание задачи
 * @apiSuccess (Success 200) {String} 		data.creation_date 		Дата создания задачи
 * @apiSuccess (Success 200) {String} 		data.in_work_time 		Время проведенное в работе в `секундах`
 * @apiSuccess (Success 200) {Status} 		data.status 			Статус задачи
 * @apiSuccess (Success 200) {Uint64} 		data.status.id 			ИД статуса
 * @apiSuccess (Success 200) {String} 		data.status.name 		Название статус
 * @apiSuccess (Success 200) {User} 		data.author				Автор задачи
 * @apiSuccess (Success 200) {Uint64} 		data.author.id 			ИД автора задачи
 * @apiSuccess (Success 200) {String} 		data.author.name 		Имя автора
 * @apiSuccess (Success 200) {User} 		data.developer 			Разработчик
 * @apiSuccess (Success 200) {Uint64} 		data.developer.id 		ИД разработчика
 * @apiSuccess (Success 200) {String} 		data.developer.name		Имя разработчика
 * @apiSuccess (Success 200) {User} 		data.customer 			Заказчик задачи
 * @apiSuccess (Success 200) {Uint64} 		data.customer.id 		ИД заказчика
 * @apiSuccess (Success 200) {String} 		data.customer.name 		Имя заказчика
 * @apiSuccess (Success 200) {Category} 	data.category 			Категория задачи
 * @apiSuccess (Success 200) {Uint64} 		data.category.id 		ИД категории
 * @apiSuccess (Success 200) {String} 		data.category.name 		Название категории
 * @apiSuccess (Success 200) {Project} 		data.project 			Проект к которому относится задача
 * @apiSuccess (Success 200) {Uint64} 		data.project.id			ИД проекта
 * @apiSuccess (Success 200) {String} 		data.project.name 		Название проекта
 * @apiSuccess (Success 200) {Bool} 		data.priority 			Приоритет задачи (`true` - экстренный, `false` - стандартный )
 * @apiSuccess (Success 200) {Uint64} 		data.exec_order			Порядок выполнения задачи. 1 - стандартный, 10 - наивысшый.
 *
 * @apiSuccessExample {type} Success-Response:
 * {
 *   "status": "success",
 *   "data": {
 *     "id": 1,
 *     "name": "test",
 *     "description": "test",
 *     "creation_date": "2021-08-31T11:21:13+03:00",
 *     "in_work_time": "1h23m20s",
 *     "status": {
 *       "id": 2,
 *       "name": "Очередь к реализации"
 *     },
 *     "author": {
 *       "id": 1,
 *       "name": ""
 *     },
 *     "developer": {
 *       "id": 2,
 *       "name": ""
 *     },
 *     "customer": {
 *       "id": 1,
 *       "name": ""
 *     },
 *     "category": {
 *       "id": 1,
 *       "name": ""
 *     },
 *     "project": null,
 *     "priority": true,
 *     "exec_order": 0
 *   }
 * }
 *
 *
 */
