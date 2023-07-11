# Gorilla Mux Learning

Based on <https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql>

Basic API with testing, running on PostgreSQL

## Docker PostgreSQL instance

Postgres is used as the database, so it's handy to have one in Docker that we can discard at will. The
task file contains scripts to run tests and reset the test environment afterwards. Some tests are destructive
and this test suite is not suitable for running against a production environment.

`docker run --name TestDonkey -it -p 5432:5432 -e POSTGRES_PASSWORD=testpassword -e POSTGRESS_DB=testgres -d postgres`
