package inmemory

import (
	"reflect"
	"testing"
	"time"
)

func Test_cacheService_Get(t *testing.T) {

	sut := New(10).(*cacheService)

	sut.Set("abc", 123)

	type args struct {
		key string
	}
	tests := []struct {
		name      string
		c         *cacheService
		args      args
		wantValue interface{}
		wantFound bool
	}{
		{
			name: "test not found",
			c:    sut,
			args: args{
				key: "xyz",
			},
			wantValue: nil,
			wantFound: false,
		},
		{
			name: "test found",
			c:    sut,
			args: args{
				key: "abc",
			},
			wantValue: 123,
			wantFound: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotFound := tt.c.Get(tt.args.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("cacheService.Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotFound != tt.wantFound {
				t.Errorf("cacheService.Get() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func Test_cacheService_Get_timeout_elapsed(t *testing.T) {

	sut := New(1).(*cacheService)

	sut.Set("abc", 123)
	time.Sleep(2 * time.Second)

	type args struct {
		key string
	}
	tests := []struct {
		name      string
		c         *cacheService
		args      args
		wantValue interface{}
		wantFound bool
	}{
		{
			name: "cache item timeout",
			c:    sut,
			args: args{
				key: "abc",
			},
			wantValue: nil,
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotFound := tt.c.Get(tt.args.key)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("cacheService.Get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotFound != tt.wantFound {
				t.Errorf("cacheService.Get() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func Test_cacheService_Set(t *testing.T) {

	sut := New(10).(*cacheService)

	type args struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name string
		c    *cacheService
		args args
	}{
		{
			name: "set value",
			c:    sut,
			args: args{
				key:   "abc",
				value: 123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.c.Set(tt.args.key, tt.args.value)
		})
	}
}
