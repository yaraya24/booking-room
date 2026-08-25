# Backend de Reserva de Salas de Reuniones

### Requisitos:
- Golang 
- Sqlite3 (gcc es necesario y hay que establecer la variable de entorno CGO_ENABLED=1)

### Instalación


1. Clona el repositorio
````clone git@github.com:yaraya24/booking-room.git```rr

2. Configura la base de datos 

```
sqlite3 ./booking_room.db < ./internal/db/setup-db.sql
```
3. Luego puedes ejecutar el servidor usando 
```
go run cmd/main.go 
```

4. O puedes compilar la aplicación para que sea un ejecutable que luego se pueda ejecutar:
```
go build ./cmd
```

### Uso

Necesitarás usar Basic Auth para acceder a la API.
Los usuarios están definidos en el archivo setub-db.sql, donde cada usuario tiene la contraseña `password`.
```
Jane:password
John:password
Sarah:password
```

Se puede crear una reserva usando el endpoint `POST localhost:8080/bookings`
Se necesita un cuerpo con un `room` y una `date`. la fecha debe tener el formato `YYYY/MM/DD`

ejemplo:
```
{
	"date": "2024-10-10",
	"room": "D"
}
```

Se puede acceder a las salas disponibles mediante `GET localhost:8080/bookings?{date}`, donde date tiene el formato `YYYY/MM/DD`.

## Errores/Problemas
1. Hay un error bastante grave, ya que los usuarios pueden reservar salas que no existen. Esto se debe a que la aplicación no verifica si una sala existe antes de hacer la reserva y confía ciegamente en el cliente. (me di cuenta de esto un poco tarde).

2. Me falta algo de validación para cuando los usuarios hacen una solicitud POST

3. Decidí usar un archivo sql para configurar la base de datos y, en consecuencia, no pude aplicar hash a las contraseñas. Ahora se almacenan en texto plano, lo cual no está bien.

4. No tenemos columnas meta en nuestras bases de datos, como marcas de tiempo de actualización y creación

5. No hacemos ping a la base de datos para asegurarnos de que realmente esté funcionando

6. Quería tener un middleware que proporcionara registro de la solicitud, id de solicitud, estado de la respuesta, etc., pero no tuve tiempo.

7. No tenía pruebas de integración reales; las únicas que agregué están en la capa de repositorio. De nuevo, por un tema de tiempo, aunque idealmente esto se haría mediante algo como Jenkins como una prueba de humo (smoke test).

## Mejoras
1. Hubiera estado bien agregar una caché como Redis o una caché en memoria para mejorar la escalabilidad de la aplicación. Al verificar las salas disponibles para una fecha, podríamos usar algo como una caché LRU y, cuando ocurre una reserva, esa fecha se puede actualizar. Esto se agrava por el hecho de que sqlite3 no permite acceso concurrente.

2. Mejorar la base de datos, ya sea migrando a MySQL o Postgres. Podría haber configurado algunas opciones para mejorar el rendimiento de sqlite, pero no tuve tiempo de analizarlo en detalle. Pero, en última instancia, se preferiría una base de datos lista para producción.

3. Por seguridad/fiabilidad, también sería bueno contar con un limitador de velocidad (rate limiter) para asegurar que nuestro servicio esté protegido contra un uso intensivo o incluso malicioso.
