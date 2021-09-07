package user

/**
 *
 * @api {POST} /task_tracker/settings/user/update Изменить группу пользователю
 * @apiName UserUpdate
 * @apiGroup 03 Пользователи
 * @apiVersion  0.0.1
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {Uint64} 	user_id				ИД пользователя
 * @apiParam  {Uint64} 	group_id			ИД новой группы пользователя

 * @apiSuccess (200) {String} 		status 						Статус выполнения запроса
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *   "status": "success"
 * }
 *
 *
 */
