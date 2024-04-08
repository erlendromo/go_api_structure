# GO API STRUCTURE

This is a template for how to structure go apis. Scalability and maintainability are the main focuses here, as well as easy deployment using docker. The structure is as follows:

```
root/
|
| ----- cmd/
|        |
|        | ----- api/
|        |        |
|        |        | ----- server/
|        |        |        |
|        |        |        | ----- server.go
|        |        |
|        |        | ----- main.go
|
| ----- deploy/
|        |
|        | ----- docker-compose.yaml
|        |
|        | ----- Dockerfile
|
| ----- docs/
|        |
|        | ----- swagger.yaml
|
| ----- internal/
|        |
|        | ----- business/
|        |        |
|        |        | ----- domains/
|        |        |
|        |        | ----- usecases/
|        |
|        | ----- config/
|        |        |
|        |        | ----- .env
|        |        |
|        |        | ----- config.go
|        |
|        | ----- constants/
|        |
|        | ----- http/
|        |
|        | ----- mocks/
|        |
|        | ----- utils/
|
| ----- .gitignore
|
| ----- go.mod
|
| ----- <go.sum>
|
| ----- README.md (this file)
```

##### `cmd` folder

This folder contains all the entry points of the application.

-   `api`: This folder contains the main entry point of the REST API server. The `main.go` file in the `server` sub-folder is responsible for starting the server and setting up all the necessary routes.

##### `deploy` folder

This folder contains the necessary configuration files for deploying the application to a production environment.

-   `Dockerfile`: This file is used to build a Docker image of the application.

-   `docker-compose.yml`: This file is used to start the application and its dependencies (such as the database) using Docker Compose.

##### `docs` folder

This folder contains the documentation for the REST API, including the `swagger.yaml` file which defines the API specification.

##### `internal` folder

This folder contains all the business logic and other implementation details of the application. It is structured as follows:

-   `business` folder

    -   `domains` folder: This folder contains domain-specific logic, such as the business rules for creating, updating, and deleting users.

    -   `usecases` folder: This folder contains the implementation of the use cases that are defined in the domains folder.

-   `config` folder

    -   `.env`: This file contains the environment variables that are used by the application.
    -   `config.go`: This file reads the environment variables and sets up the configuration for the application.

-   `constants` folder

    -   this folder contains constant values used throughout the application.

-   `http` folder

    -   `datatransfers` folder: This folder contains the implementation of data transfer objects, such as request and response objects.
    -   `handlers` folder: This folder contains the implementation of HTTP handlers, which handle incoming HTTP requests and send responses back to the client.
    -   `routes` folder: This folder contains the implementation of routes, which map URLs to handlers.

-   `mocks` folder

    -   this folder contains the implementation of mock objects used in tests.

-   `utils` folder

    -   this folder contains utility functions and classes used throughout the application.
