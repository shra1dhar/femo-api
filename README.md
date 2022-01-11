##Set up Postgres container

1. Check available images `docker images`
2. If not available then download the alpine version (since it is light) `docker pull postgres:12-alpine`
3. To check running container use `docker ps`
4. For running container, use: `docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine`

-p <machine-port>:<container-port>
-e for setting environment variables
-d to run as daemon (background)

Also note, password is not asked when we run postgres container locally

To get a interactive shell in the container, use:
`docker exec -it postgres12 psql -U root`

Commands psql, `select now();` to verify access.

To exit console, use: `\q` 

## For Docker logs
`docker logs <container_name_or_id>`

To check logs for container we created:
`docker logs postgres12`

----

For schema use: `https://dbdiagram.io/home`
For managing DB use: `TablePlus`



