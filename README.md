# Tailwind CSS with Go templates
These are my experiments in trying to use `go` templates
as web components with tailwind CSS styling.
Thus potentially combining the best of both `go` and `nodejs` worlds.

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
The http and websocket servers have been updated to use transport layer security (TLS).
place server.pem and server-key.pem at the `hello/cmd/hello` folder.
Other files and static assest have been embedded with the go binary using `embed` https://pkg.go.dev/embed

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

## Running the github user search api example
Create and user a github token if you see a "limits exceeded"
response from github.

By default I make an unauthenticated request to the github API.
This is rather low rate limit levels.


```
export GITHUB_TOKEN=<paste your token here>
go run main.go
curl -L localhost:8080/api/v1/github/{username}

eg.
curl -L localhost:8080/api/v1/github/siuyin
