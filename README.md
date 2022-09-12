# go-training

## Useful links

* [gorm](https://gorm.io/index.html)
* [JSON validator](https://github.com/go-playground/validator)
* [gorilla mux](https://github.com/gorilla/mux)

### Envs
* [envconfig](https://github.com/kelseyhightower/envconfig)
* [godotenv](https://github.com/joho/godotenv)

### Testing
* [testify](https://github.com/stretchr/testify)
* [is](https://github.com/matryer/is)

## Blogs worth following

* [developer20.com](https://developer20.com/)
* [Ardan Labs](https://www.ardanlabs.com/blog/)
* [Dave Cheney](https://dave.cheney.net/)

## Running postgres

```sh
sudo docker run  --name myPostgresDb  -p 5432:5432  -e POSTGRES_USER=postgresUser  -e POSTGRES_PASSWORD=postgresPW  -e POSTGRES_DB=postgresDB  -d  postgres
```
