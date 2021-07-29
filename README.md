# go-rest-api
A simple Go based REST API server built using [gorilla/mux](https://github.com/gorilla/mux) HTTP router and [PostgreSQL](https://www.postgresql.org/) database.

### Prerequisites
-----------------
 - Basic familiarity with Go and PostgreSQL, and
 - That you have working [Go](https://golang.org/) and [Docker](https://docs.docker.com/get-docker/) installations. You can use Docker to run a test database easily.

### Run locally
---------------
 - Start PostgreSQL container using:
    ```bash
    $ ./run.sh startdb
    ```
 - Once started, initialize database using:
    ```bash
    $ ./run.sh initdb products_db
    ```
 - Once initialized, start server using:
    ```bash
    $ ./run.sh server -port=<port_number> # default: 8010
    ```
 - Once started, you'll see a message like this:
    ```bash
    Go REST api-server listening on port: 8010
    Press Cmd+C / Ctrl+C to stop.
    ```
 - Now launch your browser and go to `http://localhost:8010/api/products`
 - Once done, stop PostgreSQL container using:
    ```bash
    $ ./run.sh stopdb
    ```
 - To clean-up database, use:
    ```bash
    $ ./run.sh cleanup
    ```

### Test
--------
To run unit tests, start PostgreSQL container then do:
```bash
$ ./run.sh test

=== RUN   TestEmptyTable
--- PASS: TestEmptyTable (0.01s)
=== RUN   TestGetNonExistentProduct
--- PASS: TestGetNonExistentProduct (0.00s)
=== RUN   TestCreateProduct
--- PASS: TestCreateProduct (0.01s)
=== RUN   TestGetProduct
--- PASS: TestGetProduct (0.01s)
=== RUN   TestUpdateProduct
--- PASS: TestUpdateProduct (0.02s)
=== RUN   TestDeleteProduct
--- PASS: TestDeleteProduct (0.02s)
PASS
ok  	github.com/adwardstark/go-rest-api/app	0.604s
```

### Endpoints
-------------
You can use [Postman](https://www.postman.com/) to perform CURD operations using following endpoints:
 - GET `/api/products`
 - POST `/api/product`
 - GET `/api/product/{id}`
 - PUT `/api/product/{id}`
 - DELETE `/api/product/{id}`

### License
-----------
Apache 2.0. See the [LICENSE](https://github.com/adwardstark/go-rest-api/blob/master/LICENSE) file for details.