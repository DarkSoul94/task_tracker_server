package task

/**
 *
 * @api {POST} /task_tracker/tasks/track Трекание задачи
 * @apiName TrackTask
 * @apiGroup 02 Задачи
 * @apiVersion  0.0.1
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {Uint64} 	task_id 	Ид задачи
 * @apiParam  {Bool} 	status 		Признак начать или закончить трекать задачу
 *
 * @apiSuccess (Success 200) {String} Status Статус выполнения запроса
 *
 * @apiParamExample  {json} Request-Example:
 * {
 *     "task_id":3,
 *     "status": true
 * }
 *
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *     "status": "success"
 * }
 *
 *
 */
