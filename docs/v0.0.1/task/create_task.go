package task

/**
 *
 * @api {POST} /task_tracker/tasks Создание задачи
 * @apiName CreateTask
 * @apiGroup 02 Задачи
 * @apiVersion  0.0.1
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiParam  {String} 	name						Название задачи
 * @apiParam  {String} 	description			Описание задачи
 * @apiParam  {Uint64} 	category_id			ИД категории к которой относится задача
 * @apiParam  {Uint64} 	project_id			ИД проекта к которому относится задача `(опционально)`
 * @apiParam  {Uint64} 	developer_id		ИД разработчика `(опционально)`. Если не передается задаче присваивается статус "Новая", елси указан - "Очередь к реализации"
 * @apiParam  {Uint64} 	customer_id			ИД заказчика `(опционально)`.
 * @apiParam  {Bool} 		priority				Приоритет `(опционально)`. Если не указан выставляется стандартный приоритет для задачи
 * @apiParam  {Uint64} 	exec_order			Указатель в каком порядке нужно выполнять задачу `(опционально)`.
 *
 * @apiSuccess (Success 200) {String} status  Статус ответа на запрос
 *
 * @apiParamExample  {json} Min-Request-Example:
 *	{
 *			"name": "test",
 *			"description": "test",
 *			"category_id": 1
 *	}
 *
 * @apiParamExample  {json} Request-Example:
 * {
 * 			"name": "test",
 * 			"description": "test",
 * 			"category_id": 1,
 * 			"developer_id": 2,
 * 			"priority": true,
 * 			"customer_id": 1
 * }
 *
 *
 * @apiSuccessExample {json} Success-Response:
 * {
 *  "status": "success"
 * }
 *
 *
 */
