import os
import soundfile as sf
from pathlib import Path

def convert_flac_to_aiff(input_dir):
    # Get all FLAC files in the directory
    flac_files = list(Path(input_dir).glob('*.flac'))
    
    if not flac_files:
        print("No FLAC files found in the directory.")
        return

    print(f"Found {len(flac_files)} FLAC files.")
    
    for flac_file in flac_files:
        try:
            # Read the FLAC file
            data, samplerate = sf.read(str(flac_file))
            
            # Create output filename (same name, different extension)
            output_file = flac_file.with_suffix('.aiff')
            
            print(f"Converting: {flac_file.name}")
            
            # Write the AIFF file
            sf.write(str(output_file), data, samplerate, format='AIFF')
            
            print(f"Successfully converted to: {output_file.name}")
            
        except Exception as e:
            print(f"Error converting {flac_file.name}: {str(e)}")

if __name__ == "__main__":
    import sys
    
    if len(sys.argv) != 2:
        print("Usage: python3 flac_converter.py <directory_path>")
        sys.exit(1)
        
    directory = sys.argv[1]
    convert_flac_to_aiff(directory)