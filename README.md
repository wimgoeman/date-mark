# date-mark
Adds date from exif tags to images in folder

- Scans a set of folders for image files
- Reads creation date from exif tags
- Runs [Image Magick](https://imagemagick.org/) to add the date to the image itself

It would be probably be more efficient to use the Wand API to add the date, but for this small tool, the compiling/linking/releasing/... would be too much hassle.

## Installation

- Compile, and put the resulting binary somewhere on the system.
- Install Image Magick on the system.
- Designate a working directory for date-mark, and inside it, create date-mark.yml, with the following structure
```yaml
magick: path/to/magick/executable
dirs:
  - path: path/to/a/folder/to/process
    logPath: path/to/a/file/for/transaction/logging
  - path: path/to/a/folder/to/process
    logPath: path/to/a/file/for/transaction/logging
```

## Running

Run the date-mark executable, and make sure to give it the correct working directory.
date-mark will create a general log file inside the working directory.
For each configured directory, date-mark will keep a transaction log at the configured logPath.
This transaction log will ensure the same file is never processed twice, and also contains potential processing errors for files.
It is structured as CSV.
