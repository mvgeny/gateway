# grpc-gateway + keycloak
The purpose of this project is to tie together [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) and keycloak authentication via swagger/grpcui.

## Usage

  1. Go to `/docker` directory and execute `docker-compose up -d` command. Keycloak instance will be running. You can login to admin console http://localhost:8085/admin/master/console/#/ with admin/admin creds.
  2. Go to project root and do `make run`
  3. Discover [swaggerui](http://localhost:8082/swaggerui/), [grpcui](http://localhost:8082/grpcui/)

  To update contracts change proto file and run `make generate`. If needed go to [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) docs and install dependencies locally.

## Implementation details
* After starting the keycloak, the auto import will be performed (custom realm **gateway**, client **myclient** and user with **dev/dev** creds). Imlicit Flow is enabled inside client. These creds should be used while authenticating inside the app!
* SwaggerUI is embeded in native go code with [statik](https://github.com/rakyll/statik) library. Swagger [dist](https://github.com/swagger-api/swagger-ui/tree/master/dist) was tuned to consume `apidocs.swagger.json`
* Swagger authorization is typical with Authorize button and client_id: **myclient**. Grpcui is auto redirecting to auth on page loaded and token is injected in metadata.
![grpcui metadata updated](/docs/grpcui.png)