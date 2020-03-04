Requirements
---------------
* Docker
 ````dockerfile
docker-py > 3.6.0
docker-compose > 1.23.2
````

* godotenv
````go
go get github.com/joho/godotenv
````

* grpc
````go
go get google.golang.org/grpc
````

Install
---------------
````go
go get github.com/tarikbauer/go_vs_py_benchmark
````

Servers
--------------
* All servers work asynchronously;

* All servers only have one API: ````/api?t={values}````, `values` is a list of integers joined by a comma, such as `/api?t=1,2,3,1`;

* For each integer value of the query param received, the server mocks an IO asynchronously requisition, that takes the correspondent value in milliseconds;

* The project compare three different servers, written in different languages and using different protocols of communication.
---
1. Written in Python and developed using Sanic (using `uvloop`) as framework (REST);
2. Written in GO and developed using Mux as framework (REST);
3. Written in GO and developed using GRPC.

Client
--------------
* To modify the number of requisitions made for every server, you just need to change the global variable `ITERATIONS` on main.go;

* To modify the query param `t`, you just need to change the global variable `VALUES` on main.go.

Run
--------------

* Running the servers
````dockerfile
docker-compose build
docker-compose up
````

* Running the benchmark
````go
go run main.go
````
 