# true-conf

Тестовое задание.

Маленькое приложение на ECHO framework предоставляющие REST API по работе с сущностью User.

##Реализовывает следующие возможности:
* Добавление User
* Получение списка User
* Получение User по Id
* Редактирование User по Id
* Удаление User по Id

##Методы:
По умолчанию запускается на порту :8080, можно изменить с помощью флага `-port`

* `POST /user?name={userName}` Добавление User
* `GET /user?id={id}` Получение User по ID
* `PUT /user?id={id}&name={userName}` Редактирование User по ID
* `DELETE /user?id={id}` Удаление User по Id

* `GET /users` Получение списка User

