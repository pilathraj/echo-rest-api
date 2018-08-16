## Project Setup
```
go get -u github.com/labstack/echo/
go get github.com/dgrijalva/jwt-go/
go get github.com/lib/pq/  not used.
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/postgres
go get gopkg.in/go-playground/validator.v9

```

## create database in the postgreSQL
#### run:
```
-- Database: "eshop-api"

-- DROP DATABASE "eshop-api";

CREATE DATABASE "eshop-api"
  WITH OWNER = postgres
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       LC_COLLATE = 'English_India.1252'
       LC_CTYPE = 'English_India.1252'
       CONNECTION LIMIT = -1;

```

### Create table
```
-- Table: public.products

-- DROP TABLE public.products;

CREATE TABLE public.products
(
  name text,
  sku uuid NOT NULL,
  description text,
  price double precision,
  CONSTRAINT sku PRIMARY KEY (sku)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE public.products
  OWNER TO postgres;
COMMENT ON TABLE public.products
  IS 'Products';
```


## Project Run
```
go run main.go
```

