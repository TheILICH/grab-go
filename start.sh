docker rmi -f theilich/pm:latest
docker rmi -f theilich/go_db:latest

docker build -t theilich/pm:latest -f app .
docker build -t theilich/go_db:latest -f db .




