# go-ddd-playground

## How To

### Create a ToDo

`curl -H "Content-Type: application/json" -d '{"name":"New Todo"}' http://localhost:8080/todos | jq`

### List ToDos

`curl localhost:8080/todos | jq`

## Requirements
- Create a new todo
- Edit an existing todo
- Delete an existing todo

## Resources
- API Example: https://thenewstack.io/make-a-restful-json-api-go/ > https://github.com/corylanou/tns-restful-json-api
- Docker Setup Example: https://github.com/callicoder/go-docker-compose
- Hot reload: https://levelup.gitconnected.com/docker-for-go-development-a27141f36ba9