# CASE CONV
Convert a string to different standard case formats like CAMEL Case, SNAKE CASE etc


## Build Porject
```
./build.sh
```

By default, this will cross compuile the GO project for linux amd64 and arm64 architectures. This will also build the docker containers for the same. 

## Run Porject
```
./run.sh
```
This will run the docker container on port 8080

# API

### Get Supported Formats
```
GET http://localhost:8080/supportedformats 
```

### Convert Case
```
GET http://localhost:8080/Query%20String/convert?format=KEBAB

Response example:
{"input":"Query String","output":"query-string","format":"KEBAB"}
```


# References
  - https://chemidy.medium.com/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324 

  - https://github.com/iancoleman/strcase

  - https://hub.docker.com/_/scratch

  - https://codeburst.io/docker-from-scratch-2a84552470c8

  - https://www.docker.com/blog/multi-arch-build-and-images-the-simple-way/

  - https://docs.docker.com/buildx/working-with-buildx/


