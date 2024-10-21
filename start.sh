sudo docker rmi -f theilich/pm:latest
sudo docker rmi -f theilich/go_db:latest

sudo docker build -t theilich/pm:latest -f app .
sudo docker build -t theilich/go_db:latest -f db .

sudo docker login

sudo docker push theilich/pm:latest
sudo docker push theilich/go_db:latest

sudo docker compose down -v 
sudo docker compose up




