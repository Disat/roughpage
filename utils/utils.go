package utils

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func RemoveDuplicates(arr *[]string) {
	templist := make([]string, 0, len(*arr))
	for _, rs := range *arr {
		for _, ls := range templist {
			if ls == rs {
				goto NextTurn
			}
		}
		templist = append(templist, rs)
	NextTurn:
	}
	*arr = templist
}

func ClearDir() {
	files, e := filepath.Glob("./public/*.html")
	if e != nil {
		fmt.Println(e)
	}
	for _, file := range files {
		e := os.Remove(file)
		if e != nil {
			fmt.Println(e)
		}
	}
}

func CreateDirectoryIfnotexist(filenames []string, parentdirectory string) {
	if _, err := os.Stat(parentdirectory); os.IsNotExist(err) {
		err := os.Mkdir(parentdirectory, 0755)
		if err != nil {
			fmt.Println(err)
		}
	}
	for _, file := range filenames {
		if _, err := os.Stat(parentdirectory + "/" + file); os.IsNotExist(err) {
			// Create the directory
			err := os.Mkdir(filepath.Join(parentdirectory, file), 0755)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}

func GetFileCreationAndModifiyTime(absolutefilepath string) (creationtime time.Time, modifytime time.Time) {
	fsinfo, err := os.Stat(absolutefilepath)
	if err != nil {
		fmt.Println(err)
	}

	// if runtime.GOOS == "windows" {
	// 	osStat, ok := fsinfo.Sys().(*syscall.Win32FileAttributeData)
	// 	if !ok {
	// 		panic("Unsupported type assertion")
	// 	}
	// 	creationTime := time.Unix(0, osStat.CreationTime.Nanoseconds())
	// 	return creationTime
	// } else if runtime.GOOS == "linux" {
	// 	stat, ok := fsinfo.Sys().(*syscall.Stat_t)
	// 	if !ok {
	// 		panic("Unsupported type assertion")
	// 	}
	// 	creationTime := time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	// 	return creationTime
	// } else {
	// 	panic("Doesn't support this OS")
	// }

	osStat, ok := fsinfo.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		panic("Unsupported type assertion")
	}
	creationtime = time.Unix(0, osStat.CreationTime.Nanoseconds())
	modifytime = fsinfo.ModTime()
	return creationtime, modifytime
}

func Copy(srcList []string) {
	rootPath := "./themes"
	publicPath := "./public"
	for _, source := range srcList {
		if matches, err := filepath.Glob(filepath.Join(rootPath, source)); err == nil {
			for _, srcPath := range matches {
				fmt.Println("Copying " + srcPath)
				file, err := os.Stat(srcPath)
				if err != nil {
					fmt.Println(err)
				}
				fileName := file.Name()
				desPath := filepath.Join(publicPath, fileName)
				if file.IsDir() {
					CopyDir(srcPath, desPath)
				} else {
					CopyFile(srcPath, desPath)
				}
			}
		} else {
			fmt.Println(err.Error())
		}
	}
}

func CopyFile(source string, dest string) {
	sourcefile, err := os.Open(source)
	if err != nil {
		fmt.Println(err.Error())
	}
	destfile, err := os.Create(dest)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer destfile.Close()
	_, err = io.Copy(destfile, sourcefile)
	if err != nil {
		fmt.Println(err.Error())
	}
	sourceinfo, err := os.Stat(source)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = os.Chmod(dest, sourceinfo.Mode())
	if err != nil {
		fmt.Println(err.Error())
	}
	sourcefile.Close()
}

func CopyDir(source string, dest string) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		fmt.Println(err.Error())
	}
	directory, _ := os.Open(source)
	defer directory.Close()
	objects, err := directory.Readdir(-1)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, obj := range objects {
		sourcefilepointer := source + "/" + obj.Name()
		destinationfilepointer := dest + "/" + obj.Name()
		if obj.IsDir() {
			CopyDir(sourcefilepointer, destinationfilepointer)
		} else {
			CopyFile(sourcefilepointer, destinationfilepointer)
		}
	}
}
