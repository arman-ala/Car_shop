# Car Shop Project

In this project I have implemented a car shop application.

## Structure of the project

- `tests`: Contains 2 types of tests, unit tests and integration tests.
- `common`: Contains common code which is shared between different modules of the application.
- `constants`: Contains constants which are used in the application.
- `docs`: Contains API list of the application using Swagger.
- `data`: Contains all the files and directories related to the data; such as cache, database, models, etc.
- `docker`: Contains Docker files which are used to build the application in a container.
- `config`: Contains configuration files which are used to configure the application; such as environment variables, database connection, etc.
- `pkg`: Contains all third-party packages which are used.
- `service`: With this directory, we will implement all the business logic of the application.
- `scripts`: If we need any scripts to run the application, we will put them in this directory.
- `api`: Contains handlers to use business logic implemented in the service layer.
- `api/forms`: DTOs(Data Transfer Objects) which are used to transfer data between the client and the server.
- `api/routers`: Contains routers which are used to define the routes of the application.
- `api/validations`: Contains validations which are used to validate the input data.
- `api/middlewares`: Contains middlewares which are used to handle the requests and responses.

## The Flow of the Application
