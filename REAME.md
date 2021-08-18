# Book management software
This program allows to interact with a book management software. The repo contains the database schema definition, a REST server to interact with the database and a cli program to manage book resources and collections of books. The first prototype implements book creation, update e retrieval.

## Getting started
To bootstrap the application follow those steps:
- run `make run-mysql-docker` to start a docker container running MySQL.
- After few seconds `make load-tables` loads the database schema
- To bootsrap the REST server run `make server`
- Finally, build the `book-cli` utility with `make build-cli`. The compiled file can be found in `pkg/book-cli`.

Additional information can be found at:
- [API Design](./pkg/apis/README.md)
- [CLI Design](./pkg/book-cli/README.md)
- [DB Design](./pkg/db/README.md)
  