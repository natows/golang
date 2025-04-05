package main

import (
	"time"
)

type Katalog struct {
	name string
	path string
	items map[string]FileSystemItem 
	createdAt time.Time
	modifiedAt time.Time
}

func (k *Katalog) Name() string {
	return k.name
}
func (k *Katalog) Path() string {
	return k.path
}
func (k *Katalog) Size() int64 {
	return 0
}
func (k *Katalog) CreatedAt() time.Time {
	return k.createdAt
}
func (k *Katalog) ModifiedAt() time.Time {
	return k.modifiedAt
}
func (k *Katalog) AddItem(item FileSystemItem) error {
	if k.items[item.Name()] != nil {
		return ErrItemExists
	}
	k.items[item.Name()] = item
	k.modifiedAt = time.Now()
	return nil

} 

func (k *Katalog) RemoveItem(name string) error { 
    item := k.items[name]
    if item == nil {
        return ErrItemNotFound
    }
    
    if dir, ok := item.(*Katalog); ok {
        for _, subItem := range dir.Items() {
            err := dir.RemoveItem(subItem.Name())
            if err != nil {
                return err
            }
        }
    }
    
    delete(k.items, name)
    k.modifiedAt = time.Now()
    return nil
}

func (k *Katalog) Items() []FileSystemItem {
	items := make([]FileSystemItem, 0 , len(k.items))
	for _, item := range k.items {
        items = append(items, item)
    }
    return items
}