# sporting-events
website hobby project showing the different sporting events I've been to. 

PostgreSQL database, GO web server, and JS website (maybe Svelte on a future iteration).

# deploy docker postgres container

following [postgres docker wiki](https://hub.docker.com/_/postgres)

`docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=<postgres pass> -d postgres`

then connect to the postgres container at `127.0.0.1:5432` either with:
- another container
    - `docker run -it --rm --link some-postgres:postgres postgres psql -h postgres -U postgres`
    - `create database <database name>` -- could just use the default postgres database
    - `\connect <database name>`
- pgadmin

then manually execute the sql files in the [sql folder](db/sql) to create the teams table

set the environment variables on machine running Go needed for Go to connect to the postgres instance
- `export POSTGRES_USER=<postgres user>`
- `export POSTGRES_PASSWORD=<postgres password>`
- `export POSTGRES_DB=<postgres database>`

# Setting up REST API

run `go run main.go` and access the rest api at `localhost:8080`

## inserting records

once up and running, use curl to interact with the REST API and postgres db
`curl -X POST http://localhost:8080/teams -H "Content-type: application/json" -d '{ "name": "Houston Dynamo", "founded_year": "2005", "dissolved_year": null, "sport": "soccer"}'`
`curl http://127.0.0.1:8080/teams`