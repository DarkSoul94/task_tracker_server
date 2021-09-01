package category

/**
 *
 * @api {GET} /task_tracker/category/list Получение списка категорий
 * @apiName GetCategoryList
 * @apiGroup 04 Категории
 * @apiVersion  0.0.1
 * @apiSampleRequest off
 *
 * @apiHeader {String} BearerToken 	Авторизационный токен
 *
 * @apiSuccess (200) {String} 		status 	Статус выполнения запроса
 * @apiSuccess (200) {[]Category} 	data 	Список категорий
 * @apiSuccess (200) {Uint64} 		id 		Ид категории
 * @apiSuccess (200) {String} 		name 	Имя категории
 *
 * @apiSuccessExample {type} Success-Response:
 * {
 *   "status": "success",
 *   "data": [
 *     {
 *       "id": 1,
 *       "name": "test"
 *     }
 *   ]
 * }
 *
 *
 */
