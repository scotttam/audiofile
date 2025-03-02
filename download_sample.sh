#!/bin/bash

# Create test directories
mkdir -p test_dir/subdir

# Download a sample FLAC file from the Internet Archive
# This is a Creative Commons licensed audio file
echo "Downloading sample FLAC file..."
curl -L "https://ia800504.us.archive.org/15/items/TenD2005-07-16.flac16/TenD2005-07-16t10Wonderboy.flac" -o test_dir/sample.flac
curl -L "https://ia800504.us.archive.org/15/items/TenD2005-07-16.flac16/TenD2005-07-16t09Tribute.flac" -o test_dir/subdir/sample2.flac

echo "Sample FLAC files downloaded to test_dir/"
echo "You can now test the converter with:"
echo "./flac2aiff -dir test_dir" 