// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memorykv_test

import (
	"context"

	libkv "github.com/bborbe/kv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/bborbe/memorykv"
)

var _ = Describe("MemoryKV Unit Tests", func() {
	var ctx context.Context
	var db libkv.DB

	BeforeEach(func() {
		ctx = context.Background()
		var err error
		db, err = memorykv.OpenMemory(ctx)
		Expect(err).To(BeNil())
	})

	AfterEach(func() {
		_ = db.Close()
	})

	Describe("DB operations", func() {
		It("should handle Remove operation", func() {
			err := db.Remove()
			Expect(err).To(BeNil())
		})

		It("should handle Sync operation", func() {
			err := db.Sync()
			Expect(err).To(BeNil())
		})
	})

	Describe("Transaction operations", func() {
		It("should list bucket names", func() {
			bucketName := libkv.BucketName("test-bucket")
			err := db.Update(ctx, func(ctx context.Context, tx libkv.Tx) error {
				_, err := tx.CreateBucket(ctx, bucketName)
				return err
			})
			Expect(err).To(BeNil())

			err = db.View(ctx, func(ctx context.Context, tx libkv.Tx) error {
				bucketNames, err := tx.ListBucketNames(ctx)
				Expect(err).To(BeNil())
				Expect(bucketNames).To(ContainElement(bucketName))
				return nil
			})
			Expect(err).To(BeNil())
		})

		It("should list empty bucket names when no buckets exist", func() {
			err := db.View(ctx, func(ctx context.Context, tx libkv.Tx) error {
				bucketNames, err := tx.ListBucketNames(ctx)
				Expect(err).To(BeNil())
				Expect(bucketNames).To(BeEmpty())
				return nil
			})
			Expect(err).To(BeNil())
		})
	})

	Describe("Item operations", func() {
		It("should return true for Exists", func() {
			bucketName := libkv.BucketName("test-bucket")
			key := []byte("test-key")
			value := []byte("test-value")

			err := db.Update(ctx, func(ctx context.Context, tx libkv.Tx) error {
				bucket, err := tx.CreateBucket(ctx, bucketName)
				if err != nil {
					return err
				}
				return bucket.Put(ctx, key, value)
			})
			Expect(err).To(BeNil())

			err = db.View(ctx, func(ctx context.Context, tx libkv.Tx) error {
				bucket, err := tx.Bucket(ctx, bucketName)
				if err != nil {
					return err
				}
				item, err := bucket.Get(ctx, key)
				Expect(err).To(BeNil())
				Expect(item.Exists()).To(BeTrue())
				return nil
			})
			Expect(err).To(BeNil())
		})

		It("should handle Value method with callback", func() {
			bucketName := libkv.BucketName("test-bucket")
			key := []byte("test-key")
			value := []byte("test-value")

			err := db.Update(ctx, func(ctx context.Context, tx libkv.Tx) error {
				bucket, err := tx.CreateBucket(ctx, bucketName)
				if err != nil {
					return err
				}
				return bucket.Put(ctx, key, value)
			})
			Expect(err).To(BeNil())

			err = db.View(ctx, func(ctx context.Context, tx libkv.Tx) error {
				bucket, err := tx.Bucket(ctx, bucketName)
				if err != nil {
					return err
				}
				item, err := bucket.Get(ctx, key)
				Expect(err).To(BeNil())

				var retrievedValue []byte
				err = item.Value(func(val []byte) error {
					retrievedValue = make([]byte, len(val))
					copy(retrievedValue, val)
					return nil
				})
				Expect(err).To(BeNil())
				Expect(retrievedValue).To(Equal(value))
				return nil
			})
			Expect(err).To(BeNil())
		})
	})

	Describe("Iterator operations", func() {
		It("should handle Close operation for forward iterator", func() {
			bucketName := libkv.BucketName("test-bucket")

			err := db.Update(ctx, func(ctx context.Context, tx libkv.Tx) error {
				_, err := tx.CreateBucket(ctx, bucketName)
				return err
			})
			Expect(err).To(BeNil())

			err = db.View(ctx, func(ctx context.Context, tx libkv.Tx) error {
				bucket, err := tx.Bucket(ctx, bucketName)
				if err != nil {
					return err
				}
				iter := bucket.Iterator()
				iter.Close() // Should not panic
				return nil
			})
			Expect(err).To(BeNil())
		})

		It("should handle Close operation for reverse iterator", func() {
			bucketName := libkv.BucketName("test-bucket")

			err := db.Update(ctx, func(ctx context.Context, tx libkv.Tx) error {
				_, err := tx.CreateBucket(ctx, bucketName)
				return err
			})
			Expect(err).To(BeNil())

			err = db.View(ctx, func(ctx context.Context, tx libkv.Tx) error {
				bucket, err := tx.Bucket(ctx, bucketName)
				if err != nil {
					return err
				}
				iter := bucket.IteratorReverse()
				iter.Close() // Should not panic
				return nil
			})
			Expect(err).To(BeNil())
		})
	})
})
