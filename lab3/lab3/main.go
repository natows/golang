package main

import (
	"time"
)

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleRemoveItem(vfs *VirtualFileSystem, name string) {
	err := vfs.RemoveItem(name)
	HandleError(err)
	_, err = vfs.GetItem(name)
	if err != nil {
		println("Item removed:", err.Error())
	} else {
		panic("Item should not be available")
	}
}
func main() {
	vfs := &VirtualFileSystem{root: &Katalog{name: "root", path: "/", items: make(map[string]FileSystemItem), createdAt: time.Now(), modifiedAt: time.Now()}}

	err := vfs.CreateFile("file1.txt", "/file1.txt", []byte("Hello World"))
	HandleError(err)
	content, err := vfs.ReadFile("file1.txt")
	HandleError(err)
	println("Odczytany plik:", string(content))

	err = vfs.WriteFile("file1.txt", []byte("Hello Go!")) 
	HandleError(err)
	content, err = vfs.ReadFile("file1.txt")
	HandleError(err)
	println("Odczytany plik po zapisie:", string(content))

	err = vfs.OverWriteFile("file1.txt", []byte("Hello Overwrite!"))
	HandleError(err)
	content, err = vfs.ReadFile("file1.txt")
	HandleError(err)
	println("Odczytany plik po nadpisaniu:", string(content))

	err = vfs.CreateDirectory("dir1", "/dir1")
	HandleError(err)
	err = vfs.CreateFile("file2.txt", "/dir1/file2.txt", []byte("Hello Directory"))
	HandleError(err)

	content, err = vfs.ReadFile("/dir1/file2.txt")
	HandleError(err)
	println("Odczytany plik z katalogu:", string(content))

	err = vfs.CreateSymlink("link_to_file1", "/link_to_file1", vfs.root.items["file1.txt"])
	HandleError(err)
	content, err = vfs.ReadFile("link_to_file1")
	HandleError(err)
	println("Odczytany plik z linku:", string(content))


	HandleRemoveItem(vfs, "file1.txt")


	HandleRemoveItem(vfs, "/dir1/file2.txt")
	

	HandleRemoveItem(vfs, "link_to_file1")


	HandleRemoveItem(vfs, "/dir1")

	// HandleRemoveItem(vfs, "/") //wyskoczy blad 
}

