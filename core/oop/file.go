package oop

type File struct {
	fd   int
	name string
}

// there's no concept of constructor in golang, as a conversion, we use NewXXX to
// create an instance of given type
func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{fd, name}
}
