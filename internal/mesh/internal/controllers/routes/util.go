// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package routes

import (
	"google.golang.org/protobuf/proto"
)

func protoClone[T proto.Message](v T) T {
	return proto.Clone(v).(T)
}

func protoSliceClone[T proto.Message](in []T) []T {
	if in == nil {
		return nil
	}
	out := make([]T, 0, len(in))
	for _, v := range in {
		out = append(out, protoClone[T](v))
	}
	return out
}
