# gin-openapi-v3

## why this project ?

[swag](https://github.com/swaggo/swag) v2 support generate openapi v3.1 docs,
but [gin-swagger](https://github.com/swaggo/gin-swagger) not support.

I referred to the design of [gin-swagger](https://github.com/swaggo/gin-swagger) and simplified the use of gin-swagger.

Now just copy the following two files to your project:

```
openapi/register.go
openapi/swagger-initializer.js
```

> In this way you can make some custom modifications according to the needs of the project.

## how to run?

1. install swag v2
    ```
    go install github.com/swaggo/swag/v2/cmd/swag@latest
    ```

2. generate openapi v3.1 docs
    ```
    swag init --v3.1 --output ./openapi --outputTypes json
    ```

   > generated `openapi/swagger.json` file, and ignored by `.gitignore` not push to git repo.

3. run app
    ```
    go run main.go
    ```

4. view docs
   open [http://localhost:8080/openapi](http://localhost:8080/openapi) in browser.

Enjoy!