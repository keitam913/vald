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

package session

type Option func(s *sess)

var (
	defaultOpts = []Option{
		WithEnableSSL(true),
	}
)

func WithEndpoint(ep string) Option {
	return func(s *sess) {
		s.endpoint = ep
	}
}

func WithRegion(rg string) Option {
	return func(s *sess) {
		s.region = rg
	}
}

func WithAccessKey(ak string) Option {
	return func(s *sess) {
		s.accessKey = ak
	}
}

func WithSecretAccessKey(sak string) Option {
	return func(s *sess) {
		s.secretAccessKey = sak
	}
}

func WithToken(tk string) Option {
	return func(s *sess) {
		s.token = tk
	}
}

func WithEnableSSL(enable bool) Option {
	return func(s *sess) {
		s.enableSSL = enable
	}
}
