package database

import (
	"fmt"
	"log"

	"go.etcd.io/bbolt"
)

type BoltDB struct {
	DB *bbolt.DB
}

// 打开数据库
func (db *BoltDB) Open(filepath string) error {
	var err error
	db.DB, err = bbolt.Open(filepath, 0600, nil)
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}
	return nil
}

// 关闭数据库
func (db *BoltDB) Close() {
	if err := db.DB.Close(); err != nil {
		log.Printf("failed to close the database: %v", err)
	}
}

// 执行查询操作，使用回调函数进行事务处理
func (db *BoltDB) View(fn func(tx *bbolt.Tx) error) error {
	return db.DB.View(fn)
}

// 执行修改操作，使用回调函数进行事务处理
func (db *BoltDB) Update(fn func(tx *bbolt.Tx) error) error {
	return db.DB.Update(fn)
}

// 创建一个桶
func (db *BoltDB) CreateBucketIfNotExists(bucketName string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		return err
	})
}

// 获取一个桶
func (db *BoltDB) GetBucket(bucketName string) (*bbolt.Bucket, error) {
	var bucket *bbolt.Bucket
	err := db.View(func(tx *bbolt.Tx) error {
		bucket = tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}
		return nil
	})
	return bucket, err
}

// 插入数据
func (db *BoltDB) Put(bucketName string, key, value []byte) error {
	return db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}
		return bucket.Put(key, value)
	})
}

// 获取数据
func (db *BoltDB) Get(bucketName string, key []byte) ([]byte, error) {
	var value []byte
	err := db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}
		value = bucket.Get(key)
		if value == nil {
			return fmt.Errorf("key %s not found", key)
		}
		return nil
	})
	return value, err
}

// 删除数据
func (db *BoltDB) Delete(bucketName string, key []byte) error {
	return db.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}
		return bucket.Delete(key)
	})
}

// 获取指定桶的所有键值对
func (db *BoltDB) GetAllKeysAndValues(bucketName string) (map[string]string, error) {
	result := make(map[string]string)

	// 以读取模式打开数据库
	err := db.View(func(tx *bbolt.Tx) error {
		// 获取指定的桶
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return fmt.Errorf("bucket %s not found", bucketName)
		}

		bucket.ForEach(func(k, v []byte) error {
			result[string(k)] = string(v)
			return nil
		})

		return nil
	})

	return result, err
}
