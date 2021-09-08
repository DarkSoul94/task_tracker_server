package user

/**
 *
 * @api {GET} /task_tracker/settings/group/:id Получение одной группы по ее ID
 * @apiName GetGroup
 * @apiGroup 03 Пользователи
 * @apiVersion  0.0.1
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {String} 				status 														Статус выполнения запроса
 * @apiSuccess (Success 200) {[]Group} 				data 															Данные(список пользователей)
 * @apiSuccess (Success 200) {Uint64} 				data.id 													Ид группы
 * @apiSuccess (Success 200) {String} 				data.name 												Имя группы
 * @apiSuccess (Success 200) {Permissions} 		data.permissions 									Имя группы
 * @apiSuccess (Success 200) {Permissions} 		data.permissions.perm_sub_group 	Подгруппа доступов (в примере `settings`, `group` и т.д) которая также может содержать другие подгруппы доступов, а также список конечных действий
 * @apiSuccess (Success 200) {[]String} 			data.permissions.actions 					Список конечных действий для подгруппы доступов
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "status": "success",
 *   "data": {
 *     "id": 2,
 *     "name": "Admin",
 *     "permissions": {
 *       "settings": {
 *         "group": {
 *           "actions": [
 *             "update",
 *             "create"
 *           ]
 *         },
 *         "user": {
 *           "actions": [
 *             "update"
 *           ]
 *         }
 *       },
 *       "task": {
 *         "actions": [
 *           "create",
 *           "update"
 *         ],
 *         "get": {
 *           "actions": [
 *             "author",
 *             "dev",
 *             "customer"
 *           ]
 *         }
 *       }
 *     }
 *   }
 * }
 *
 *
 */
