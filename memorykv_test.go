// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memorykv_test

import (
	"context"

	libkv "github.com/bborbe/kv"
	"github.com/bborbe/memorykv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("BadgerKV", func() {
	var ctx context.Context
	var db libkv.DB
	var err error
	var provider libkv.ProviderFunc = func(ctx context.Context) (libkv.DB, error) {
		return db, nil
	}
	BeforeEach(func() {
		ctx = context.Background()
		db, err = memorykv.OpenMemory(ctx)
		Expect(err).To(BeNil())
	})
	AfterEach(func() {
		_ = db.Close()
	})
	libkv.BucketTestSuite(provider)
	libkv.BasicTestSuite(provider)
	libkv.IteratorTestSuite(provider)
})
