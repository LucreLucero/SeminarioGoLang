## Seminario GoLang 2020
___________
## Lucero Lucrecia - Sede Tandil

### Pasos a seguir para ejecutar nuestra api

1. Clonar el repositorio en nuestro GOPATH
2. Utilizar la consola con:

> go run cmd/candyshop/candysrv.go -config ./config/config.yaml

3. Navegar en forma local por el puerto 8080 

### GET : todos los caramelos

> http://localhost:8080/candies

### GET : un caramelo

> http://localhost:8080/candy

### POST : agrega un caramelo  

- Con body {"Text": lollipop}

> http://localhost:8080/candy

### PUT : actualizar un caramelo

- Con body {"Text": Caramel cookie}

> http://localhost:8080/candy

### DELETE : eliminar un caramelo

> http://localhost:8080/candy

## Consideraciones
- Realice la prueba en forma local en Chrome
- Para realizar el trabajo utilice Windows 7 y al ejecutar consola aparecia una advertencia, presionando continuar funciona correctamente



