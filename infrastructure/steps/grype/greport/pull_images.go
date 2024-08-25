package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	// Create a Docker client
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.WithVersion("1.45"), client.FromEnv)
	if err != nil {
		panic(err)
	}

	// Check if the Docker socket is mounted (if using Docker socket binding)
	if _, err := os.Stat("/var/run/docker.sock"); err != nil {
		fmt.Println("Docker socket is not mounted. Please ensure it's accessible.")
		return
	}

	// Create the logs directory if it doesn't exist
	if err := os.MkdirAll("./logs", os.ModePerm); err != nil {
		fmt.Printf("Error creating logs directory: %v\n", err)
		return
	}

	// Open the log file for writing
	logFile, err := os.OpenFile("./logs/grype_scan.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}
	defer logFile.Close() // Ensure log file is closed

	// Redirect stdout and stderr to the log file
	cmdOut := io.MultiWriter(os.Stdout, logFile)

	// Get a list of running containers
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		fmt.Printf("Error getting container list: %v\n", err)
		return
	}

	// Iterate through each container and inspect it to get the image ID
	for _, container := range containers {
		fmt.Fprintf(cmdOut, "Inspecting container: %s\n", container.ID)

		// Inspect the container to get the image ID
		inspect, err := cli.ContainerInspect(ctx, container.ID)
		if err != nil {
			fmt.Fprintf(cmdOut, "Error inspecting container: %v\n", err)
			continue
		}

		// Scan the image using grype
		cmd := exec.Command("grype", inspect.Image)
		cmd.Stdout = cmdOut // Redirect grype output to combined writer
		cmd.Stderr = cmdOut // Redirect grype error output to combined writer

		err = cmd.Run() // Run the command
		if err != nil {
			fmt.Fprintf(cmdOut, "Error scanning image: %v\n", err)
			continue
		}
	}
}