# gRPC File Service

## Overview

The `grpc-file-service` application is built using gRPC with Go, following a server-client model. The server listens on port `50053` for client requests, allowing clients to request the content of specific files by providing the file name. It is assumed that the files are already mounted at the server's location. The server streams the file content back to the client.

## Getting Started

### Clone the Repository

To get started, clone the repository using the following command:

```bash
git clone https://github.com/ssahoo11/grpc-file-service.git
```

### Change Directory

Navigate to the project directory:

```bash
cd grpc-file-service/
```

### Build Docker Images

Run the following commands to create the server and client Docker images:

```bash
docker build -f Dockerfile.server --no-cache -t grpc-file-server:latest .
docker build -f Dockerfile.client --no-cache -t grpc-file-client:latest .
```

### Check New Images

You can check the newly created images with:

```bash
docker images
```

### Create a Custom Network

Create a custom network for communication between the server and client:

```bash
docker create network grpc-network
```

### Run the Server

Run the server in interactive mode to see the stdout logs:

```bash
docker run -it --name grpc-file-server --network grpc-network -p 50053:50053 grpc-file-server:latest
```

### Run the Client

In a separate shell, run the client container in interactive mode:

```bash
docker run -it --name grpc-client --network grpc-network grpc-file-client:latest
```

## Example Outputs

### Server Container Output

```bash
$ docker run -it --name grpc-file-server --network grpc-network -p 50053:50053 grpc-file-server
creating file /app/dummy_files/testfile_0.txt with size 1073741824 bytes
2025/04/05 06:03:36 Server is running on port 50053...
```

### Client Container Output

```bash
$ docker run -it --name grpc-client --network grpc-network grpc-file-client
Reading from server file /app/dummy_files/testfile_0.txt
Processed testfile_0.txt in 1.24246659s

Total time taken to process 1073741824 size data: 1.24246659s
$
```

```bash
$ docker run -it --network grpc-network --name grpc-file-java-client grpc-file-java-client
Time taken to read and write the file of size 1073741824 is : 2.662626447 seconds.
$
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to open issues or submit pull requests for any improvements or bug fixes.
```

Feel free to adjust any sections or add additional information as needed!
