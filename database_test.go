package skirmish

import (
	"testing"
)

// TODO(sbrow): Fix TestConnect.
func TestConnect(t *testing.T) {
	type args struct {
		Host string
		Port int
		User string
		Name string
		SSL  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"standard", args{"localhost", 5432, "postgres", "skirmish", "disable"}, false},
		{"ssl_Enabled", args{"localhost", 5432, "postgres", "skirmish", "require"}, false},

		// TODO(sbrow): find out why connecting to database with non-existant user/database doesn't throw an error.
		{"wrong_user", args{"localhost", 5432, "butts", "skirmish", "disable"}, false},
		{"wrong_database", args{"localhost", 5432, "postgres", "skirmish23", "disable"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Connect(tt.args.Host, tt.args.Port, tt.args.Name, tt.args.User, tt.args.SSL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
