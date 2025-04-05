# Use a minimal base image
FROM mcr.microsoft.com/windows/servercore:ltsc2019

# Set the working directory
WORKDIR /app

# Copy the client executable
COPY client.exe .

# Command to run the client
CMD ["client.exe"]
