package wap

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
)

type Zhengzaitv struct {
	Cookies  string
	UserInfo *UserInfo
	Client   *resty.Client
}

func NewZhengzaitv(cookies string) *Zhengzaitv {
	return &Zhengzaitv{
		Cookies: cookies,
		Client:  resty.New(),
	}
}

// GetTicketIds 获取票种列表
func (z *Zhengzaitv) GetTicketIds(performanceId string) ([]TicketList, error) {
	// 获取票种 id 列表
	if performanceId == "" {
		return nil, fmt.Errorf("场次 id 为空")
	}

	url := fmt.Sprintf("https://kylin.zhengzai.tv/kylin/performance/partner/%s", performanceId)
	resp, err := z.Client.R().
		SetQueryParam("isAgent", "0").
		SetResult(&TicketListResult{}).
		SetHeaders(z.GetKylinHeaders()).
		Get(url)
	if err != nil {
		log.Println("错误 resp: ", resp.String())
		return nil, err
	}

	v, ok := resp.Result().(*TicketListResult)
	if !ok {
		return nil, fmt.Errorf("获取票种列表 断言失败")
	}

	if !v.Success {
		return nil, fmt.Errorf("接口请求失败 success: false")
	}

	res := make([]TicketList, 0)

	for i := 0; i < len(v.Data.TicketTimesList); i++ {
		timesList := v.Data.TicketTimesList[i]
		res = append(res, timesList.TicketList...)
	}

	return res, nil
}

// GetEntersPeople 获取用户观演人
func (z *Zhengzaitv) GetEntersPeople() ([]GetEntersPeopleItem, error) {
	url := "https://adam.zhengzai.tv/adam/enters/list"
	resp, err := z.Client.R().
		SetResult(&GetEntersPeopleResp{}).
		SetHeaders(z.GetAdamHeaders()).
		Get(url)
	if err != nil {
		log.Println("错误 resp: ", resp.String())
		return nil, err
	}

	v, ok := resp.Result().(*GetEntersPeopleResp)
	if !ok {
		return nil, fmt.Errorf("获取观演人列表 断言失败")
	}

	return v.Data, nil
}

// SmsLogin 用户登录接口
func (z *Zhengzaitv) SmsLogin(mobile, code string) (*SmsLoginData, error) {
	url := "https://adam.zhengzai.tv/adam/login/sms"
	resp, err := z.Client.R().
		SetFormData(map[string]string{
			"mobile": mobile,
			"code":   code,
		}).
		SetResult(&SmsLoginResp{}).
		SetHeaders(z.GetSmsLoginHeaders()).
		Post(url)
	if err != nil {
		log.Println("错误 resp: ", resp.String())
		return nil, err
	}

	v, ok := resp.Result().(*SmsLoginResp)
	if !ok {
		return nil, fmt.Errorf("登录用户 断言失败")
	}

	if !v.Success {
		return nil, fmt.Errorf("登录失败, success: false")
	}

	return &v.Data, nil
}

// InjectUserInfo 注入用户信息
func (z *Zhengzaitv) InjectUserInfo() error {
	info, err := z.GetUserInfo()
	if err != nil {
		return err
	}

	z.UserInfo = info
	return nil
}

// 获取用户个人信息
func (z *Zhengzaitv) GetUserInfo() (*UserInfo, error) {
	url := "https://adam.zhengzai.tv/adam/user/info"

	resp, err := z.Client.R().
		SetResult(&GetUserInfoResp{}).
		SetHeaders(z.GetAdamHeaders()).
		Post(url)
	if err != nil {
		log.Println("错误 resp: ", resp.String())
		return nil, err
	}

	v, ok := resp.Result().(*GetUserInfoResp)
	if !ok {
		return nil, fmt.Errorf("获取用户个人信息 断言失败")
	}

	if !v.Success {
		return nil, fmt.Errorf("获取用户个人信息 success: false")
	}

	return &v.Data.UserInfo, nil
}

func (z *Zhengzaitv) GetAddrList() ([]GetAddrListData, error) {
	url := "https://adam.zhengzai.tv/adam/addr/list"

	resp, err := z.Client.R().
		SetResult(&GetAddrListResp{}).
		SetHeaders(z.GetAdamHeaders()).
		Get(url)

	if err != nil {
		log.Println("错误 resp: ", resp.String())
		return nil, err
	}

	v, ok := resp.Result().(*GetAddrListResp)
	if !ok || !v.Success {
		return nil, fmt.Errorf("获取地址列表失败")
	}

	return v.Data, nil
}
