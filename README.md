# so-go-grpc-client_project
Provides Client / Project relations

### Domain Interfaces

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




#### App Configurations
Store configs in env variables. It replaces config files.

- [Idea from "12 Factor App"](https://12factor.net/config)

- [Tutorial about ENV and ARG in docker](https://vsupalov.com/docker-arg-env-variable-guide/)


##### Run with env-files

How to execute multiple env-file with `docker run`:

> docker run --env-file=database.env --env-file=app.env alpine env


How to define in `docker-compose` file:

```
version: '3'

services:
  plex:
    image: alpine
      env_file: 
        - ./database.env 
        - ./app.env
        - /opt/secrets.env
```

##### Default Environments

is under folder default_env:
- default_env/database.env


---

#### Docker

2 version of build:

- build/Dockerfile - use two-step build for optimize output image
- build/dev/Dockerfile - with hot-reload, executes `go build` every time after code changes


##### First version

###### Build:

> docker build -t so-client_project --rm -f build/Dockerfile .

1. `--force-rm` always delete intermediate containers; 
you can replace it to `--rm` which delete only when build was success;
2. `-f` specify Dockerfile independently of the build context (workdir on host).

###### Run:
> docker run -it -p 8080:8080 --env-file=default_env/database.env so-client_project

###### Debug:

Getting inside a container:
> docker run -it -p 8080:8080 so-client_project /bin/sh


##### Second version

There is special go library for watch files and compile go sources:
 
- [github.com/githubnemo/CompileDaemon](github.com/githubnemo/CompileDaemon)

- [Article](https://www.zachjohnsondev.com/posts/go-docker-hot-reload-example/)

###### Build

> docker build -t so-client_project-dev --rm -f build/dev/Dockerfile .

###### Run:

> docker-compose -f build/dev/docker-compose.yml up

---

##### Problem / Solutions:


###### Alpine cannot resolve host for download APKINDEX.

Docker file runs `apk update` for install git, 
which is required for install some `go get` packages.


###### Solution:

run following:

> sudo networksetup -setdnsservers Wi-Fi 8.8.8.8 8.8.4.4 74.82.42.42

why it works:

- [Issue for the problem](https://github.com/gliderlabs/docker-alpine/issues/279)
- [Official DNS for alpine](https://wiki.alpinelinux.org/wiki/Configure_Networking#Configuring_DNS)
- [Solution how create /etc/resolv on macOS](https://serverfault.com/a/478540)



###### Cannot connect to localhost of Docker Host from container

This problem relates to Docker on Mac.

###### Solution

> From 18.03 onwards our recommendation is to connect to the special DNS name 
> `host.docker.internal`, which resolves to the internal IP address used by the host. 
> This is for development purpose and will not work in a production environment outside of Docker for Mac.

[Official docs](https://docs.docker.com/docker-for-mac/networking/#there-is-no-docker0-bridge-on-macos#i-want-to-connect-from-a-container-to-a-service-on-the-host)

