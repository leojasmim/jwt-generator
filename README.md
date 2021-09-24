# jwt-generator

API para facilitar a geração de tokens JWT

### Run
  ```dockerfile
  docker-compose up
  ```

### Test
  ```curl
curl --location --request POST 'http://localhost:8080/jwt' \
--form 'payload=\"{ \"message\" : \"teste\" }\"' \
--form 'sk="{{private_key}}"'
```

[Collection Postman](https://www.getpostman.com/collections/1a7b61a1c638d478e16a)