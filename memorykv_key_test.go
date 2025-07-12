// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memorykv

import (
	libkv "github.com/bborbe/kv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bucket Key Operations", func() {
	Describe("BucketToPrefix", func() {
		It("should add separator to bucket name", func() {
			bucket := libkv.BucketName("test-bucket")
			result := BucketToPrefix(bucket)
			expected := []byte("test-bucket_")
			Expect(result).To(Equal(expected))
		})

		It("should handle empty bucket name", func() {
			bucket := libkv.BucketName("")
			result := BucketToPrefix(bucket)
			expected := []byte{bucketKeySeperator}
			Expect(result).To(Equal(expected))
		})
	})

	Describe("BucketAddKey", func() {
		It("should combine bucket and key with separator", func() {
			bucket := libkv.BucketName("test-bucket")
			key := []byte("test-key")
			result := BucketAddKey(bucket, key)

			expected := []byte("test-bucket_test-key")
			Expect(result).To(Equal(expected))
		})

		It("should handle empty key", func() {
			bucket := libkv.BucketName("test-bucket")
			key := []byte{}
			result := BucketAddKey(bucket, key)

			expected := []byte("test-bucket_")
			Expect(result).To(Equal(expected))
		})

		It("should handle empty bucket and key", func() {
			bucket := libkv.BucketName("")
			key := []byte{}
			result := BucketAddKey(bucket, key)

			expected := []byte{bucketKeySeperator}
			Expect(result).To(Equal(expected))
		})
	})

	Describe("BucketRemoveKey", func() {
		It("should remove bucket prefix from key", func() {
			bucket := libkv.BucketName("test-bucket")
			originalKey := []byte("test-key")

			// First add the bucket prefix
			fullKey := BucketAddKey(bucket, originalKey)

			// Then remove it
			result := BucketRemoveKey(bucket, fullKey)
			Expect(result).To(Equal(originalKey))
		})

		It("should handle key with only bucket prefix", func() {
			bucket := libkv.BucketName("test-bucket")
			originalKey := []byte{}

			fullKey := BucketAddKey(bucket, originalKey)
			result := BucketRemoveKey(bucket, fullKey)
			Expect(result).To(Equal(originalKey))
		})

		It("should handle empty bucket", func() {
			bucket := libkv.BucketName("")
			originalKey := []byte("test-key")

			fullKey := BucketAddKey(bucket, originalKey)
			result := BucketRemoveKey(bucket, fullKey)
			Expect(result).To(Equal(originalKey))
		})
	})
})
