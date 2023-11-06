docker compose down && docker compose build && docker compose up -d
docker rmi $(docker images -f "dangling=true" -q) #Removes dangling docker images
