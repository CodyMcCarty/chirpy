# chirpy

personal notes:

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
