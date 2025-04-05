package main

import (
	"bufio"
	"log"
	"net"
	"os"

	pb "GRPC-POC/fileservice"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFileServiceServer
}

func (s *server) ReadFile(req *pb.ReadFileRequest, stream pb.FileService_ReadFileServer) error {
	filename := req.Filename
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024*1024)

	for {
		n, err := reader.Read(buffer)
		if n > 0 {
			if err := stream.Send(&pb.ReadFileResponse{Content: buffer[:n]}); err != nil {
				return err
			}
		}
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFileServiceServer(grpcServer, &server{})
	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
