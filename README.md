# Tailwind CSS with Go templates
These are my experiments is trying to use `go` templates
as web components with tailwind CSS styling.

Potentially combining the best of both the `go` and `nodejs` worlds.

`go` is robust, lean on resources and fast, but lacks
a healthy frontend development community.

`nodejs` has a very healthy frontend development community but, I feel,
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

I made my `go` web development experience quicker by installing
`gow` -- go watcher from https://github.com/mitranim/gow

The following is my gow watch command:
```
cd hello/cmd/hello
gow -v -e=go -e=mod -e=css -e=html run .
```
