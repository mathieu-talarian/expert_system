package filehandler

import (
	"bytes"
	"github.com/mathmoul/expert_system/src/esfmt"
	"os"
)

// File struct handle file opener, and getter / setters
type File struct {
	*os.File
}

// OpenedFile is a pointer on File struct, define an file Descriptor
var OpenedFile *File

// OpenFile to open file with fileName
func OpenFile(fileName string) (err error) {
	openFile, err := os.Open(fileName)
	OpenedFile = &File{
		File: openFile,
	}
	return err
}

func (f *File) set(file *os.File) {
	f.File = file
}
func (f File) get() *File {
	return OpenedFile
}

// ToString Convert file *os.File to string
func (f File) ToString() (str string) {
	buf := bytes.NewBuffer(nil)
	if _, err := buf.ReadFrom(f); err != nil {
		esfmt.LogAndExit(err.Error())
	}
	return buf.String()
}
