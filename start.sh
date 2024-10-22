sudo docker compose down

sudo docker rmi -f theilich/pm:latest

sudo docker build -t theilich/pm:latest -f app .

sudo docker login

sudo docker push theilich/pm:latest

sudo docker compose up




