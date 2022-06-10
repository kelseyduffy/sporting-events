# sporting-events
website hobby project showing the different sporting events I've been to. 

PostgreSQL database, GO web server, and JS website (maybe Svelte on a future iteration).

# deploy docker postgres container

following [postgres docker wiki](https://hub.docker.com/_/postgres)

`docker run --name some-postgres -p 5432:5432 -e POSTGRES_PASSWORD=mysecretpassword -d postgres`

Connect:
- locally from 127.0.0.1:5432 with things like pgAdmin
- another container such as
  - `docker run -it --rm --link some-postgres:postgres postgres psql -h postgres -U postgres`

to stop and start the container:
`docker stop some-postgres`
`docker start some-postgres`