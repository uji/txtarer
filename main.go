package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/tools/txtar"
)

func main() {
	var output string
	var gitMode bool
	flag.StringVar(&output, "output", "output.txtar", "The output file to write the txtar archive to. Defaults to 'output.txtar' if not specified.")
	flag.BoolVar(&gitMode, "gitmode", false, "Enable git mode to exclude files based on .gitignore")
	help := flag.Bool("help", false, "Show help message")

	flag.Parse()

	if *help {
		fmt.Println("Usage of txtarer:")
		fmt.Println("  txtarer -output <output file> -gitmode <directory>")
		fmt.Println("The <directory> argument specifies the directory to combine into a txtar archive.")
		flag.PrintDefaults()
		return
	}

	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: txtarer -output <output file> -gitmode <directory>")
		os.Exit(1)
	}

	dir := flag.Arg(0)

	archive := &txtar.Archive{}
	var files []string

	if gitMode {
		cmd := exec.Command("git", "-C", dir, "ls-files")
		output, err := cmd.Output()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error running git ls-files: %v\n", err)
			os.Exit(1)
		}

		scanner := bufio.NewScanner(strings.NewReader(string(output)))
		for scanner.Scan() {
			files = append(files, filepath.Join(dir, scanner.Text()))
		}
	} else {
		err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if !info.IsDir() {
				files = append(files, path)
			}

			return nil
		})

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error walking through files: %v\n", err)
			os.Exit(1)
		}
	}

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", file, err)
			continue
		}

		formattedName := strings.TrimPrefix(file, dir)
		if strings.HasPrefix(formattedName, string(os.PathSeparator)) {
			formattedName = formattedName[1:]
		}

		archive.Files = append(archive.Files, txtar.File{
			Name: formattedName,
			Data: data,
		})
	}

	err := os.WriteFile(output, []byte(txtar.Format(archive)), 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing output file: %v\n", err)
		os.Exit(1)
	}
}
