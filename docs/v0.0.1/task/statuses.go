package task

/**
 *
 * @api {GET} /task_tracker/tasks/statuses Статусы задачи
 * @apiName GetTaskStatusList
 * @apiGroup 02 Задачи
 * @apiVersion  0.0.1
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {String} 		status 				Статус ответа на запрос
 * @apiSuccess (Success 200) {[]Status} 	data 				Массив объектов "статус задачи"
 * @apiSuccess (Success 200) {Uint64} 		data.id 			ИД статуса задачи
 * @apiSuccess (Success 200) {String} 		data.name 			Имя статуса
 *
 * @apiSuccessExample {type} Success-Response:
 * {
 *   "status": "success",
 *   "data": [
 *     {
 *       "id": 1,
 *       "name": "Новая"
 *     },
 *     {
 *       "id": 2,
 *       "name": "Очередь к реализации"
 *     },
 *     {
 *       "id": 3,
 *       "name": "В работе"
 *     },
 *     {
 *       "id": 4,
 *       "name": "Приостановлена"
 *     },
 *     {
 *       "id": 5,
 *       "name": "Ожидание"
 *     },
 *     {
 *       "id": 6,
 *       "name": "Отклонена"
 *     },
 *     {
 *       "id": 7,
 *       "name": "Готово к тестированию"
 *     },
 *     {
 *       "id": 8,
 *       "name": "Выполнено"
 *     }
 *   ]
 * }
 *
 *
 */
