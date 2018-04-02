# remove images and containers
docker image prune -f
docker rmi -f imgserver
docker rm -f imgserver
# build docker file
docker build -t  imgserver .
# run container
docker run --name imgserver \
    --link mongodb-img:mgodb \
    --env-file env \
    -p 8090:8090 \
    -v "$(pwd)":/go/src/github.com/ivancduran/imgserver \
    -d imgserver
