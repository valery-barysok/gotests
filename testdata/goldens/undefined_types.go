package undefinedtypes

import (
	"reflect"
	"testing"
	"time"

	"github.com/gojuno/minimock"
)

func TestUndefined_Do(t *testing.T) {
	type args struct {
		es Something
	}
	tests := []struct {
		name    string
		u       *Undefined
		args    args
		want    *Unknown
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		mc := minimock.NewController(t)
		defer mc.Wait(time.Second)
		got, err := tt.u.Do(tt.args.es)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Undefined.Do() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("%q. Undefined.Do() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
