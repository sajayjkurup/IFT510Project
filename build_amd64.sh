env GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o target/caseconv_amd64
docker build --build-arg ARCHPLATFORM=amd64 -t caseconv:amd64 .
docker save caseconv:amd64 > dockerimages/caseconv-amd64-docker.tar.gz