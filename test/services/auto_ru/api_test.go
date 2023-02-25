package auto_ru

import (
	"io"
	"net/url"
	"reflect"
	"testing"

	"github.com/tonatos/goal-tracker/pkg/services/auto_ru"
)

func TestAutoRUApi_request(t *testing.T) {
	type fields struct {
		BaseURL string
		Headers map[string]string
	}
	type args struct {
		method string
		url    *url.URL
		data   io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &auto_ru.AutoRUApi{
				BaseURL: tt.fields.BaseURL,
				Headers: tt.fields.Headers,
			}
			got, err := api.Request(tt.args.method, tt.args.url, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("AutoRUApi.request() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AutoRUApi.request() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAutoRUApi_BuildURL(t *testing.T) {
	type fields struct {
		BaseURL string
		Headers map[string]string
	}
	type args struct {
		urlName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *url.URL
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &auto_ru.AutoRUApi{
				BaseURL: tt.fields.BaseURL,
				Headers: tt.fields.Headers,
			}
			if got := api.BuildURL(tt.args.urlName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AutoRUApi.BuildURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAutoRUApi_Post(t *testing.T) {
	type fields struct {
		BaseURL string
		Headers map[string]string
	}
	type args struct {
		url  *url.URL
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &auto_ru.AutoRUApi{
				BaseURL: tt.fields.BaseURL,
				Headers: tt.fields.Headers,
			}
			got, err := api.Post(tt.args.url, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("AutoRUApi.Post() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AutoRUApi.Post() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAutoRUApi_Get(t *testing.T) {
	type fields struct {
		BaseURL string
		Headers map[string]string
	}
	type args struct {
		url *url.URL
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &auto_ru.AutoRUApi{
				BaseURL: tt.fields.BaseURL,
				Headers: tt.fields.Headers,
			}
			got, err := api.Get(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("AutoRUApi.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AutoRUApi.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
