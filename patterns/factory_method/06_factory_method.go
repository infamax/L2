package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const (
	distStorage = iota
	cache
	fileStorage
)

// Фабричный метод
// На примере некоторого хранилища данных: диск, кэш, база данных, файловый носитель и т.д

type Store interface {
	SaveData(data string) (int, error)
	ReadData(id int) (string, error)
}

type DiskStorage struct {
	storage []string
	limit   int
}

func (ds *DiskStorage) SaveData(data string) (int, error) {
	fmt.Println("limit = ", ds.limit)
	if len(ds.storage)+1 > ds.limit {
		return 0, errors.New("not enough space in disk")
	}
	ds.storage = append(ds.storage, data)
	log.Println("data successfully saved on disk")
	return len(ds.storage) - 1, nil
}

func (ds *DiskStorage) ReadData(id int) (string, error) {
	if id >= len(ds.storage) {
		return "", errors.New("not found object in this id in disk")
	}
	return ds.storage[id], nil
}

type Cache struct {
	cache  map[int]string
	lastID int
	limit  int
}

func (c *Cache) SaveData(data string) (int, error) {
	c.lastID++
	if c.lastID > c.limit {
		return 0, errors.New("not enough space in cache")
	}
	c.cache[c.lastID] = data
	return c.lastID, nil
}

func (c *Cache) ReadData(id int) (string, error) {
	if id > c.lastID {
		return "", errors.New("not found object in this id in cache")
	}
	return c.cache[id], nil
}

type FileStorage struct {
	limit int
	file  *os.File
}

func (f *FileStorage) SaveData(data string) (int, error) {
	return 0, nil
}

func (f *FileStorage) ReadData(id int) (string, error) {
	return "", nil
}

type Service struct{}

func (s *Service) CreateStorage(action int) Store {
	switch action {
	case distStorage:
		return &DiskStorage{
			storage: make([]string, 0),
			limit:   10,
		}
	case cache:
		return &Cache{
			cache: map[int]string{},
			limit: 20,
		}
	case fileStorage:
		return &FileStorage{}
	default:
		return nil
	}
}

func main() {
	srv := Service{}
	store := srv.CreateStorage(distStorage)
	id, err := store.SaveData("test")
	if err != nil {
		log.Fatal(err)
	}
	data, _ := store.ReadData(id)
	fmt.Println(data)
}
