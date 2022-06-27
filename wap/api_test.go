package wap

import (
	"testing"

	"github.com/k0kubun/pp"
)

func getTestZhengzaitv() *Zhengzaitv {
	cookies := "your cookies"
	z := NewZhengzaitv(cookies)
	return z
}

func TestZhengzaitv_GetTicketIds(t *testing.T) {
	type args struct {
		performanceId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "one",
			args:    args{"1158398181545697282777463"},
			wantErr: false,
		},
	}
	z := getTestZhengzaitv()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := z.GetTicketIds(tt.args.performanceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTicketIds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != 0 {
				t.Log("len is,", len(got))
				pp.Println(got[0])
			}
		})
	}
}

func TestGetEntersPeople(t *testing.T) {
	z := getTestZhengzaitv()
	got, err := z.GetEntersPeople()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", got)
}

func TestSmsLogin(t *testing.T) {
	z := getTestZhengzaitv()
	z.Cookies = "" // 手动设置为空
	mobile := "xxxx"
	code := "907548"
	got, err := z.SmsLogin(mobile, code)
	if err != nil {
		t.Error(err)
		return
	}
	pp.Println(got)
	z.Cookies = got.Token
	got1, err := z.GetEntersPeople()
	if err != nil {
		t.Error(err)
		return
	}
	pp.Println(got1)
}

func TestGetUserInfo(t *testing.T) {
	z := getTestZhengzaitv()
	userInfo, err := z.GetUserInfo()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", userInfo)
}

func TestGetAddrList(t *testing.T) {
	z := getTestZhengzaitv()
	resp, err := z.GetAddrList()
	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("%+v\n", resp)
}
