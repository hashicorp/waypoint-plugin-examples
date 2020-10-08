package utils

import (
	"io"
	"os"
	"path/filepath"
)

func CreateSimlink(srcFile, dstFile string) error {
	d, _ := filepath.Split(dstFile)

	// create the output directory
	os.MkdirAll(d, os.ModePerm)

	// if exists delete the destination file
	os.Remove(dstFile)

	err := os.Link(srcFile, dstFile)
	if err != nil {
		return err
	}

	return nil
}

func DeploymentCount(dstDir string) int {
	files, _ := filepath.Glob(filepath.Join(dstDir, "*.deployment"))
	return len(files)
}

func Filename(filePath string) string {
	_, f := filepath.Split(filePath)
	return f
}

func Directory(filePath string) string {
	d, _ := filepath.Split(filePath)
	return d
}

func CopyFile(srcFile, dstDir string) (string, error) {
	// get the filename from the filepath
	_, f := filepath.Split(srcFile)

	// create the output directory
	os.MkdirAll(dstDir, os.ModePerm)

	// open the file for reading
	src, err := os.Open(srcFile)
	if err != nil {
		return "", err
	}
	defer src.Close()

	// generate the output filename
	outputPath := filepath.Join(dstDir, f)

	// if exists delete the destination file
	os.Remove(outputPath)

	// create the destination file
	dst, err := os.Create(outputPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// copy the source file to the destination
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}

	// ensure the output has the same permissions
	perms, _ := src.Stat()
	os.Chmod(outputPath, perms.Mode())

	return outputPath, nil
}
