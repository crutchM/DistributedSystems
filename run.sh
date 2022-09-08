docker build -t go-app .
docker run --name=web-app -p 8080:8080 go-app