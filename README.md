# true-conf

### Тестовое задание:
"Вам нужно написать маленькое приложение на ECHO framework предоставляющие REST API по работе с сущностью User.

REST API должно удовлетворять следующие возможности:
* Добавление User
* Получение списка User
* Получение User по Id
* Редактирование User по Id
* Удаление User по Id


REST API должно работать с форматом данных JSON.


Сущность User должно состоять минимум из следующих полей:
* Идентификатор пользователя
* Отображаемое имя


В качестве хранилища данных нужно использовать файл в формате JSON.

### Реализовывает следующие возможности:
* Добавление User
* Получение списка User
* Получение User по Id
* Редактирование User по Id
* Удаление User по Id

### Методы:
По умолчанию запускается на порту :8080, можно изменить с помощью флага `-port`

* `POST /user?name={userName}` Добавление User
* `GET /user?id={id}` Получение User по ID
* `PUT /user?id={id}&name={userName}` Редактирование User по ID
* `DELETE /user?id={id}` Удаление User по Id

* `GET /users` Получение списка User

