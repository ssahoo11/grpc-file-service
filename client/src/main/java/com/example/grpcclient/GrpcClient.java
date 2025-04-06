package com.example.grpcclient;

import fileservice.FileServiceGrpc;
import fileservice.FileServiceGrpc.FileServiceBlockingStub;
import fileservice.FileServiceOuterClass.ReadFileRequest;
import fileservice.FileServiceOuterClass.ReadFileResponse;
import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

import java.io.FileOutputStream;
import java.io.IOException;
import java.io.OutputStream;
import java.util.Iterator;

public class GrpcClient {

    public static void main(String[] args) {
        String serverAddress = "grpc-file-server"; // Update with the server address
        int serverPort = 50053; // The port you defined in your server
        String outputFilePath = "outputfile.txt"; // Specify the output file path

        // Create a channel
        ManagedChannel channel = ManagedChannelBuilder.forAddress(serverAddress, serverPort)
                .usePlaintext()
                .build();

        // Create a stub
        FileServiceBlockingStub stub = FileServiceGrpc.newBlockingStub(channel);

        // Create a request for reading a file
        ReadFileRequest request = ReadFileRequest.newBuilder()
                .setFilename("/app/dummy_files/testfile_0.txt") // Specify the filename to read
                .build();

        // Start measuring time
        long startTime = System.nanoTime(); // or System.currentTimeMillis() for millisecond precision

        // Call the server and get the response
        Iterator<ReadFileResponse> responses = stub.readFile(request);
	long totalLen = 0;

        // Open output file stream to write the content
        try (OutputStream outputStream = new FileOutputStream(outputFilePath)) {
            // Print and write the file content (streamed)
            while (responses.hasNext()) {
                ReadFileResponse response = responses.next();
                byte[] content = response.getContent().toByteArray();
                
                // Write the content to the file
                outputStream.write(content);
		totalLen += content.length;
            }

            // Calculate the total time taken
            long endTime = System.nanoTime(); // or System.currentTimeMillis()
            long duration = endTime - startTime;
            double seconds = duration / 1_000_000_000.0;
            // Print the total time taken (in nanoseconds or milliseconds)
            System.out.println("Time taken to read and write the file of size " + totalLen +" is : " + seconds + " seconds.");

        } catch (IOException e) {
            System.err.println("Error writing to file: " + e.getMessage());
        }

        // Shutdown the channel
        channel.shutdown();
    }
}


