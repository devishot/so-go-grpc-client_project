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
    
    - [Paper 1](http://www.u.arizona.edu/~rubinson/scrawl/Rubinson.2007.Nulls_Three-Valued_Logic_and_Ambiguity_in_SQL.pdf)
    
    - [Paper 2](https://www.dcs.warwick.ac.uk/~hugh/TTM/Missing-info-without-nulls.pdf)

2. 