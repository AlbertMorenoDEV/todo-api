# todo-api

[![codecov](https://codecov.io/gh/AlbertMorenoDEV/todo-api/branch/master/graph/badge.svg)](https://codecov.io/gh/AlbertMorenoDEV/todo-api)
[![build](https://github.com/AlbertMorenoDEV/todo-api/workflows/Build%20and%20Test/badge.svg)](https://github.com/AlbertMorenoDEV/todo-api/actions?query=workflow%3A%22Build+and+Test%22)

Simple playground project to try out DDD and Hexagonal approaches using Golang.

## How To

### Create a ToDo

`curl --location --request POST 'http://localhost:3000/todos' --header 'Content-Type: application/json' --data-raw '{"id":"1", "title":"New Todo", "due":1609459200}' | jq`

### List ToDos

`curl localhost:3000/todos | jq`

## Requirements
- [x] Create a new todo
- [x] Edit an existing todo
- [ ] Delete an existing todo
- [ ] Github test pipeline
- [ ] AWS ECS deployment
- [ ] Logging improvement

## Resources
- API Example: https://thenewstack.io/make-a-restful-json-api-go/ > https://github.com/corylanou/tns-restful-json-api
- Docker Setup Example: https://github.com/callicoder/go-docker-compose
- Hot reload: https://levelup.gitconnected.com/docker-for-go-development-a27141f36ba9
- API Example: https://github.com/friendsofgo/gopherapi/tree/v0.3.3
- CQRS Example: https://github.com/jetbasrawi/go.cqrs/tree/master/examples/simplecqrs
- VO Example: https://github.com/codeinabox/go-go-valueobject
- Tests on Docker: https://www.ardanlabs.com/blog/2019/03/integration-testing-in-go-executing-tests-with-docker.html
- E2E test framework: https://github.com/gavv/httpexpect
- Testing package: https://github.com/stretchr/testify