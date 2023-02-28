package auto_ru

import (
	"reflect"
	"testing"

	"github.com/tonatos/goal-tracker/internal/services"
	"github.com/tonatos/goal-tracker/internal/services/auto_ru"
)

func TestAutoRU_GetCatalogLink(t *testing.T) {
	type fields struct {
		Amount       float32
		Urls         auto_ru.AutoRUApiURLs
		Api          services.Api
		FilterParams auto_ru.AutoRUCountAdsRequest
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := &auto_ru.AutoRU{
				Amount:       tt.fields.Amount,
				Urls:         tt.fields.Urls,
				Api:          tt.fields.Api,
				FilterParams: tt.fields.FilterParams,
			}
			got, err := ar.GetCatalogLink()
			if (err != nil) != tt.wantErr {
				t.Errorf("AutoRU.GetCatalogLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AutoRU.GetCatalogLink() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAutoRU_CountAds(t *testing.T) {
	type fields struct {
		Amount       float32
		Urls         auto_ru.AutoRUApiURLs
		Api          services.Api
		FilterParams auto_ru.AutoRUCountAdsRequest
	}
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ar := &auto_ru.AutoRU{
				Amount:       tt.fields.Amount,
				Urls:         tt.fields.Urls,
				Api:          tt.fields.Api,
				FilterParams: tt.fields.FilterParams,
			}
			got, err := ar.CountAds()
			if (err != nil) != tt.wantErr {
				t.Errorf("AutoRU.CountAds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AutoRU.CountAds() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAutoruInit(t *testing.T) {
	type args struct {
		accumulated_sum float32
		goal_amount     float32
	}
	tests := []struct {
		name string
		args args
		want *auto_ru.AutoRU
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := auto_ru.AutoruInit(tt.args.goal_amount, tt.args.accumulated_sum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AutoruInit() = %v, want %v", got, tt.want)
			}
		})
	}
}
