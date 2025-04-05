package main

import (
	"time"
)

type Plik struct {
	name string
	path string
	content []byte
	createdAt time.Time
	modifiedAt time.Time
}

func (p *Plik) Name() string {
	return p.name
}
func (p *Plik) Path() string {
	return p.path
}
func (p *Plik) Size() int64 {
	return int64(len(p.content))
}
func (p *Plik) CreatedAt() time.Time {
	return p.createdAt
}
func (p *Plik) ModifiedAt() time.Time {
	return p.modifiedAt
}

func (p *Plik) Read(c []byte) (n int, err error) {
	if len(p.content) == 0 {
		return 0, nil
	}
	n = copy(c, p.content)
	return n, nil
}

func (p *Plik) Write(c []byte) (n int, err error){
	p.content = append(p.content, c...)
	p.modifiedAt = time.Now()
	return len(c), nil
}

func (p *Plik) OverWrite(c []byte) (n int, err error){
	p.content = c
	p.modifiedAt = time.Now()
	return len(c), nil
}




type PlikDoOdczytu struct {
	Plik
}


func (p *PlikDoOdczytu) Write(c []byte) (n int, err error) {
    return 0, ErrPermissionDenied
}
