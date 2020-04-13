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
go run src/api/main.go
```

Then
```
curl http://localhost:8080
```

# Resources

* Switch Case - https://yourbasic.org/golang/switch-statement/
* Errors - https://gobyexample.com/errors
* Errors - https://blog.golang.org/error-handling-and-go
* Errors - https://blog.golang.org/go1.13-errors
* webapp - https://www.sohamkamani.com/blog/2017/09/13/how-to-build-a-web-application-in-golang/#2-set-up-your-directory-structure