# grpc-file-service

This application follows a server-client model built using gRPC with Go. The server listens on port 50053 for client requests, where clients can request the content of a specific file by providing the file name. It is assumed that the file is already mounted at the server's location. The server streams the file content back to the client.

### Clone this repo with below command
```git clone https://github.com/ssahoo11/grpc-file-service.git```

### Change the directory
```cd grpc-file-service/```

### Run below docker command to create server and client images
```docker build -f Dockerfile.server --no-cache -t grpc-file-server:latest . ```
```docker build -f Dockerfile.client --no-cache -t grpc-file-client:latest . ```

### Check new images
``` docker images ```

### Create a custom network for the communication
```docker create network grpc-network```

### Run the server with interactive mode to see the stdout logs
``` docker run -it --name grpc-file-server --network grpc-network -p 50053:50053 grpc-file-server:latest```

### On a separate shell run the client container with interactive mode

```docker run -it --name grpc-client --network grpc-network -p 50053:50053 grpc-file-client:latest```


### Output of server container
```
$ docker run -it --name grpc-file-server --network grpc-network -p 50053:50053 ssahoo1d/grpc-file-server
creating file /app/dummy_files/testfile_0.txt with size 1073741824 bytes
2025/04/05 06:03:36 Server is running on port 50053...

```

### Output of client container
```
$ docker run -it --name grpc-client --network grpc-network 240fd80b1d0b
Reading from server file /app/dummy_files/testfile_0.txt
Processed testfile_0.txt in 1.24246659s

Total time taken to process 1073741824 size data: 1.24246659s
$

```

