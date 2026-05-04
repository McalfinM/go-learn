package storage

import "sync"

var (
	instance Storage
	once     sync.Once
)

func GetStorage() Storage {
	once.Do(func() {
		// 👉 nanti bisa switch provider di sini
		s, err := NewCloudinaryStorage()
		if err != nil {
			panic(err)
		}
		instance = s
	})

	return instance
}
