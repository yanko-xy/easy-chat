/**
 * @author: Yanko/xiaoxiaoyang-sheep
 * @doc:
 **/

package logic

import (
	"context"
	"easy-chat/apps/user/rpc/user"
	"testing"
)

func TestRegisterLogic_Register(t *testing.T) {
	type args struct {
		in *user.RegisterReq
	}
	tests := []struct {
		name      string
		args      args
		wantPrint bool
		wantErr   bool
	}{
		{"1", args{in: &user.RegisterReq{
			Phone:    "13700000000",
			Nickname: "Yanko",
			Password: "123456",
			Avatar:   "demo.png",
			Sex:      0,
		}}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewRegisterLogic(context.Background(), svcCtx)
			got, err := l.Register(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantPrint {
				t.Log(tt.name, got)
			}
		})
	}
}