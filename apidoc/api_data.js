define({ "api": [
  {
    "type": "POST",
    "url": "/task_tracker/auth/signin",
    "title": "Авторизация",
    "name": "SignIn",
    "group": "01_Авторизация",
    "version": "0.0.1",
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "username",
            "description": "<p>Доменная электронная почта</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "password",
            "description": "<p>Пароль от доменной электронной почты</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"username\": \"test@limefin.com\",\n\t\t \"password\": \"Qwerty1234\"\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          },
          {
            "group": "Success 200",
            "type": "Data",
            "optional": false,
            "field": "data",
            "description": "<p>Объект данных</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.user_id",
            "description": "<p>ИД пользователя</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.user_name",
            "description": "<p>Имя пользователя</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.token",
            "description": "<p>Авторизационный токен</p>"
          },
          {
            "group": "Success 200",
            "type": "Group",
            "optional": false,
            "field": "data.group",
            "description": "<p>Объект группы пользователя</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.group.group_id",
            "description": "<p>ИД группы</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.group.group_name",
            "description": "<p>Название группы</p>"
          },
          {
            "group": "Success 200",
            "type": "JSON",
            "optional": false,
            "field": "data.group.permissions",
            "description": "<p>JSON объект с доступами которые есть у группы</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"success\",\n  \"data\": {\n    \"user_id\": 2,\n    \"user_name\": \"Евгений Николаевич Табаков\",\n    \"token\": \"<token>\",\n    \"group\": {\n      \"group_id\": 2,\n      \"group_name\": \"Admin\",\n      \"permissions\": {\n        \t\"task\": [\n        \t  \"task.create\",\n        \t  \"task.get.all\"\n        \t],\n        \t\"user\": [\n        \t  \"user.update\",\n        \t  \"user.get\"\n        \t]\n      }\n    }\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./auth/docs/sign_in.go",
    "groupTitle": "01_Авторизация"
  },
  {
    "type": "POST",
    "url": "/task_tracker/tasks",
    "title": "Создание задачи",
    "name": "CreateTask",
    "group": "02_Задачи",
    "version": "0.0.1",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Название задачи</p>"
          },
          {
            "group": "Parameter",
            "type": "String",
            "optional": false,
            "field": "description",
            "description": "<p>Описание задачи</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "category_id",
            "description": "<p>ИД категории к которой относится задача</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "project_id",
            "description": "<p>ИД проекта к которому относится задача <code>(опционально)</code></p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "developer_id",
            "description": "<p>ИД разработчика <code>(опционально)</code>. Если не передается задаче присваивается статус &quot;Новая&quot;, елси указан - &quot;Очередь к реализации&quot;</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "customer_id",
            "description": "<p>ИД заказчика <code>(опционально)</code>.</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "priority",
            "description": "<p>Приоритет <code>(опционально)</code>. Если не указан выставляется стандартный приоритет для задачи</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "exec_order",
            "description": "<p>Указатель в каком порядке нужно выполнять задачу  от 1 до 10 <code>(опционально)</code>. Если не указан по умолчанию указывается 1</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Min-Request-Example:",
          "content": "{\n\t\t\t\"name\": \"test\",\n\t\t\t\"description\": \"test\",\n\t\t\t\"category_id\": 1\n}",
          "type": "json"
        },
        {
          "title": "Request-Example:",
          "content": "{\n\t\t\t\"name\": \"test\",\n\t\t\t\"description\": \"test\",\n\t\t\t\"category_id\": 1,\n\t\t\t\"developer_id\": 2,\n\t\t\t\"priority\": true,\n\t\t\t\"customer_id\": 1\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n \"status\": \"success\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/v0.0.1/task/create_task.go",
    "groupTitle": "02_Задачи"
  },
  {
    "type": "GET",
    "url": "/task_tracker/tasks/task",
    "title": "Получение задачи",
    "name": "GetTask",
    "group": "02_Задачи",
    "version": "0.0.1",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "id",
            "description": "<p>ИД задачи которую нужно получить</p>"
          }
        ]
      }
    },
    "examples": [
      {
        "title": "Example usage:",
        "content": "http://localhost:8585/task_tracker/tasks/task?id=1",
        "type": "json"
      }
    ],
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          },
          {
            "group": "Success 200",
            "type": "Data",
            "optional": false,
            "field": "data",
            "description": "<p>Объект с данными по запросу</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.id",
            "description": "<p>ИД задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.name",
            "description": "<p>Название задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.description",
            "description": "<p>Описание задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.creation_date",
            "description": "<p>Дата создания задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.in_work_time",
            "description": "<p>Время проведенное в работе в <code>секундах</code></p>"
          },
          {
            "group": "Success 200",
            "type": "Status",
            "optional": false,
            "field": "data.status",
            "description": "<p>Статус задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.status.id",
            "description": "<p>ИД статуса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.status.name",
            "description": "<p>Название статус</p>"
          },
          {
            "group": "Success 200",
            "type": "User",
            "optional": false,
            "field": "data.author",
            "description": "<p>Автор задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.author.id",
            "description": "<p>ИД автора задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.author.name",
            "description": "<p>Имя автора</p>"
          },
          {
            "group": "Success 200",
            "type": "User",
            "optional": false,
            "field": "data.developer",
            "description": "<p>Разработчик</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.developer.id",
            "description": "<p>ИД разработчика</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.developer.name",
            "description": "<p>Имя разработчика</p>"
          },
          {
            "group": "Success 200",
            "type": "User",
            "optional": false,
            "field": "data.customer",
            "description": "<p>Заказчик задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.customer.id",
            "description": "<p>ИД заказчика</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.customer.name",
            "description": "<p>Имя заказчика</p>"
          },
          {
            "group": "Success 200",
            "type": "Category",
            "optional": false,
            "field": "data.category",
            "description": "<p>Категория задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.category.id",
            "description": "<p>ИД категории</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.category.name",
            "description": "<p>Название категории</p>"
          },
          {
            "group": "Success 200",
            "type": "Project",
            "optional": false,
            "field": "data.project",
            "description": "<p>Проект к которому относится задача</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.project.id",
            "description": "<p>ИД проекта</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.project.name",
            "description": "<p>Название проекта</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "data.priority",
            "description": "<p>Приоритет задачи (<code>true</code> - экстренный, <code>false</code> - стандартный )</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.exec_order",
            "description": "<p>Порядок выполнения задачи. 1 - стандартный, 10 - наивысшый.</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "data.tracked",
            "description": "<p>Признак показывающий работает программист над задачей или нет.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"success\",\n  \"data\": {\n    \"id\": 1,\n    \"name\": \"test\",\n    \"description\": \"test\",\n    \"creation_date\": \"2021-08-31T11:21:13+03:00\",\n    \"in_work_time\": \"1h23m20s\",\n    \"status\": {\n      \"id\": 2,\n      \"name\": \"Очередь к реализации\"\n    },\n    \"author\": {\n      \"id\": 1,\n      \"name\": \"\"\n    },\n    \"developer\": {\n      \"id\": 2,\n      \"name\": \"\"\n    },\n    \"customer\": {\n      \"id\": 1,\n      \"name\": \"\"\n    },\n    \"category\": {\n      \"id\": 1,\n      \"name\": \"\"\n    },\n    \"project\": null,\n    \"priority\": true,\n    \"exec_order\": 0,\n    \"tracked\": false\n  }\n}",
          "type": "type"
        }
      ]
    },
    "filename": "./docs/v0.0.1/task/get_task.go",
    "groupTitle": "02_Задачи"
  },
  {
    "type": "GET",
    "url": "/task_tracker/tasks/statuses",
    "title": "Статусы задачи",
    "name": "GetTaskStatusList",
    "group": "02_Задачи",
    "version": "0.0.1",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Status",
            "optional": false,
            "field": "data",
            "description": "<p>Массив объектов &quot;статус задачи&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.id",
            "description": "<p>ИД статуса задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.name",
            "description": "<p>Имя статуса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"success\",\n  \"data\": [\n    {\n      \"id\": 1,\n      \"name\": \"Новая\"\n    },\n    {\n      \"id\": 2,\n      \"name\": \"Очередь к реализации\"\n    },\n    {\n      \"id\": 3,\n      \"name\": \"В работе\"\n    },\n    {\n      \"id\": 4,\n      \"name\": \"Приостановлена\"\n    },\n    {\n      \"id\": 5,\n      \"name\": \"Ожидание\"\n    },\n    {\n      \"id\": 6,\n      \"name\": \"Отклонена\"\n    },\n    {\n      \"id\": 7,\n      \"name\": \"Готово к тестированию\"\n    },\n    {\n      \"id\": 8,\n      \"name\": \"Выполнено\"\n    }\n  ]\n}",
          "type": "type"
        }
      ]
    },
    "filename": "./docs/v0.0.1/task/statuses.go",
    "groupTitle": "02_Задачи"
  },
  {
    "type": "GET",
    "url": "/task_tracker/tasks",
    "title": "Получение списка задач",
    "name": "GetTasksList",
    "group": "02_Задачи",
    "version": "0.0.1",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус ответа на запрос</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Task]",
            "optional": false,
            "field": "data",
            "description": "<p>Массив объектов &quot;задача&quot;</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.id",
            "description": "<p>ИД задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.name",
            "description": "<p>Название задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.description",
            "description": "<p>Описание задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.creation_date",
            "description": "<p>Дата создания задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.in_work_time",
            "description": "<p>Время работы над задачей</p>"
          },
          {
            "group": "Success 200",
            "type": "Status",
            "optional": false,
            "field": "data.status",
            "description": "<p>Статус задачи</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.status.id",
            "description": "<p>ИД статуса</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.status.name",
            "description": "<p>Название статус</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "data.priority",
            "description": "<p>Приоритет задачи (<code>true</code> - экстренный, <code>false</code> - стандартный )</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.exec_order",
            "description": "<p>Порядок выполнения задачи. 1 - стандартный, 10 - наивысшый.</p>"
          },
          {
            "group": "Success 200",
            "type": "Bool",
            "optional": false,
            "field": "data.tracked",
            "description": "<p>Признак показывающий работает программист над задачей или нет.</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"success\",\n  \"data\": [\n    {\n      \"id\": 1,\n      \"name\": \"test\",\n      \"description\": \"test\",\n      \"creation_date\": \"2021-08-31T11:21:13+03:00\",\n      \"in_work_time\": \"1h23m20s\",\n      \"status\": {\n        \"id\": 2,\n        \"name\": \"Очередь к реализации\"\n      },\n      \"priority\": true,\n      \"exec_order\": 0,\n    \t \"tracked\": false\n    },\n    {\n      \"id\": 2,\n      \"name\": \"banner\",\n      \"description\": \"banner\",\n      \"creation_date\": \"2021-08-31T11:21:28+03:00\",\n      \"in_work_time\": \"0s\",\n      \"status\": {\n        \"id\": 1,\n        \"name\": \"Новая\"\n      },\n      \"priority\": false,\n      \"exec_order\": 0,\n    \t \"tracked\": false\n    }\n  ]\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/v0.0.1/task/get_tasks_list.go",
    "groupTitle": "02_Задачи"
  },
  {
    "type": "POST",
    "url": "/task_tracker/tasks/track",
    "title": "Трекание задачи",
    "name": "TrackTask",
    "group": "02_Задачи",
    "version": "0.0.1",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "task_id",
            "description": "<p>Ид задачи</p>"
          },
          {
            "group": "Parameter",
            "type": "Bool",
            "optional": false,
            "field": "status",
            "description": "<p>Признак начать или закончить трекать задачу</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Request-Example:",
          "content": "{\n    \"task_id\":3,\n    \"status\": true\n}",
          "type": "json"
        }
      ]
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "Status",
            "description": "<p>Статус выполнения запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n    \"status\": \"success\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/v0.0.1/task/track_task.go",
    "groupTitle": "02_Задачи"
  },
  {
    "type": "GET",
    "url": "/task_tracker/settings/group/:id",
    "title": "Получение одной группы по ее ID",
    "name": "GetGroup",
    "group": "03_Пользователи",
    "version": "0.0.1",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус выполнения запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Group",
            "optional": false,
            "field": "data",
            "description": "<p>Данные(список пользователей)</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.id",
            "description": "<p>Ид группы</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.name",
            "description": "<p>Имя группы</p>"
          },
          {
            "group": "Success 200",
            "type": "Permissions",
            "optional": false,
            "field": "data.permissions",
            "description": "<p>Имя группы</p>"
          },
          {
            "group": "Success 200",
            "type": "Permissions",
            "optional": false,
            "field": "data.permissions.perm_sub_group",
            "description": "<p>Подгруппа доступов (в примере <code>settings</code>, <code>group</code> и т.д) которая также может содержать другие подгруппы доступов, а также список конечных действий</p>"
          },
          {
            "group": "Success 200",
            "type": "[]String",
            "optional": false,
            "field": "data.permissions.actions",
            "description": "<p>Список конечных действий для подгруппы доступов</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"success\",\n  \"data\": {\n    \"id\": 2,\n    \"name\": \"Admin\",\n    \"permissions\": {\n      \"settings\": {\n        \"group\": {\n          \"actions\": [\n            \"update\",\n            \"create\"\n          ]\n        },\n        \"user\": {\n          \"actions\": [\n            \"update\"\n          ]\n        }\n      },\n      \"task\": {\n        \"actions\": [\n          \"create\",\n          \"update\"\n        ],\n        \"get\": {\n          \"actions\": [\n            \"author\",\n            \"dev\",\n            \"customer\"\n          ]\n        }\n      }\n    }\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/v0.0.1/user/group_get.go",
    "groupTitle": "03_Пользователи"
  },
  {
    "type": "GET",
    "url": "/task_tracker/settings/group/list",
    "title": "Список групп",
    "name": "GetGroupList",
    "group": "03_Пользователи",
    "version": "0.0.1",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус выполнения запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "[]Group",
            "optional": false,
            "field": "data",
            "description": "<p>Данные(список групп)</p>"
          },
          {
            "group": "Success 200",
            "type": "Uint64",
            "optional": false,
            "field": "data.id",
            "description": "<p>Ид группы</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "data.name",
            "description": "<p>Имя группы</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"success\",\n  \"data\": [\n    {\n      \"id\": 1,\n      \"name\": \"Regular user\"\n    },\n    {\n      \"id\": 2,\n      \"name\": \"Admin\"\n    },\n    {\n      \"id\": 3,\n      \"name\": \"PM\"\n    },\n    {\n      \"id\": 4,\n      \"name\": \"Developer\"\n    }\n  ]\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/v0.0.1/user/group_list.go",
    "groupTitle": "03_Пользователи"
  },
  {
    "type": "GET",
    "url": "/task_tracker/settings/permissions/list",
    "title": "Список разрешений",
    "name": "GetPermissionsList",
    "group": "03_Пользователи",
    "version": "0.0.1",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус выполнения запроса</p>"
          },
          {
            "group": "Success 200",
            "type": "Permissions",
            "optional": false,
            "field": "data",
            "description": "<p>Список доступов</p>"
          },
          {
            "group": "Success 200",
            "type": "Permissions",
            "optional": false,
            "field": "data.perm_sub_group",
            "description": "<p>Подгруппа доступов (в примере <code>settings</code>, <code>group</code> и т.д) которая также может содержать другие подгруппы доступов, а также список конечных действий</p>"
          },
          {
            "group": "Success 200",
            "type": "[]String",
            "optional": false,
            "field": "data.actions",
            "description": "<p>Список конечных действий для подгруппы доступов</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"success\",\n  \"data\": {\n    \"settings\": {\n      \"group\": {\n        \"actions\": [\n          \"update\",\n          \"create\"\n        ]\n      },\n      \"user\": {\n        \"actions\": [\n          \"update\"\n        ]\n      }\n    },\n    \"task\": {\n      \"actions\": [\n        \"create\",\n        \"update\"\n      ],\n      \"get\": {\n        \"actions\": [\n          \"author\",\n          \"dev\",\n          \"customer\"\n        ]\n      }\n    }\n  }\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/v0.0.1/user/permissions.go",
    "groupTitle": "03_Пользователи"
  },
  {
    "type": "GET",
    "url": "/task_tracker/settings/user/list",
    "title": "Список пользователей",
    "name": "GetUserList",
    "group": "03_Пользователи",
    "version": "0.0.1",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус выполнения запроса</p>"
          },
          {
            "group": "200",
            "type": "[]User",
            "optional": false,
            "field": "data",
            "description": "<p>Данные(список пользователей)</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "data.id",
            "description": "<p>Ид пользователя</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "data.name",
            "description": "<p>Имя пользователя</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "data.group_id",
            "description": "<p>Ид группы</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "data.department",
            "description": "<p>Отдел пользователя в домене</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"success\",\n  \"data\": [\n    {\n      \"id\": 1,\n      \"name\": \"Евгений Николаевич Табаков\",\n      \"group_id\": 2\n    },\n    {\n      \"id\": 2,\n      \"name\": \"Вячеслав Викторович Тищенко\",\n      \"group\": 2\n    }\n  ]\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/v0.0.1/user/user_list.go",
    "groupTitle": "03_Пользователи"
  },
  {
    "type": "POST",
    "url": "/task_tracker/settings/user/update",
    "title": "Изменить группу пользователю",
    "name": "UserUpdate",
    "group": "03_Пользователи",
    "version": "0.0.1",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "user_id",
            "description": "<p>ИД пользователя</p>"
          },
          {
            "group": "Parameter",
            "type": "Uint64",
            "optional": false,
            "field": "group_id",
            "description": "<p>ИД новой группы пользователя</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус выполнения запроса</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"success\"\n}",
          "type": "json"
        }
      ]
    },
    "filename": "./docs/v0.0.1/user/user_update.go",
    "groupTitle": "03_Пользователи"
  },
  {
    "type": "GET",
    "url": "/task_tracker/category/list",
    "title": "Получение списка категорий",
    "name": "GetCategoryList",
    "group": "04_Категории",
    "version": "0.0.1",
    "header": {
      "fields": {
        "Header": [
          {
            "group": "Header",
            "type": "String",
            "optional": false,
            "field": "BearerToken",
            "description": "<p>Авторизационный токен</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "200": [
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "status",
            "description": "<p>Статус выполнения запроса</p>"
          },
          {
            "group": "200",
            "type": "[]Category",
            "optional": false,
            "field": "data",
            "description": "<p>Список категорий</p>"
          },
          {
            "group": "200",
            "type": "Uint64",
            "optional": false,
            "field": "id",
            "description": "<p>Ид категории</p>"
          },
          {
            "group": "200",
            "type": "String",
            "optional": false,
            "field": "name",
            "description": "<p>Имя категории</p>"
          }
        ]
      },
      "examples": [
        {
          "title": "Success-Response:",
          "content": "{\n  \"status\": \"success\",\n  \"data\": [\n    {\n      \"id\": 1,\n      \"name\": \"test\"\n    }\n  ]\n}",
          "type": "type"
        }
      ]
    },
    "filename": "./docs/v0.0.1/category/get_category_list.go",
    "groupTitle": "04_Категории"
  }
] });
