package user

/**
 *
 * @api {GET} /task_tracker/settings/permissions/list Список разрешений
 * @apiName GetPermissionsList
 * @apiGroup 03 Пользователи
 * @apiVersion  0.0.1
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (Success 200) {String} 				status 								Статус выполнения запроса
 * @apiSuccess (Success 200) {Permissions} 		data 									Список доступов
 * @apiSuccess (Success 200) {Permissions} 		data.perm_sub_group 	Подгруппа доступов (в примере `settings`, `group` и т.д) которая также может содержать другие подгруппы доступов, а также список конечных действий
 * @apiSuccess (Success 200) {[]String} 			data.actions 					Список конечных действий для подгруппы доступов
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "status": "success",
 *   "data": {
 *     "settings": {
 *       "group": {
 *         "actions": [
 *           "update",
 *           "create"
 *         ]
 *       },
 *       "user": {
 *         "actions": [
 *           "update"
 *         ]
 *       }
 *     },
 *     "task": {
 *       "actions": [
 *         "create",
 *         "update"
 *       ],
 *       "get": {
 *         "actions": [
 *           "author",
 *           "dev",
 *           "customer"
 *         ]
 *       }
 *     }
 *   }
 * }
 *
 *
 */
