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
 * @apiSuccess (Success 200) {String} 		status 				Статус ответа на запрос
 * @apiSuccess (Success 200) {Data} 		data 				Объект с данными по запросу
 * @apiSuccess (Success 200) {Uint64} 		data.id				ИД задачи
 * @apiSuccess (Success 200) {String} 		data.name 			Название задачи
 * @apiSuccess (Success 200) {String} 		data.description 	Описание задачи
 * @apiSuccess (Success 200) {String} 		data.creation_date 	Дата создания задачи
 * @apiSuccess (Success 200) {String} 		data.in_work_time 	Время работы над задачей
 * @apiSuccess (Success 200) {Status} 		data.status 		Статус задачи
 * @apiSuccess (Success 200) {Uint64} 		data.status.id 		ИД статуса
 * @apiSuccess (Success 200) {String} 		data.status.name 	Название статус
 * @apiSuccess (Success 200) {Bool} 		data.priority 		Приоритет задачи (`true` - экстренный, `false` - стандартный )
 * @apiSuccess (Success 200) {Uint64} 		data.exec_order		Порядок выполнения задачи
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "status": "success",
 *   "data": [
 *     {
 *       "id": 1,
 *       "name": "test",
 *       "description": "test",
 *       "creation_date": "2021-08-31T11:21:13+03:00",
 *       "in_work_time": "1h23m20s",
 *       "status": {
 *         "id": 2,
 *         "name": "Очередь к реализации"
 *       },
 *       "priority": true,
 *       "exec_order": 0
 *     },
 *     {
 *       "id": 2,
 *       "name": "banner",
 *       "description": "banner",
 *       "creation_date": "2021-08-31T11:21:28+03:00",
 *       "in_work_time": "0s",
 *       "status": {
 *         "id": 1,
 *         "name": "Новая"
 *       },
 *       "priority": false,
 *       "exec_order": 0
 *     }
 *   ]
 * }
 *
 *
 */
