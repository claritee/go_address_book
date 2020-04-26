# Project setup

Make sure GOPATH is set e.g. `~/Projects/golang`

Repo needs to be in `$GOPATH/src` e.g. `~/Projects/golang/src/go_address_book`

# Dependencies

```
dep ensure
```

```
go mod vendor
```

# Running the CLI App

Go to directory

```
cd src/cli
```

To run by default

```
go run main.go
```

To run specifying another file to parse

```
go run main.go file-to-parse
```

The other usages

```
go run main.go people.csv oldest
go run main.go people.csv youngest
go run main.go people.csv city Melbourne
go run main.go people.csv name Lily
```

# Running the webapp

```
go build
go ./api
```

Then
```
curl http://localhost:8080
```

With a path

```
curl http://localhost:8080/hello
```

```
curl http://localhost:8080/bird
```

To create a bird
```
http://localhost:8080/assets/
```

## CREATE

```
psql
```

```
CREATE DATABASE bird_encyclopedia;
\list;
\c bird_encyclopedia;
```

Create test DB

```
CREATE DATABASE bird_encyclopedia_test;
```

### Create tables

```
CREATE TABLE birds (
  id SERIAL PRIMARY KEY,
  bird VARCHAR(256),
  description VARCHAR(1024),
  species VARCHAR(1024)
);
```
## OR Docker PG image

https://github.com/lib/pq/blob/master/TESTS.md#example-setup-docker

# Dependencies

Commands

```
go mod init
go mod tidy # clean up libs
go mod vendor # external libs go here
```

# Running tests

```
go test api/*
```

OR

```
cd api
go test . # or go test -v
```

OR (all)

```
go test ./...
```

# Resources

* Switch Case - https://yourbasic.org/golang/switch-statement/
* Errors - https://gobyexample.com/errors
* Errors - https://blog.golang.org/error-handling-and-go
* Errors - https://blog.golang.org/go1.13-errors
* webapp - https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/#2-set-up-your-directory-structure
* `dep` dependency management - https://github.com/golang/dep#setup
* go projects - https://medium.com/mindorks/create-projects-independent-of-gopath-using-go-modules-802260cdfb51
* webapp - https://www.sohamkamani.com/blog/2017/10/18/golang-adding-database-to-web-application/
* webapp repo - https://github.com/sohamkamani/blog_example__go_web_db
* go modules outside GOPATH - https://blog.francium.tech/go-modules-go-project-set-up-without-gopath-1ae601a4e868
* project layout - https://github.com/golang-standards/project-layout
* Tests - https://medium.com/@benbjohnson/structuring-tests-in-go-46ddee7a25c
* Example webapp https://github.com/sohamkamani/blog_example__go_web_app
* Example webapp with DB https://github.com/sohamkamani/blog_example__go_web_db
* Go webapp API https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
*