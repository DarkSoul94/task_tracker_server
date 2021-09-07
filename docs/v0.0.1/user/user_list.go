package user

/**
 *
 * @api {GET} /task_tracker/settings/user/list Список пользователей
 * @apiName GetUserList
 * @apiGroup 03 Пользователи
 * @apiVersion  0.0.1
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (200) {String} 		status 						Статус выполнения запроса
 * @apiSuccess (200) {[]User} 		data 							Данные(список пользователей)
 * @apiSuccess (200) {Uint64} 		data.id 					Ид пользователя
 * @apiSuccess (200) {String} 		data.name 				Имя пользователя
 * @apiSuccess (200) {Uint64} 		data.group_id 		Ид группы
 * @apiSuccess (200) {String} 		data.department 	Отдел пользователя в домене
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "status": "success",
 *   "data": [
 *     {
 *       "id": 1,
 *       "name": "Евгений Николаевич Табаков",
 *       "group_id": 2
 *     },
 *     {
 *       "id": 2,
 *       "name": "Вячеслав Викторович Тищенко",
 *       "group": 2
 *     }
 *   ]
 * }
 *
 *
 */
