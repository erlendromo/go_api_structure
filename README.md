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

`cmd/`

`deploy/`

`docs/`

`internal/`
