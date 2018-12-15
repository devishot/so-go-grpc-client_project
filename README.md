# so-go-grpc-client_project
Provides Client / Project relations

### Interfaces

#### gRPC

##### How to install submodule
1. Check `.gitmodules` file
2. Just clone a parent project

##### How to add submodules

Run command

`git submodule add --name [NAME] [repository] [path]`

Example:

> git submodule add --name so-gRPC-proto-files git@github.com:devishot/grpc-protofiles.git domain_interface/grpc/protofiles

---

### Programming Patterns and more

1. SQL: Fully avoid using NULL in any field. 

    > Nullable columns are annoying and lead to a lot of ugly code. If you can, avoid them.
    
    from `database/sql` [tutorial](http://go-database-sql.org/nulls.html).
    
    How to manage data without NULL and make JOINs see:
    
    - [Discussion](https://stackoverflow.com/questions/3079885/options-for-eliminating-nullable-columns-from-a-db-model-in-order-to-avoid-sql#)
    
    - Book: `O’Relly / C.J. Date: "Database in Depth: Relational Theory for Practitioners”`
    
    - [Paper 1: Critique for this approach](http://www.u.arizona.edu/~rubinson/scrawl/Rubinson.2007.Nulls_Three-Valued_Logic_and_Ambiguity_in_SQL.pdf)
    
    - [Paper 2: How to manage database without NULL value](https://www.dcs.warwick.ac.uk/~hugh/TTM/Missing-info-without-nulls.pdf)

2. 


---

### Server configurations

#### Database

Every microservices have own database.
For dev environment, on local machine, they have single user.

##### Install
- Postgresql 11 for Macos
https://postgresapp.com/
- PgAdmin GUI for postgresql databases. Runs as local web server
https://www.pgadmin.org/

##### Create database

run in terminal:
> psql postgres

1. for create database `so_client_project`:

    > postgres=> CREATE DATABASE so_client_project;

2. for grant all access for db user `devishot`, which same as system user:

    > postgres=> GRANT ALL PRIVILEGES ON DATABASE so_client_project TO devishot;

3. for connect into created database:

    > postgres=> \connect so_client_project;

4. for display tables in the database:

    > postgres=> \dt

5. for exit:

    > \q


#### Configurations
Store configs in env variables. It replaces config files.

- [Idea from "12 Factor App"](https://12factor.net/config)

- [Tutorial about ENV and ARG in docker](https://vsupalov.com/docker-arg-env-variable-guide/)


##### Run with env-files

How to execute multiple env-file with `docker run`:

> docker run --env-file=database.env --env-file=app.env alpine env


How to define in `docker-compose` file:

```dockerfile
version: '3'

services:
  plex:
    image: alpine
      env_file: 
        - ./database.env 
        - ./app.env
        - /opt/secrets.env
```


#### Docker

##### Build:

> docker build -t so-client_project --force-rm .


`--force-rm` always delete intermediate containers; 
you can replace it to `--rm` which delete only when build was success.

##### Run:
> docker run -it -p 8080:8080 so-client_project

##### Debug:

Getting inside a container:
> docker run -it -p 8080:8080 so-client_project /bin/sh

