#docker rm -f project; docker run --name="project" -p 8080:8080 caseconv:amd64
docker pull ghcr.io/sajayjkurup/caseconv:snapshot
docker rm -f project; docker run --name="project" -p 8080:8080 ghcr.io/sajayjkurup/caseconv:snapshot