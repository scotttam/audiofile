# FLAC to AIFF Converter

A command-line application written in Go that converts FLAC audio files to AIFF format.

## Features

- Recursively scans a directory for FLAC files
- Converts each FLAC file to AIFF format
- Preserves the original FLAC files
- Processes files concurrently for faster conversion
- Configurable number of concurrent workers
- Dry-run mode to preview conversions
- Quiet mode for reduced output

## Requirements

- Go 1.16 or higher
- FFmpeg (must be installed and available in your PATH)

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/scotttam/audiophile.git
   cd audiophile
   ```

2. Build the application:
   ```
   go build -o flac2aiff
   ```

## Usage

Run the application with the `-dir` flag to specify the directory containing FLAC files:

```
./flac2aiff -dir /path/to/your/flac/files
```

### Command-line Options

| Flag | Default | Description |
|------|---------|-------------|
| `-dir` | | Directory containing FLAC files to convert (required) |
| `-workers` | Number of CPU cores | Number of concurrent conversion processes |
| `-quiet` | false | Reduce output verbosity |
| `-dry-run` | false | Show what would be converted without actually converting |
| `-recursive` | true | Recursively search for FLAC files in subdirectories |

### Examples

Convert all FLAC files in a directory:
```
./flac2aiff -dir ~/Music/Albums
```

Convert files with 4 concurrent workers:
```
./flac2aiff -dir ~/Music/Albums -workers 4
```

Preview which files would be converted without actually converting them:
```
./flac2aiff -dir ~/Music/Albums -dry-run
```

Convert files with minimal output:
```
./flac2aiff -dir ~/Music/Albums -quiet
```

Convert only files in the specified directory (not in subdirectories):
```
./flac2aiff -dir ~/Music/Albums -recursive=false
```

## How It Works

The application uses FFmpeg to perform the actual audio conversion. It spawns multiple goroutines to process files concurrently, with the number of simultaneous conversions limited by the `-workers` flag.

## License

MIT 