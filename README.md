# txtarer

`txtarer` is a CLI tool that recursively reads files in a specified directory and combines them into a single file in txtar format. When working with a Git repository, it offers an option to exclude files according to the `.git` directory and `.gitignore` file rules.

## Features

- Recursively reads and combines files in a specified directory into a txtar format.
- The `-gitmode` option enables exclusion of files based on the `.gitignore` file in a Git repository.
- Output filename can be specified with the `-output` option. If not specified, the default output file is "output.txtar" in the current directory.

## Installation

Ensure you have a Go language environment set up. Then, you can install the tool using the following command:

```
go install github.com/uji/txtarer@latest
```

## Usage

Basic usage:

```
txtarer [options] <directory>
```

### Options

- `-output <file>`: Specifies the name of the output file. Default is "output.txtar".
- `-gitmode`: Excludes files based on the `.gitignore` file in a Git repository.
- `-help`: Displays the help message.

### Examples

- To combine files from a regular directory:

  ```
  txtarer -output combined.txtar /path/to/directory
  ```

- To combine files from a Git repository:

  ```
  txtarer -gitmode -output combined.txtar /path/to/git-repository
  ```
