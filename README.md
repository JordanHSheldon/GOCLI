# GOCLI
Tool for managing directories for users

## Requirements

Write a command line utility in Go ( or your favorite language ) which takes as arguments a list of directories. The program should output the sizes of each of the individual directories passed as well as a cumulative total. If a --recursive flag is provided, output the sizes of each of the individual directories passed and sub-directories recursively as well as a cumulative total. If a "--human" flag is passed, format the sizes to be human friendly by outputting the size in the most appropriate unit of bytes. For example, 304K for 304,000 bytes and 300M for 300000000 bytes.

- [ ] Takes list of directories as arguements
- [ ] Outputs file sizes of each of individual directories passed as well as total
- [ ] --recursive flag is provided, output the sizes of each subdirectory and total
- [ ] --human flag format the file sizes to be more human readable.


## considerations:

1. files being bigger than int64 maxvalue do not work as they have potential to show in complete value
    - Could make slice of all values before they become bigger than max value and show that to user instead of the negative value.
2. Does not work currently to indicate the size of a singular file.
