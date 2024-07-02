# Hotel Reservation Service Backend

## Project Outline

- users -> book room from an hotel
- admins -> check reservations / bookings
- Authentication & Authorizations -> JWT Tokens
- Hotels -> CRUD API -> JSON
- Rooms -> CRUD API -> JSON
- Scripts -> database managements -> Seeding, database migration etc.

## Resources

### Mongo db driver

Documentation

```
https://www.mongodb.com/docs/drivers/go/current/quick-start/
```

Installing mongodb client

```
go get go.mongodb.org/mongo-driver/mongo
```

### gofiber

Documentation

```
https://docs.gofiber.io/api/fiber
```

Installing gofiber

```
go get github.com/gofiber/fiber/v2
```

Usage

```
make run
```

This will run the server on port 5000 if default
