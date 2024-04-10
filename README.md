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
|        |        | ----- config.go
|        |
|        | ----- constants/
|        |
|        | ----- html/
|        |        |
|        |        | ----- index.html
|        |
|        | ----- http/
|        |        |
|        |        | ----- handlers/ (http-handlers for different endpoints)
|        |        |
|        |        | ----- route/
|        |        |        |
|        |        |        | ----- router.go (route requests to the different handlers)
|        |
|        | ----- mocks/
|        |
|        | ----- utils/
|
| ----- .env (remember to .gitignore this file in actual projects)
|
| ----- .gitignore
|
| ----- docker-compose.yaml
|
| ----- Dockerfile
|
| ----- go.mod
|
| ----- <go.sum> (will be created when using go get <package>)
|
| ----- README.md (this file)
```

##### `cmd` folder

This folder contains all the entry points of the application.

-   `api`: This folder contains the main entry point of the REST API server. The `main.go` file in the `server` sub-folder is responsible for starting the server and setting up all the necessary routes.

##### `internal` folder

This folder contains all the business logic and other implementation details of the application. It is structured as follows:

-   `business` folder

    -   `domains` folder: This folder contains domain-specific logic, such as the business rules for creating, updating, and deleting users.

    -   `usecases` folder: This folder contains the implementation of the use cases that are defined in the domains folder.

-   `config` folder

    -   `config.go`: This file reads the environment variables and sets up the configuration for the application.

-   `constants` folder

    -   this folder contains constant values used throughout the application.

-   `html` folder

    -   this folder contains html-files to be deployed in the application.

-   `http` folder

    -   `handlers` folder: This folder contains the implementation of HTTP handlers, which handle incoming HTTP requests and send responses back to the client.
    -   `route` folder: This folder contains the implementation of routes, which map URLs to handlers.

-   `mocks` folder

    -   this folder contains the implementation of mock objects used in tests.

-   `utils` folder

    -   this folder contains utility functions and classes used throughout the application.
