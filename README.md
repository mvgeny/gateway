# grpc-gateway + keycloak
The purpose of this project is to tie together [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) and keycloak authentication via swagger/grpcui.

Blog post also is available on [nedo.tech](https://nedo.tech/post/grpc-gateway-auth/)
## Usage

  1. Go to `/docker` directory and execute `docker-compose up -d` command. Keycloak instance will be running.
  2. Go to project root and do `make run`
  3. Discover [swaggerui](http://localhost:8082/swaggerui/), [grpcui](http://localhost:8082/grpcui/)
  4. At swaggerui go to Authorize and use `myclient` as client_id. Login/pass is **dev/dev**.

## Implementation details
* After starting the keycloak, the auto import will be performed (custom realm **gateway**, client **myclient** and user with **dev/dev** creds). Implicit Flow is enabled inside client. These creds should be used while authenticating inside the app!
* Also you can login to [admin console](http://localhost:8085/admin/master/console/#/) with **admin/admin** creds.
* SwaggerUI is embedded in native go code with [statik](https://github.com/rakyll/statik) library. Swagger [dist](https://github.com/swagger-api/swagger-ui/tree/master/dist) was tuned to consume `apidocs.swagger.json`
* Swagger authorization is typical with Authorize button and client_id: **myclient**. Grpcui is redirecting to auth after page loaded and token is auto injected in metadata.
![grpcui metadata updated](/docs/grpcui.png)