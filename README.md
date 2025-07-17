# finder

a command line tool for

- searching and finding text files in a linux directory.
- get line count of files.
- get top 10 most frequent words.

## Understanding Code

To understand the code structure we have 2 main packages -> cli, files.

- main.go is our starting point to read the code.
  - we try to match the command given by user and handle logic of printing the result required by user.
- files package handles scanning files and directories and operations performed in them.
  - word counter, it should be executed as a go routine, it counts words in file and sends them over a channel.
  - line counter, it should be executed as a go routine, it counts number of lines in a file and sends them over a channel in a specific format.
  - file scanner, it scans all the files in a given path along with files in the sub-directory.
  - file handler is an abstraction to main.go for handling filescanner, wordcounter, linecounter functions.
- cli package as the name suggests it manages command line. I have used kong package here.
- utils package for some utility functions
- models package for storing all the models

## Execution

- I have written a make file to install this tool to your linux system.
- Run make command with sudo permissions since it should install to /usr/bin/

  ``` bash
  sudo make build_and_install
  ```

- To run the tool use the following command
  
  ``` bash
  sudo finder list <path> --recursive --word-count
  ```

- For help and exploring tool
  
  ``` bash
  sudo finder list -h
  ```

## Testing

- Unit Tests are written for each service in files package
- command for testing
- ```go test ./...```
