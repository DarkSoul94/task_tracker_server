package user

/**
 *
 * @api {GET} /task_tracker/user/list Список пользователей
 * @apiName GetUserList
 * @apiGroup 03 Пользователи
 * @apiVersion  0.0.1
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (200) {String} 		status 					Статус выполнения запроса
 * @apiSuccess (200) {[]User} 		data 					Данные(список пользователей)
 * @apiSuccess (200) {Uint64} 		data.id 				Ид пользователя
 * @apiSuccess (200) {String} 		data.name 				Имя пользователя
 * @apiSuccess (200) {Group} 		data.group 				Группа к которой относиться пользователь
 * @apiSuccess (200) {Uint64} 		data.group.id 			Ид группы
 * @apiSuccess (200) {String} 		data.group.name 		Имя группы
 * @apiSuccess (200) {Permissions} 	data.group.permissions 	Набор доступов
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "status": "success",
 *   "data": [
 *     {
 *       "id": 1,
 *       "name": "Евгений Николаевич Табаков",
 *       "group": {
 *         "id": 2,
 *         "name": "",
 *         "permissions": null
 *       }
 *     },
 *     {
 *       "id": 2,
 *       "name": "Вячеслав Викторович Тищенко",
 *       "group": {
 *         "id": 2,
 *         "name": "",
 *         "permissions": null
 *       }
 *     }
 *   ]
 * }
 *
 *
 */
