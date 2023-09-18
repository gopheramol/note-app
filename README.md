# note-app
A Simple and Efficient Web-Based Note-taking App in Go.

## Stack
### Go
### Mongo

## TODO : Dockerize , more features - Created time, suuport to store URL link


# Installtion 

## Up the the MongoDB Container 

```                   
  docker run -d --name note-mongodb \          
  -p 27017:27017 \
  -e MONGO_INITDB_ROOT_USERNAME=root \
  -e MONGO_INITDB_ROOT_PASSWORD=password \
  mongo
``````
## Run the application

```go run main.go```

## Browse localhost

``` http://localhost:8081/ ```

## Screenshot
