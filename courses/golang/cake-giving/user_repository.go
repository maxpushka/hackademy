package main

import (
	"crypto/md5"
	"errors"
	"os"
	"sync"
)

type InMemoryUserStorage struct {
	lock    sync.RWMutex
	storage map[string]User
}

func NewInMemoryUserStorage() *InMemoryUserStorage {

	passwordDigest := md5.New().Sum([]byte(os.Getenv("CAKE_ADMIN_PASSWORD")))
	sa := User{
		Email:          os.Getenv("CAKE_ADMIN_EMAIL"),
		PasswordDigest: string(passwordDigest),
		Role:           "superadmin",
		FavoriteCake:   "BiscuitCake",
		IsBan:          false,
	}

	newStorage := make(map[string]User)
	newStorage[os.Getenv("CAKE_ADMIN_EMAIL")] = sa

	numberOfRegistered.Inc()
	return &InMemoryUserStorage{
		lock:    sync.RWMutex{},
		storage: newStorage,
	}
}

func (userStorage InMemoryUserStorage) Add(key string, u User) error {
	_, ok := userStorage.storage[key]
	if ok || u.Email == os.Getenv("CAKE_ADMIN_EMAIL") {
		err := errors.New("this user is already exists")
		return err
	}
	userStorage.storage[key] = u
	return nil
}

func (userStorage InMemoryUserStorage) Get(key string) (User, error) {
	if u, ok := userStorage.storage[key]; ok {
		return u, nil
	}
	err := errors.New("there is no such user")
	empty := User{}
	return empty, err
}

func (userStorage InMemoryUserStorage) Update(key string, u User) error {
	userStorage.storage[key] = u
	return nil
}

func (userStorage InMemoryUserStorage) Delete(key string) (User, error) {
	if u, ok := userStorage.storage[key]; ok {
		delete(userStorage.storage, key)
		return u, nil
	}
	err := errors.New("there is no such user")
	empty := User{}
	return empty, err
}
