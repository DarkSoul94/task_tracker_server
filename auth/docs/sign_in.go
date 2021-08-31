package docs

/**
 *
 * @api {POST} /task_tracker/auth/signin Авторизация
 * @apiName SignIn
 * @apiGroup 01 Авторизация
 * @apiVersion  0.0.1
 * @apiSampleRequest off
 *
 *
	* @apiParam  {String} username Доменная электронная почта
	* @apiParam  {String} password Пароль от доменной электронной почты
 *
 * @apiSuccess (Success 200) {String} 	status 									Статус ответа на запрос
 * @apiSuccess (Success 200) {Data} 	data 										Объект данных
 * @apiSuccess (Success 200) {Uint64} 	data.user_id						ИД пользователя
 * @apiSuccess (Success 200) {String} 	data.user_name					Имя пользователя
 * @apiSuccess (Success 200) {String} 	data.token							Авторизационный токен
 * @apiSuccess (Success 200) {Group} 	data.group 							Объект группы пользователя
 * @apiSuccess (Success 200) {Uint64} 	data.group.group_id			ИД группы
 * @apiSuccess (Success 200) {String} 	data.group.group_name		Название группы
 * @apiSuccess (Success 200) {JSON} 		data.group.permissions 	JSON объект с доступами которые есть у группы
 *
 *
 * @apiParamExample  {json} Request-Example:
 * {
 *     "username": "test@limefin.com",
 *		 "password": "Qwerty1234"
 * }
 *
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "status": "success",
 *   "data": {
 *     "user_id": 2,
 *     "user_name": "Евгений Николаевич Табаков",
 *     "token": "<token>",
 *     "group": {
 *       "group_id": 2,
 *       "group_name": "Admin",
 *       "permissions": {
 *         	"task": [
 *         	  "task.create",
 *         	  "task.get.all"
 *         	],
 *         	"user": [
 *         	  "user.update",
 *         	  "user.get"
 *         	]
 *       }
 *     }
 *   }
 * }
 *
 *
*/
