# Requirements
You are developing a software to account for maintenance tasks performed during a
working day. This application has two types of users (Manager, Technician).
The technician performs tasks and is only able to see, create or update his own
performed tasks.
The manager can see tasks from all the technicians, delete them, and should be
notified when some tech performs a task.
A task has a summary (max: 2500 characters) and a date when it was performed, the
summary from the task can contain personal information.
## Notes:
- If you don’t have enough time to complete the test you should prioritize
complete features (with tests) over many features;
- We’ll evaluate security, quality and readability of your code;
- This test is suitable for all levels of developers, so make sure to prove yours.
## Development
### Features:
- Create API endpoint to save a new task;
- Create API endpoint to list tasks;
- Notify manager of each task performed by the tech (This notification can be
just a print saying “The tech X performed the task Y on date Z”);
- This notification should not block any http request.
### Tech Requirements:
- Use either Go or Node to develop this HTTP API;
- Create a local development environment using docker containing this service
and a MySQL database;
- Use MySQL database to persist data from the application;
- Features should have unit tests to ensure they are working properly
### Bonus
- Use a message broker to decouple notification logic from the application flow;
- Create Kubernetes object files needed to deploy this application

## My strategy to tackle the challenge

### 1 - Configure the dev environment
- Create a `docker-compose.yml` to boot up the entire environment with two applications and the MySQL database:
  - API - which will handle the application logic
  - Notifications - which will handle the notifications, using NSQ(https://github.com/nsqio/nsq)
  - MySQL (https://hub.docker.com/_/mysql#:~:text=Example%20docker%2Dcompose.yml%20for%20mysql%3A)
- Use the library air(https://github.com/cosmtrek/air) for hot reload.
- Setup migrate(https://github.com/golang-migrate/migrate) to handle migrations. Check this SO answer for reference https://stackoverflow.com/a/55779980/13701400

### 2 - The API application
- Create a "api" package to boot up the application "controllers"
- Create a middleware to authenticate the user making the request
- Use the chi library to configure the API routes (https://github.com/go-chi/chi)
  - `GET /task`
  - `POST /task`
  - `PATCH /task`
  - `DELETE /task`
- Use the validator lib to validate the requests (https://github.com/go-playground/validator)
- Create a "entities" package with the Manager and Technician structs.
- Use the MySQL lib (https://github.com/go-sql-driver/mysql)
  - Create a "driver" package to handle the setup of the MySQL connection
  - Create a "datastore" package with the logic to interact with the database
- Configure the NSQ client, example: https://github.com/nsqio/nsq/blob/master/apps/nsq_to_http/nsq_to_http.go