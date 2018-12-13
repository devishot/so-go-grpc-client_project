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

> git submodule add --name gRPC-proto-files git@github.com:devishot/grpc-protofiles.git interface/grpc/protofiles

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

