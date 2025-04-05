package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	pb "grpc-file-service/fileservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// Number of files to test
	numFiles = 1
	// Size of each file (1MB)
	fileSize = 1024 * 1024 * 1024
	// Folder to store dummy files
	folderName         = "../dummy_files"
	receivedFolderName = "../dummy_files_received"
)

func main() {
	absFolderPath, err := filepath.Abs(folderName)
	if err != nil {
		log.Fatalf("Failed to get absolute folder path: %v", err)
	}

	// Get the absolute path for the folder where received files will be stored
	receivedFolderPath, err := filepath.Abs(receivedFolderName)
	if err != nil {
		log.Fatalf("Failed to get absolute folder path for received files: %v", err)
	}

	// Create the directory to store dummy files (if not already created)
	if err := os.MkdirAll(receivedFolderPath, 0755); err != nil {
		log.Fatalf("Failed to create folder: %v", err)
	}

	// Connect to the gRPC server
	conn, err := grpc.NewClient("grpc-file-server:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//conn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	// Track total time
	var totalTime time.Duration

	var size int
	// Loop to send multiple files
	for i := 0; i < numFiles; i++ {
		// Generate the filename and its full path
		filename := fmt.Sprintf("testfile_%d.txt", i)
		file_path := filepath.Join(absFolderPath, filename)

		fmt.Println("Reading from server file " + file_path)
		// Measure the time taken to read and send the file
		startTime := time.Now()

		// Pass the full file path to the server
		stream, err := client.ReadFile(context.Background(), &pb.ReadFileRequest{Filename: file_path})
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		// Write the received stream data to a new file
		receivedFilePath := filepath.Join(receivedFolderPath, filename)
		outputFile, err := os.Create(receivedFilePath)
		if err != nil {
			log.Fatalf("Error creating output file: %v", err)
		}
		defer outputFile.Close()

		// Read the stream and write to the output file
		for {
			resp, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break // End of stream
				}
				log.Fatalf("Error receiving stream: %v", err)
			}
			response := resp.Content

			//fmt.Println(resp.Content)
			size = size+len(response)
			outputFile.Write(resp.Content)
		}

		// Calculate the time taken for this file and add it to total time
		elapsedTime := time.Since(startTime)
		totalTime += elapsedTime
		fmt.Printf("Processed %s in %v\n", filename, elapsedTime)
	}

	// Print the total time taken to process all files
	fmt.Printf("\nTotal time taken to process %d all files: %v\n", size, totalTime)
}

