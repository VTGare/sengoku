package sengoku

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestNewSengoku(t *testing.T) {
	type args struct {
		config []Config
	}
	tests := []struct {
		name string
		args args
		want *Sengoku
	}{
		{
			name: "With config",
			args: args{
				[]Config{{"apikey", 5, false, 5}},
			},
			want: &Sengoku{&Config{"apikey", 5, false, 5}, &http.Client{}, "https://saucenao.com"},
		},
		{
			name: "Without config",
			args: args{nil},
			want: &Sengoku{DefaultConfig(), &http.Client{}, "https://saucenao.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSengoku(tt.args.config...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSengoku() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfigQuery(t *testing.T) {
	type args struct {
		uri      *url.URL
		imageURL string
	}
	tests := []struct {
		name string
		c    *Config
		args args
		want string
	}{
		{"Test mode enabled", &Config{"1", 123, true, 5}, args{&url.URL{}, "https://imgur.com/image.png"}, "api_key=1&db=123&num_res=5&output_type=2&test_mode=1&url=https%3A%2F%2Fimgur.com%2Fimage.png"},
		{"Without API key", &Config{"", 123, false, 5}, args{&url.URL{}, "https://imgur.com/image.png"}, "db=123&num_res=5&output_type=2&url=https%3A%2F%2Fimgur.com%2Fimage.png"},
		{"Without DB", &Config{"1", 0, false, 5}, args{&url.URL{}, "https://imgur.com/image.png"}, "api_key=1&db=999&num_res=5&output_type=2&url=https%3A%2F%2Fimgur.com%2Fimage.png"},
		{"Without results", &Config{"1", 999, false, 0}, args{&url.URL{}, "https://imgur.com/image.png"}, "api_key=1&db=999&output_type=2&url=https%3A%2F%2Fimgur.com%2Fimage.png"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.query(tt.args.uri, tt.args.imageURL); got != tt.want {
				t.Errorf("Config.query() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSengokuSearch(t *testing.T) {
	type args struct {
		imageURL string
	}
	tests := []struct {
		name    string
		s       *Sengoku
		args    args
		want    *Sauce
		wantErr bool
	}{
		{"Test", NewSengoku(), args{"https://api.kotori.love/pixiv/image/i.pximg.net/c/600x1200_90_webp/img-master/img/2020/11/14/17/43/50/85663331_p0_master1200.jpg"}, nil, false},
		{"Test", NewSengoku(), args{"https://6hf997t63vhfe.hjvvk1hhwfv4m.mangadex.network/99T_ZD6OgN-AYrbjhe5DtVxsaplJx05ZR7Yx8uhY5uOp4zlvy8hsv4MhiCTYPXEQ1g6eR763h4NxIRk4D5Nar3wUYn9k8Tfb8R1GDN85r7HCSr2-4rgYLJDdrjjDE1lDJvLZJDZm6PNGN791FKPqID5JhbtQ-Oweu61uxo1mfIxQK-7PkKuVLhcNv7A85HjjXBR6guUvfcY8GVeOyebW7B7cM34NmA/data/341ab55d1dd3c0a5c71b2813d0175de3/1-fd9f59c8e0adc0647af1d721886487ce1320a234395c8ab5e1a357302e1606d5.png"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Search(tt.args.imageURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sengoku.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sengoku.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
