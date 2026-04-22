# API_REST_GO_MYSQL

API REST construida con Go conectada a una base de datos MySQL. Este proyecto sirve como base para aprender a integrar un servicio HTTP con persistencia real de datos.

## Caracteristicas

- Operaciones CRUD basicas (crear, leer, actualizar, eliminar)
- Conexion a MySQL con driver oficial
- Respuestas en formato JSON
- Coleccion de Postman incluida en /docs/postman para probar los endpoints

## Tecnologias

- Go (Golang)
- MySQL
- Postman (para pruebas)

## Configuracion de la base de datos

Antes de ejecutar el proyecto, asegurate de tener MySQL corriendo localmente y una base de datos creada. Actualiza las credenciales de conexion en main.go con las tuyas.

## Como ejecutarlo

git clone https://github.com/ninjadiego/API_REST_GO_MYSQL.git
cd API_REST_GO_MYSQL
go mod tidy
go run main.go

## Pruebas

Puedes importar la coleccion de Postman que se encuentra en /docs/postman para probar todos los endpoints disponibles de forma rapida.

## Autor

ninjadiego - https://github.com/ninjadiego- - - - - - - 
