//
// Copyright (C) 2019-2020 Vdaas.org Vald team ( kpango, rinx, kmrmt )
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Package observer provides storage observer
package observer

import (
	"github.com/vdaas/vald/internal/errgroup"
	"github.com/vdaas/vald/internal/timeutil"
	"github.com/vdaas/vald/pkg/agent/sidecar/service/storage"
)

type Option func(o *observer) error

var (
	defaultOpts = []Option{
		WithErrGroup(errgroup.Get()),
		WithBackupDuration("10m"),
		WithPostStopTimeout("10s"),
	}
)

func WithBackupDuration(dur string) Option {
	return func(o *observer) error {
		if dur == "" {
			return nil
		}
		d, err := timeutil.Parse(dur)
		if err != nil {
			return nil
		}
		o.checkDuration = d
		return nil
	}
}

func WithPostStopTimeout(dur string) Option {
	return func(o *observer) error {
		if dur == "" {
			return nil
		}
		d, err := timeutil.Parse(dur)
		if err != nil {
			return nil
		}
		o.postStopTimeout = d
		return nil
	}
}

func WithErrGroup(eg errgroup.Group) Option {
	return func(o *observer) error {
		if eg != nil {
			o.eg = eg
		}
		return nil
	}
}

func WithDir(dir string) Option {
	return func(o *observer) error {
		if dir == "" {
			return nil
		}

		o.dir = dir

		return nil
	}
}

func WithBlobStorage(storage storage.Storage) Option {
	return func(o *observer) error {
		if storage != nil {
			o.storage = storage
		}
		return nil
	}
}
