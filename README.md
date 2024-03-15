# Проект "Заметки"

Этот проект представляет собой веб-приложение для создания и управления заметками.

## Описание

Проект включает в себя функциональность для создания, редактирования и удаления заметок, а также просмотра списка заметок. Каждая заметка содержит текстовое описание и может быть привязана к определенному пользователю.

## API Structure

### Авторизация и аутентификация

- `POST /auth/sign-up`: Регистрация нового пользователя.
- `POST /auth/sign-in`: Авторизация пользователя.

### Заметки

- `POST /api/notes/`: Создание новой заметки.
- `GET /api/notes/`: Получение всех заметок.
- `GET /api/notes/:id`: Получение заметки по ID.
- `PUT /api/notes/:id`: Обновление заметки.
- `DELETE /api/notes/:id`: Удаление заметки.


  ## Структура базы данных

### Таблица "users"

- `id` (serial primary key): Идентификатор пользователя.
- `name` (varchar(255) not null): Имя пользователя.
- `username` (varchar(255) not null unique): Уникальное имя пользователя.
- `password_hash` (varchar(255) not null): Хэш пароля пользователя.

### Таблица "todo_lists"

- `id` (serial primary key): Идентификатор списка задач.
- `title` (varchar(255) not null): Название списка задач.
- `description` (varchar(255)): Описание списка задач.

### Таблица "users_lists"

- `id` (serial primary key): Идентификатор связи пользователя с списком задач.
- `user_id` (int references users(id) on delete cascade not null): Идентификатор пользователя (внешний ключ).
- `list_id` (int references todo_lists(id) on delete cascade not null): Идентификатор списка задач (внешний ключ).

### Таблица "todo_items"

- `id` (serial primary key): Идентификатор задачи.
- `title` (varchar(255) not null): Название задачи.
- `description` (varchar(255)): Описание задачи.
- `done` (boolean primary key): Статус завершения задачи.

### Таблица "lists_items"

- `id` (serial primary key): Идентификатор связи списка задач с задачей.
- `item_id` (int references todo_items(id) on delete cascade not null): Идентификатор задачи (внешний ключ).
- `lists_id` (int references todo_lists(id) on delete cascade not null): Идентификатор списка задач (внешний ключ).

 ## Author
 ### Yersultan Zhetibayev - 22B031250
