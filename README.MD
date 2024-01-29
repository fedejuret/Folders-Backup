# Folder Backup Script

A simple Go script for recursively backing up files from a source directory to a destination directory, retaining the folder structure, and copying only files created more than a specified number of months ago.

## Usage

```bash
files-backup-golang -origin /path/to/source -destination /path/to/destination -months 6 -delete-originals
```

- -origin: Source directory for backup.
- -destination: Destination directory for backup.
- -months: Number of months ago to consider for file backup (default is 6 months).
- -delete-originals: Optional flag. If provided, files will be deleted from the source directory after being successfully copied to the destination.

## Features

- Recursively backs up files from the source directory to the destination directory.
- Retains the original folder structure in the destination directory.
- Copies only files created more than the specified number of months ago.
- Option to delete original files from the source directory after successful backup.

# Author
Federico Juretich fedejuret@gmail.com

# Acknowledgments

This script utilizes the following third-party packages:

- github.com/common-nighthawk/go-figure for ASCII art text display.
- github.com/fatih/color for colorful console output.
