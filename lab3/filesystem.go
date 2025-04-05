package main

import (
	"strings"
	"time"
)

type VirtualFileSystem struct {
	root *Katalog
}

func splitPath(path string) (dirPath string, fileName string) {
	lastSlashIndex := strings.LastIndex(path, "/")

	if lastSlashIndex == -1 {
		return "/", path
	}

	if lastSlashIndex == 0 {
		return "/", path[1:]
	}

	return path[:lastSlashIndex], path[lastSlashIndex+1:]
}

func (vfs *VirtualFileSystem) findDirectory(path string) (*Katalog, error) {
	if path == "/" {
		return vfs.root, nil
	}

	if path[0] == '/' {
		path = path[1:]
	}

	components := strings.Split(path, "/")

	current := vfs.root

	for _, component := range components {
		if component == "" {
			continue
		}
		items := current.Items()
		found := false
		var item FileSystemItem
		for _, i := range items {
			if i.Name() == component {
				item = i
				found = true
				break
			}
		}
		if !found {
			return nil, ErrItemNotFound
		}

		dir, ok := item.(*Katalog)
		if !ok {
			return nil, ErrNotDirectory
		}

		current = dir
	}

	return current, nil
}

func (vfs *VirtualFileSystem) Root() *Katalog {
	return vfs.root
}

func (vfs *VirtualFileSystem) CreateFile(name string, path string, content []byte) error {
	dirPath, fileName := splitPath(path)

	dir, err := vfs.findDirectory(dirPath)
	if err != nil {
		return err
	}

	file := &Plik{
		name:       fileName,
		path:       path,
		content:    content,
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
	}

	return dir.AddItem(file)
}

func (vfs *VirtualFileSystem) CreateDirectory(name string, path string) error {
	dirPath, dirName := splitPath(path)

	parent, err := vfs.findDirectory(dirPath)
	if err != nil {
		return err
	}

	dir := &Katalog{
		name:       dirName,
		path:       path,
		items:      make(map[string]FileSystemItem),
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
	}

	return parent.AddItem(dir)
}

func (vfs *VirtualFileSystem) CreateSymlink(name string, path string, target FileSystemItem) error {
	dirPath, linkName := splitPath(path)

	dir, err := vfs.findDirectory(dirPath)
	if err != nil {
		return err
	}

	symlink := &SymLink{
		name:       linkName,
		path:       path,
		target:     target,
		createdAt:  time.Now(),
		modifiedAt: time.Now(),
	}

	return dir.AddItem(symlink)
}

func (vfs *VirtualFileSystem) RemoveItem(path string) error {
	if path == "/" {
		return ErrPermissionDenied
	}
	
	dirPath, itemName := splitPath(path)

	dir, err := vfs.findDirectory(dirPath)
	if err != nil {
		return err
	}

	return dir.RemoveItem(itemName)
}

func (vfs *VirtualFileSystem) GetItem(path string) (FileSystemItem, error) {
	dirPath, itemName := splitPath(path)

	dir, err := vfs.findDirectory(dirPath)
	if err != nil {
		return nil, err
	}

	for _, item := range dir.Items() {
        if item.Name() == itemName {
            return item, nil
        }
    }
	return nil, ErrItemNotFound
}

func (vfs *VirtualFileSystem) ListDirectory(path string) ([]string, error) {
    dir, err := vfs.findDirectory(path)
    if err != nil {
        return nil, err
    }
    
    items := dir.Items()
    names := make([]string, len(items))
    
    for i, item := range items {
        names[i] = item.Name()
    }
    
    return names, nil
}

func (vfs *VirtualFileSystem) ReadFile(path string) ([]byte, error) {
    item, err := vfs.GetItem(path)
    if err != nil {
        return nil, err
    }

    if symlink, ok := item.(*SymLink); ok {
        if symlink.target == nil {
            return nil, ErrItemNotFound
        }
        item = symlink.target
    }

    if file, ok := item.(*Plik); ok {
        buffer := make([]byte, file.Size())
        n, err := file.Read(buffer)
        if err != nil {
            return nil, err
        }
        return buffer[:n], nil
    }
    
    return nil, ErrNotImplemented
}

func (vfs *VirtualFileSystem) WriteFile(path string, content []byte) error {
	item, err := vfs.GetItem(path)
	if err != nil {
		return err
	}

	if file, ok := item.(*Plik); ok {
		_, err = file.Write(content)
		return err
	}

	return ErrNotImplemented
}

func (vfs *VirtualFileSystem) OverWriteFile(path string, content []byte) error {
	item, err := vfs.GetItem(path)
	if err != nil {
		return err
	}

	if file, ok := item.(*Plik); ok {
		_, err = file.OverWrite(content)
		return err
	}
	return ErrNotImplemented
}

func (vfs *VirtualFileSystem) SeeModifiedAt(path string) (time.Time, error) {
	item ,err := vfs.GetItem(path)
	if err != nil {
		return time.Time{}, err
	}
	return item.ModifiedAt(), nil
}

func (vfs *VirtualFileSystem) SeeCreatedAt(path string) (time.Time, error) {
	item ,err := vfs.GetItem(path)
	if err != nil {
		return time.Time{}, err
	}
	return item.CreatedAt(), nil
}

func (vfs *VirtualFileSystem) SeeSize(path string) (int64, error) {
	item ,err := vfs.GetItem(path)
	if err != nil {
		return 0, err
	}
	return item.Size(), nil
}
