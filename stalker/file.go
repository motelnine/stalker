package stalker

import (
	"io/ioutil"
	"fmt"
	"os"
	"path/filepath"
)

// Filesize returns size of file as int64
func FileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
	}
	return info.Size()
}

//DirSize returns size of directory as int64
func DirSize(path string) int64 {
    var size int64
    filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
        if err != nil {
            fmt.Println(err)
        }
        if !info.IsDir() {
            size += info.Size()
        }
       return err
    })
    return size
}

// ReadFile reads file data and returns as string
func ReadFile(filename string) string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return string(content)
}
