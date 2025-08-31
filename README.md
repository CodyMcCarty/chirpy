# Chirpy
Twitter clone. Based on a course with boot.dev

# Features
- HTTP server in Go, *without the use of a framework*
- JSON, headers, and status codes to communicate with clients via a RESTful API
- Type safe SQL to store and retrieve data from a Postgres database with [sqlc](https://sqlc.dev/)
  - SQLC is an amazing Go program that generates Go code from SQL queries. It's not exactly an ORM, but rather a tool that makes working with raw SQL easy and type-safe.
- Authentication/authorization system with well-tested cryptography libraries
  - github.com/golang-jwt/jwt & golang.org/x/crypto/bcrypt
- Webhooks and API keys
- Database migrations with [goose](https://github.com/pressly/goose)
  - Goose is a database migration tool written in Go. It runs migrations from a set of SQL files, making it a perfect fit for this project (we wanna stay close to the raw SQL).

# WIP
- [] Postman. Partial complete. I used curl or a short js script in chrome dev tools for testing.   
[Postman link](https://www.postman.com/security-specialist-67141284/chirpy/collection/gn4z824/chirpy?action=share&source=copy-link&creator=18651993) OR  
[![Run in Postman](https://run.pstmn.io/button.svg)]()
- [] [Swagger //localhost](http://localhost:8080/swagger-ui.html#/) todo

---

# personal notes:

How to get into this db?
```bash
sudo -u postgres psql
\c chirpy
or
psql "postgres://postgres:postgres@localhost:5432/chirpy"
```
how to migration?    
goose postgres protocol://username:password@host:port/database  (up or down)  
```bash
cd sql/schema/
goose postgres postgres://postgres:postgres@localhost:5432/chirpy up
```

how to new sql queries
```bash
cd project root

sqlc generate
```