# FLAC to AIFF Converter

A simple GUI application to convert FLAC audio files to AIFF format.

## Requirements
- Python 3.7 or higher
- pip (Python package installer)
- libsndfile (system dependency)

### Installing Requirements on Mac

1. Install Homebrew (if not already installed):
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

2. Install Python and libsndfile using Homebrew:
```bash
brew install python libsndfile
```

3. Verify the installation:
```bash
python3 --version
pip3 --version
```

## Installation

1. Clone this repository
2. Install the required dependencies:
```bash
pip3 install -r requirements.txt
```

## Usage

1. Run the converter:
```bash
python3 flac_converter.py
```

2. When prompted, enter the directory path containing your FLAC files
3. The converter will:
   - Find all FLAC files in the specified directory
   - Convert each file to AIFF format in the same directory
   - Display progress in the terminal

## Notes
- The converter maintains the original audio quality
- The output files will have the same names as the input files with .aiff extension
- Progress and any errors are shown in the terminal