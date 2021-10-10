env GOOS=linux GOARCH=arm64 go build -ldflags="-w -s" -o target/caseconv_arm64
docker build --build-arg ARCHPLATFORM=arm64 -t caseconv:arm64 .
docker save caseconv:arm64 > dockerimages/caseconv-arm64-docker.tar.gz