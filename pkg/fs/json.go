package fs

import (
	"encoding/json"
	"io"
	"os"
)

type Json[T any] struct {
	Path string
}

func NewJson[T any](filePath string) *Json[T] {
	return &Json[T]{
		Path: filePath,
	}
}

func (f *Json[T]) Read() (T, error) {
	file, err := os.OpenFile(f.Path, os.O_RDONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		var zero T
		return zero, err
	}
	defer file.Close()

	var data T
	if err = json.NewDecoder(file).Decode(&data); err != nil && err != io.EOF {
		var zero T
		return zero, err
	}

	return data, nil
}

func (f *Json[T]) Write(data T) error {
	file, err := os.OpenFile(f.Path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(data)
}

func (f *Json[T]) Modify(update func(T) (T, error)) error {
	data, err := f.Read()
	if err != nil {
		return err
	}
	data, err = update(data)
	if err != nil {
		return err
	}
	if err := f.Write(data); err != nil {
		return err
	}
	return nil
}
