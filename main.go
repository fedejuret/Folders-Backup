package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

func main() {
	myFigure := figure.NewColorFigure("Folder Backup", "", "cyan", true)
	myFigure.Print()
	fmt.Println(color.YellowString(" By Federico Juretich <fedejuret@gmail.com>"))
	fmt.Println()
	fmt.Println()
	fmt.Println()

	origin := flag.String("origin", "", "Source directory for backup")
	destination := flag.String("destination", "", "Destination directory for backup")
	months := flag.Int("months", 6, "How many months ago did we create the backup?")
	deleteOriginals := flag.Bool("delete-originals", false, "Whether or not you want to delete files from the source after they are copied to the destination")
	flag.Parse()

	if *origin == "" || *destination == "" {
		fmt.Println("Error: Source and destination directories must be provided.")
		flag.PrintDefaults()
		os.Exit(1)
	}

	backupOrigin, err := filepath.Abs(*origin)
	if err != nil {
		fmt.Println("Error getting absolute path of source directory:", err)
		os.Exit(1)
	}

	backupDestination, err := filepath.Abs(*destination)
	if err != nil {
		fmt.Println("Error getting absolute path of destination directory:", err)
		os.Exit(1)
	}

	monthsAgo := time.Now().AddDate(0, -*months, 0)

	err = BackupFolder(backupOrigin, backupDestination, monthsAgo, *deleteOriginals)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func BackupFolder(origin, destination string, monthsAgo time.Time, deleteOriginals bool) error {

	if err := os.MkdirAll(destination, os.ModePerm); err != nil {
		return err
	}

	files, err := os.ReadDir(origin)
	if err != nil {
		return err
	}

	for _, file := range files {
		originPath := filepath.Join(origin, file.Name())
		backupPath := filepath.Join(destination, file.Name())

		if file.IsDir() {
			err := BackupFolder(originPath, backupPath, monthsAgo, deleteOriginals)
			if err != nil {
				return err
			}
		} else {
			info, err := file.Info()
			if err != nil {
				return err
			}

			if info.ModTime().Before(monthsAgo) {
				err := CopyFile(originPath, backupPath, deleteOriginals)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func CopyFile(origin, destination string, deleteOriginals bool) error {

	fileOrigin, err := os.Open(origin)
	if err != nil {
		return err
	}
	defer fileOrigin.Close()

	fileDestination, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer fileDestination.Close()

	_, err = io.Copy(fileDestination, fileOrigin)

	if err := fileOrigin.Close(); err != nil {
		return err
	}

	if err := fileDestination.Close(); err != nil {
		return err
	}

	if deleteOriginals {
		err = os.Remove(origin)
	}
	return err
}
