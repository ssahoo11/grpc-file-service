# grpc-file-service

git clone https://github.com/ssahoo11/grpc-file-service.git

cd grpc-file-service/

docker build -f Dockerfile.server --no-cache -t grpc-file-server:latest .
docker build -f Dockerfile.client --no-cache -t grpc-file-client:latest .

docker create network grpc-network

docker run -it --name grpc-file-server --network grpc-network -p 50053:50053 grpc-file-server:latest

#On a seperate shell run
docker run -it --name grpc-client --network grpc-network -p 50053:50053 grpc-file-client:latest


