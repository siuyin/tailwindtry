# Tailwind CSS with Go templates
These are my experiments is trying to use `go` templates
as web components with tailwind CSS styling.

Here is combine the best of both the `go` and `nodejs` worlds.

`go` is robust, lean on resources and fast, but lacks
a healthy fontend development community.

`nodejs` has a very healthy fontend development community but, I feel,
lacks the software engineering discipline of `go`

## Generate tailwind css 
```
npx tailwindcss -i ./src/styles.css -o ./cmd/hello/static/styles.css --watch
npx tailwindcss -i ./src/styles.css -o ./cmd/hello/static/styles.min.css --watch --minify

```

## Run the go web server
```
cd hello/cmd/hello
go run main.go // serves to port 8080
```
