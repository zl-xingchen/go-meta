package retry

import (
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	type args struct {
		max      uint16
		fn       Fn
		notifyFn NotifyFn
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{max: 5, fn: func() (bail bool, err error) {
				return false, errors.New("err")
			}, notifyFn: func(attempt uint16, err error) {
				t.Errorf("attempt:%d,err:%s", attempt, err.Error())
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Run(tt.args.max, tt.args.fn, tt.args.notifyFn); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
