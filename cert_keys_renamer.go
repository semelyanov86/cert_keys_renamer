package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
)

const DefaultPath = "/etc/letsencrypt/archive/keys.sergeyem.ru"

const DefaultTarget = "/opt/bitwarden/bwdata/ssl/keys.sergeyem.ru"

func main() {
	fileMapping := make(map[string]string)
	fileMapping["chain.pem"] = "ca.crt"
	//fileMapping["cert.pem"] = "certificate.crt"
	fileMapping["privkey.pem"] = "private.key"
	fileMapping["fullchain.pem"] = "certificate.crt"
	// Define command line flags
	pathToFind := flag.String("path", DefaultPath, "path where we will find certificates")
	targetPath := flag.String("target", DefaultTarget, "new name of the file")
	flag.Parse()

	// Check if required flags are present
	if *pathToFind == "" || *targetPath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	for target, result := range fileMapping {
		err := FindAndCopyFile(target, result, pathToFind, targetPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%s copied to %s\n", target, result)
	}
}

func FindAndCopyFile(findPtr string, newPtr string, pathToFind *string, targetPath *string) error {
	// Get a list of files in certs directory
	files, err := os.ReadDir(*pathToFind)
	if err != nil {
		return err
	}
	// Filter files matching the nameToFind pattern
	var matchingFiles []os.DirEntry

	pattern := regexp.MustCompile(fmt.Sprintf("^%s\\d*\\.pem$", regexp.QuoteMeta(fileNameWithoutExtSliceNotation(findPtr))))
	for _, file := range files {
		if pattern.MatchString(file.Name()) {
			matchingFiles = append(matchingFiles, file)
		}
	}

	// Sort matching files by name in descending order
	sort.Slice(matchingFiles, func(i, j int) bool {
		return matchingFiles[i].Name() > matchingFiles[j].Name()
	})

	// Check if any matching files were found
	if len(matchingFiles) == 0 {
		return errors.New("No files matching " + findPtr + " found in " + *targetPath)
	}

	// Get the latest matching file
	latestFile := matchingFiles[0]

	// Copy the latest matching file to /opt/certs with the new name
	src := filepath.Join(*pathToFind, latestFile.Name())
	dst := filepath.Join(*targetPath, newPtr)
	err = copyFile(src, dst)
	return err
}

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
