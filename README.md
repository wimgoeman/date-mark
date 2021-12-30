# date-mark
Adds date from exif tags to images in folder

- Scans a folder for image files
- Reads creation date from exif tags
- Runs [Image Magick](https://imagemagick.org/) to add the date to the image itself

It would be probably be more efficient to use the Wand API to add the date, but for this small tool, the compiling/linking/releasing/... would be too much hassle.

Make sure to have ImageMagick installed and available on the PATH before running this tool.