// Copyright (c) 2026 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package memorykv

import (
	"context"

	"github.com/bborbe/errors"
	libkv "github.com/bborbe/kv"
)

// Stats returns per-bucket key counts for the in-memory database.
// Bucket sizes in bytes are not reported (left at zero) because computing
// them would require walking every key/value pair.
func (b *memorydb) Stats(ctx context.Context) (libkv.Stats, error) {
	s := libkv.Stats{Backend: "memory"}
	err := b.View(ctx, func(ctx context.Context, tx libkv.Tx) error {
		names, err := tx.ListBucketNames(ctx)
		if err != nil {
			return errors.Wrapf(ctx, err, "list bucket names failed")
		}
		for _, name := range names {
			bucket, err := tx.Bucket(ctx, name)
			if err != nil {
				return errors.Wrapf(ctx, err, "get bucket %s failed", name)
			}
			count, err := libkv.Count(ctx, bucket)
			if err != nil {
				return errors.Wrapf(ctx, err, "count bucket %s failed", name)
			}
			s.Buckets = append(s.Buckets, libkv.BucketStats{
				Name:     name,
				KeyCount: count,
			})
		}
		return nil
	})
	if err != nil {
		return libkv.Stats{}, errors.Wrapf(ctx, err, "stats failed")
	}
	return s, nil
}
