# Cache

**cache** This is a basic cache that leverages Redis commands to store and retrieve data from an in-memory data store. 

### Usage

Clone the repository and build the server.go file as well as the client.go file. 
You can run the server by specifying any port of your choice as the preferred port to listen on.
To run the client, set the corresponding server port as the port to communicate with the server on. 

```bash
go run server.go 8080
go run client.go localhost:8080
```
