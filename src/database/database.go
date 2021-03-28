package database

// Local imports
import (
)

// External imports
import (
	bolt "go.etcd.io/bbolt"
)

type Database struct {
	connection *bolt.DB
	bucket []byte
}

func New(dbPath string) (db *Database, closeCallback func() error, err error) {
	connection, err := bolt.Open(dbPath, 0600, nil)

	if err != nil {
		return nil, nil, err
	}

	db = &Database{connection: connection, bucket: []byte("default")}
	closeCallback = connection.Close

	return
}

func (d *Database) SetBucket(bucket string) {
	d.bucket = []byte(bucket)
}

func (d *Database) Get(key string) (value []byte, err error) {
	err = d.connection.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(d.bucket)

		value = b.Get([]byte(key))

		return nil
	})

	return
}

func (d *Database) Set(key string, value []byte) error {
	return d.connection.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(d.bucket)

		return b.Put([]byte(key), value)
	})
}