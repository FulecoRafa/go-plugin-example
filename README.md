# Go plugin example

This is a simple application that has only one job: show todos from a JSON file

However, we could have it display extra fields! So I wrote a plugin to add
priority to all todos displayed!

In the end, is just a simple example on how to create and import [go plugins](https://pkg.go.dev/plugin).

## Building
- Have Go
- *Make* your life easier
- `make` should build the plugin
- `go run .` will run the application
