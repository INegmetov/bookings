package render

import (
	"github.com/inegmetov/bookings/pkg/config"
	"github.com/inegmetov/bookings/pkg/models"
	"net/http"
	"reflect"
	"testing"
	"text/template"
)

func TestAddDefaultData(t *testing.T) {
	type args struct {
		td *models.TemplateData
	}
	tests := []struct {
		name string
		args args
		want *models.TemplateData
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddDefaultData(tt.args.td); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDefaultData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTemplateCache(t *testing.T) {
	tests := []struct {
		name    string
		want    map[string]*template.Template
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateTemplateCache()
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTemplateCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateTemplateCache() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTemplates(t *testing.T) {
	type args struct {
		a *config.AppConfig
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewTemplates(tt.args.a)
		})
	}
}

func TestRenderTemplate(t *testing.T) {
	type args struct {
		w    http.ResponseWriter
		tmpl string
		td   *models.TemplateData
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RenderTemplate(tt.args.w, tt.args.tmpl, tt.args.td)
		})
	}
}
