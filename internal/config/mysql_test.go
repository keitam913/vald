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

// Package config providers configuration type and load configuration logic
package config

import (
	"reflect"
	"testing"

	"github.com/vdaas/vald/internal/errors"
)

func TestMySQL_Bind(t *testing.T) {
	type fields struct {
		DB                   string
		Host                 string
		Port                 int
		User                 string
		Pass                 string
		Name                 string
		Charset              string
		Timezone             string
		InitialPingTimeLimit string
		InitialPingDuration  string
		ConnMaxLifeTime      string
		MaxOpenConns         int
		MaxIdleConns         int
		TLS                  *TLS
		TCP                  *TCP
	}
	type want struct {
		want *MySQL
	}
	type test struct {
		name       string
		fields     fields
		want       want
		checkFunc  func(want, *MySQL) error
		beforeFunc func()
		afterFunc  func()
	}
	defaultCheckFunc := func(w want, got *MySQL) error {
		if !reflect.DeepEqual(got, w.want) {
			return errors.Errorf("got = %v, want %v", got, w.want)
		}
		return nil
	}
	tests := []test{
		// TODO test cases
		/*
		   {
		       name: "test_case_1",
		       fields: fields {
		           DB: "",
		           Host: "",
		           Port: 0,
		           User: "",
		           Pass: "",
		           Name: "",
		           Charset: "",
		           Timezone: "",
		           InitialPingTimeLimit: "",
		           InitialPingDuration: "",
		           ConnMaxLifeTime: "",
		           MaxOpenConns: 0,
		           MaxIdleConns: 0,
		           TLS: TLS{},
		           TCP: TCP{},
		       },
		       want: want{},
		       checkFunc: defaultCheckFunc,
		   },
		*/

		// TODO test cases
		/*
		   func() test {
		       return test {
		           name: "test_case_2",
		           fields: fields {
		           DB: "",
		           Host: "",
		           Port: 0,
		           User: "",
		           Pass: "",
		           Name: "",
		           Charset: "",
		           Timezone: "",
		           InitialPingTimeLimit: "",
		           InitialPingDuration: "",
		           ConnMaxLifeTime: "",
		           MaxOpenConns: 0,
		           MaxIdleConns: 0,
		           TLS: TLS{},
		           TCP: TCP{},
		           },
		           want: want{},
		           checkFunc: defaultCheckFunc,
		       }
		   }(),
		*/
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if test.beforeFunc != nil {
				test.beforeFunc()
			}
			if test.afterFunc != nil {
				defer test.afterFunc()
			}
			if test.checkFunc == nil {
				test.checkFunc = defaultCheckFunc
			}
			m := &MySQL{
				DB:                   test.fields.DB,
				Host:                 test.fields.Host,
				Port:                 test.fields.Port,
				User:                 test.fields.User,
				Pass:                 test.fields.Pass,
				Name:                 test.fields.Name,
				Charset:              test.fields.Charset,
				Timezone:             test.fields.Timezone,
				InitialPingTimeLimit: test.fields.InitialPingTimeLimit,
				InitialPingDuration:  test.fields.InitialPingDuration,
				ConnMaxLifeTime:      test.fields.ConnMaxLifeTime,
				MaxOpenConns:         test.fields.MaxOpenConns,
				MaxIdleConns:         test.fields.MaxIdleConns,
				TLS:                  test.fields.TLS,
				TCP:                  test.fields.TCP,
			}

			got := m.Bind()
			if err := test.checkFunc(test.want, got); err != nil {
				tt.Errorf("error = %v", err)
			}

		})
	}
}
