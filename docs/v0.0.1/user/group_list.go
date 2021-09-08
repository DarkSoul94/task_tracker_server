package user

/**
 *
 * @api {GET} /task_tracker/settings/group/list Список групп
 * @apiName GetGroupList
 * @apiGroup 03 Пользователи
 * @apiVersion  0.0.1
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {String} 		status 						Статус выполнения запроса
 * @apiSuccess (Success 200) {[]Group} 		data 							Данные(список групп)
 * @apiSuccess (Success 200) {Uint64} 		data.id 					Ид группы
 * @apiSuccess (Success 200) {String} 		data.name 				Имя группы
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "status": "success",
 *   "data": [
 *     {
 *       "id": 1,
 *       "name": "Regular user"
 *     },
 *     {
 *       "id": 2,
 *       "name": "Admin"
 *     },
 *     {
 *       "id": 3,
 *       "name": "PM"
 *     },
 *     {
 *       "id": 4,
 *       "name": "Developer"
 *     }
 *   ]
 * }
 *
 *
 */
