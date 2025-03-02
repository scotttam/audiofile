package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

func main() {
	// Define command-line flags
	var (
		dirPath    string
		workers    int
		quietMode  bool
		dryRun     bool
		recursive  bool
	)

	flag.StringVar(&dirPath, "dir", "", "Directory containing FLAC files to convert")
	flag.IntVar(&workers, "workers", runtime.NumCPU(), "Number of concurrent workers")
	flag.BoolVar(&quietMode, "quiet", false, "Reduce output verbosity")
	flag.BoolVar(&dryRun, "dry-run", false, "Show what would be converted without actually converting")
	flag.BoolVar(&recursive, "recursive", true, "Recursively search for FLAC files in subdirectories")
	flag.Parse()

	// Check if directory path is provided
	if dirPath == "" {
		fmt.Println("Please provide a directory path using the -dir flag")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check if the directory exists
	info, err := os.Stat(dirPath)
	if err != nil {
		fmt.Printf("Error accessing directory: %v\n", err)
		os.Exit(1)
	}
	if !info.IsDir() {
		fmt.Printf("%s is not a directory\n", dirPath)
		os.Exit(1)
	}

	// Find all FLAC files in the directory
	flacFiles, err := findFlacFiles(dirPath, recursive)
	if err != nil {
		fmt.Printf("Error finding FLAC files: %v\n", err)
		os.Exit(1)
	}

	if len(flacFiles) == 0 {
		fmt.Printf("No FLAC files found in %s\n", dirPath)
		os.Exit(0)
	}

	if !quietMode {
		fmt.Printf("Found %d FLAC files to convert\n", len(flacFiles))
	}

	if dryRun {
		fmt.Println("Dry run mode - the following files would be converted:")
		for _, file := range flacFiles {
			aiffFile := strings.TrimSuffix(file, filepath.Ext(file)) + ".aiff"
			fmt.Printf("%s -> %s\n", file, aiffFile)
		}
		os.Exit(0)
	}

	// Create a channel to limit concurrency
	semaphore := make(chan struct{}, workers)
	var wg sync.WaitGroup

	// Convert FLAC files to AIFF
	for _, flacFile := range flacFiles {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			
			// Acquire a token from the semaphore
			semaphore <- struct{}{}
			defer func() { <-semaphore }() // Release the token when done
			
			err := convertFlacToAiff(file, quietMode)
			if err != nil {
				fmt.Printf("Error converting %s: %v\n", file, err)
			} else if !quietMode {
				fmt.Printf("Successfully converted %s\n", file)
			}
		}(flacFile)
	}

	wg.Wait()
	if !quietMode {
		fmt.Println("All conversions completed")
	}
}

// findFlacFiles finds all FLAC files in the given directory
func findFlacFiles(dirPath string, recursive bool) ([]string, error) {
	var flacFiles []string

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		
		// Skip subdirectories if not recursive
		if !recursive && info.IsDir() && path != dirPath {
			return filepath.SkipDir
		}
		
		if !info.IsDir() && strings.ToLower(filepath.Ext(path)) == ".flac" {
			flacFiles = append(flacFiles, path)
		}
		return nil
	}

	err := filepath.Walk(dirPath, walkFunc)
	return flacFiles, err
}

// convertFlacToAiff converts a FLAC file to AIFF format using FFmpeg
func convertFlacToAiff(flacFile string, quietMode bool) error {
	// Create the output AIFF file path by replacing the extension
	aiffFile := strings.TrimSuffix(flacFile, filepath.Ext(flacFile)) + ".aiff"

	// Use FFmpeg to convert the file
	args := []string{"-i", flacFile, "-f", "aiff"}
	
	// Add quiet flag if in quiet mode
	if quietMode {
		args = append(args, "-loglevel", "error")
	}
	
	args = append(args, aiffFile)
	cmd := exec.Command("ffmpeg", args...)
	
	if !quietMode {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	return cmd.Run()
} 