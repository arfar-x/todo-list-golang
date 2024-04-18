## Todo list with multiple interfaces - REST API and CLI

**THIS PROJECT IS ONLY FOR EDUCATIONAL PURPOSES.**
*My very first project in Go*

This project is defined to gain a hands-on ability to write dynamic applications using Golang.
In this project, I tried to bring the ability to switch between HTTP and CLI while avoiding code duplication for each
interface.

So, to run the project with HTTP APIs, you can simply clone it (whether you want to compile it or not); then run it via:

```shell
./todo_list http
```
After you run the command above, HTTP server (Gin) starts.

### Endpoints and commands for CRUD

#### List

```shell
# Create a task via REST API
curl --location '127.0.0.1:8001/list'

# To show the lists
./todo_list cli list
```
And here is the HTTP response 200 OK:
```json
{
  "data": {
    "id": 1,
    "name": "Run Golang",
    "done": true,
    "created_at": "2024-01-01T12:00:00+00:00",
    "updated_at": "2024-01-01T12:00:00+00:00"
  },
  "status": 200
}
```
If you run the CLI command, the JSON response will be written to a file next to the executable file.

The interesting feature about switching between HTTP and CLI is, you don't have to customize any column or database query for
each interface. The columns are specified via struct `tags` and `reflect` dynamically to allow which column gets
updated.
If one column does not have the tag `flag`, it would not be queried.
-------------
#### Create
```shell
curl --location '127.0.0.1:8001/list' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Make awesome",
    "done": true
}'

# To create a task
./todo_list cli create --name=task-one --done=1 --not-existed-col=no --no=way
# The task is created with the two specified variables of columns `name` and `done`.
# The flags '--not-existed-col' and '--no' are ignored.
```

Response 201 Created:
```json
{
    "data": {
        "id": 1,
        "name": "Solutions",
        "done": false,
        "created_at": "2024-01-01T00:00:00.000+00:00",
        "updated_at": "2024-01-01T00:00:00.000+00:00"
    },
    "status": 201
}
```
-------------
#### Update
```shell
curl --location --request PUT '127.0.0.1:8001/list/50' \
--header 'Content-Type: application/json' \
--data '{
    "name": "Soft",
    "done": false
}'

# To update a task by ID
./todo_list cli update --id=1 --name=updated-task-two --done=1 --existed-col=no --no=way
# The task is found by `--id=1`  then the two columns `name` and `done` get affected.
```
Response 200 OK:
```json
{
    "data": {
        "id": 1,
        "name": "updated-task-two",
        "done": true,
        "created_at": "2024-01-01T00:00:00.000+00:00",
        "updated_at": "2024-01-01T00:00:00.000+00:00"
    },
    "status": 200
}
```
-------------
```shell
curl --location --request DELETE '127.0.0.1:8001/list/92' \
--header 'Content-Type: application/json' \
--data '{
    "name": "networks",
    "done": false
}'

# To delete a task by ID
./todo_list cli delete --id=2
```
Response 204 No Content:
```json5
// :D
```

### Creating `tasks` table
```mysql
create table tasks
(
    id         bigint unsigned auto_increment primary key,
    name       varchar(255)                         not null,
    done       tinyint(1) default 0                 not null,
    created_at timestamp  default CURRENT_TIMESTAMP null,
    updated_at timestamp  default CURRENT_TIMESTAMP null,
    deleted_at timestamp                            null
);
```

### Final speech
As it was mentioned above, this project is only for educational purposes and lacks production-ready key features such middlewares,
database error handling, etc, to increase fault tolerance.
