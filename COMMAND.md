# command docker

docker container create mongo:latest
docker container create --name mongoku mongo:latest
docker container create --name mongoku -p 2727:27017 mongo:latest
docker container start mongoku
docker container stop mongoku
docker container rm mongoku
docker container ls --all
docker container inspect nginx
docker container logs nginx
 
docker network create nginx-deployment
docker network ls
docker network connect nginx-deployment nginx
docker network connect nginx-deployment mongo

docker-compose up -d
docker-compose down
docker-compose start
docker-compose stop

docker system df
