package storage

import (
	"reflect"
	"testing"

	"github.com/google/go-github/v26/github"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		want *client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_SaveFile(t *testing.T) {
	type fields struct {
		Client *github.Client
	}
	type args struct {
		filename string
		content  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "Commit without error",
			fields:  fields{Client: NewClient().Client},
			args:    args{"test-note.md", "# This is a test note"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				Client: tt.fields.Client,
			}
			if err := c.SaveFile(tt.args.filename, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("client.SaveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
