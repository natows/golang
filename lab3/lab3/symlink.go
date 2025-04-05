package main

import (
	"time"

)

type SymLink struct {
	name string
	path string
	target FileSystemItem
	createdAt time.Time
	modifiedAt time.Time
}

func (s *SymLink) Name() string {
	return s.name
}

func (s *SymLink) Path() string{
	return s.path
}

func (s *SymLink) Size() int64{
	if s.target == nil {
		return 0
	}
	return s.target.Size()
}

func (s *SymLink) CreatedAt() time.Time {
	return s.createdAt
}

func (s *SymLink) ModifiedAt() time.Time {
	return s.modifiedAt
}