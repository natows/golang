package main


import (
	"time"
)

type FileSystemItem interface {
	Name() string
	Path() string
	Size() int64
	CreatedAt() time.Time
	ModifiedAt() time.Time
}

type Readable interface {
	Read(p []byte) (n int, err error)
}

type Writable interface {
	Write(p []byte) (n int, err error)
}

type Directory interface {
	FileSystemItem
	AddItem(item FileSystemItem) error
	RemoveItem(name string) error
	Items() []FileSystemItem
}
