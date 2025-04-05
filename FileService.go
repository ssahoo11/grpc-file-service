package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"fmt"
	"strconv"
	"io/ioutil"
	"path/filepath"
	pb "grpc-file-service/fileservice"

	"google.golang.org/grpc"
)

const (
	// Number of files to test
	numFiles = 1
	// Size of each file (1MB)
	fileSize = 1024 * 1024 * 1024
	// Folder to store dummy files
	folderName         = "dummy_files"
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
	
	absFolderPath, err := filepath.Abs(folderName)
	if err != nil {
		log.Fatalf("Failed to get absolute folder path: %v", err)
	}

	// Create the directory to store dummy files (if not already created)
	if err := os.MkdirAll(absFolderPath, 0755); err != nil {
		log.Fatalf("Failed to create folder: %v", err)
	}
	
		// Loop to send multiple files
	for i := 0; i < numFiles; i++ {
		// Generate the filename and its full path
		filename := fmt.Sprintf("testfile_%d.txt", i)
		file_path := filepath.Join(absFolderPath, filename)

		// Create a 1MB file with dummy data
		if err := createDummyFile(file_path, fileSize); err != nil {
			log.Fatalf("Error creating dummy file: %v", err)
		}
	}
	
	lis, err := net.Listen("tcp", ":50053")
	//lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterFileServiceServer(grpcServer, &server{})
	log.Println("Server is running on port 50053...")
	//log.Println("Server is running on port 50054...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// createDummyFile creates a dummy file with the specified size
func createDummyFile(filepath string, size int) error {
	data := make([]byte, size)
	for i := range data {
		data[i] = 'A' // Fill with dummy data (character 'A')
	}
	fmt.Println("creating file " + filepath + " with size " + strconv.Itoa(size) + " bytes")
	return ioutil.WriteFile(filepath, data, 0644)
}
