package rancher

import (
	"reflect"
	"testing"
)

func TestClient_GetNumberofProjects(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    int
		wantErr bool
	}{
		{"test-1", testClient, 4, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Client{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.GetNumberofProjects()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetNumberofProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetNumberofProjects() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetProjectAnnotations(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    []projectAnnotation
		wantErr bool
	}{
		{"test-1", testClient, []projectAnnotation{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Client{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.GetProjectAnnotations()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProjectAnnotations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProjectAnnotations() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetProjectLabels(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    []projectLabel
		wantErr bool
	}{
		{"test-1", testClient, []projectLabel{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Client{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.GetProjectLabels()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProjectLabels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProjectLabels() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_GetProjectResourceQuota(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		want    []projectResource
		wantErr bool
	}{
		{"test-1", testClient, []projectResource{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Client{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.GetProjectResourceQuota()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProjectResourceQuota() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProjectResourceQuota() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_clusterIdToName(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		args    string
		want    string
		wantErr bool
	}{
		{"test-1", testClient, "", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Client{
				Client: tt.fields.Client,
				Config: tt.fields.Config,
			}
			got, err := r.clusterIdToName(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("clusterIdToName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("clusterIdToName() got = %v, want %v", got, tt.want)
			}
		})
	}
}
