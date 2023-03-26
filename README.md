# Tailwind CSS with Go templates

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
