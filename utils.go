package gradexpath

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/timdrysdale/copy"
)

//need to be case insensitive
func IsPdf(path string) bool {
	suffix := strings.ToLower(filepath.Ext(path))
	return strings.Compare(suffix, ".pdf") == 0
}

func IsTxt(path string) bool {
	suffix := strings.ToLower(filepath.Ext(path))
	return strings.Compare(suffix, ".txt") == 0
}

func IsZip(path string) bool {
	suffix := strings.ToLower(filepath.Ext(path))
	return strings.Compare(suffix, ".zip") == 0
}

func Copy(source, destination string) error {
	// last param is buffer size ...
	if IsTxt(source) {
		return copy.Copy(source, destination, 1024)
	} else {
		return copy.Copy(source, destination, 130048)
	}

}

func EnsureDir(dirName string) error {

	err := os.Mkdir(dirName, 0755) //probably umasked with 22 not 02

	os.Chmod(dirName, 0755)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

func EnsureDirAll(dirName string) error {

	err := os.MkdirAll(dirName, 0755) //probably umasked with 22 not 02

	os.Chmod(dirName, 0755)

	if err == nil || os.IsExist(err) {
		return nil
	} else {
		return err
	}
}

func GetFileList(dir string) ([]string, error) {

	paths := []string{}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			paths = append(paths, path)
		}

		return nil
	})

	return paths, err

}

func CopyIsComplete(source, dest []string) bool {

	sourceBase := BaseList(source)
	destBase := BaseList(dest)

	for _, item := range sourceBase {

		if !itemExists(destBase, item) {
			return false
		}
	}

	return true

}

func BaseList(paths []string) []string {

	bases := []string{}

	for _, path := range paths {
		bases = append(bases, filepath.Base(path))
	}

	return bases
}

// Mod from array to slice,
// from https://www.golangprograms.com/golang-check-if-array-element-exists.html
func itemExists(sliceType interface{}, item interface{}) bool {
	slice := reflect.ValueOf(sliceType)

	if slice.Kind() != reflect.Slice {
		panic("Invalid data-type")
	}

	for i := 0; i < slice.Len(); i++ {
		if slice.Index(i).Interface() == item {
			return true
		}
	}

	return false
}