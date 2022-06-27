package wap

type TicketListResult struct {
	Code    string         `json:"code"`
	Message interface{}    `json:"message"`
	Data    TicketListData `json:"data"`
	Success bool           `json:"success"`
}

type TicketList struct {
	Mid                 int     `json:"mid"`
	TicketsID           string  `json:"ticketsId"` // order use
	TimeID              string  `json:"timeId"`    // order use
	Title               string  `json:"title"`
	Type                int     `json:"type"`
	Price               float64 `json:"price"` // order use
	PriceExpress        float64 `json:"priceExpress"`
	MemberPrice         float64 `json:"memberPrice"`
	DiscountPrice       float64 `json:"discountPrice"`
	Describes           string  `json:"describes"`
	DescribeExpress     string  `json:"describeExpress"`
	DescribeElectronic  string  `json:"describeElectronic"`
	TimeStart           string  `json:"timeStart"`
	TimeEnd             string  `json:"timeEnd"`
	MemberTimeStart     string  `json:"memberTimeStart"`
	TimeEndExpress      string  `json:"timeEndExpress"`
	UseStart            string  `json:"useStart"`
	UseEnd              string  `json:"useEnd"`
	SaleRemindMinute    int     `json:"saleRemindMinute"`
	IsStudent           int     `json:"isStudent"`
	IsElectronic        int     `json:"isElectronic"` // order use
	IsExpress           int     `json:"isExpress"`    // order use
	SysDamai            int     `json:"sysDamai"`
	Counts              int     `json:"counts"`
	Status              int     `json:"status"`
	StatusExchange      int     `json:"statusExchange"`
	IsLackRegister      int     `json:"isLackRegister"`
	ExpressType         int     `json:"expressType"` // order use
	IsTrueName          int     `json:"isTrueName"`
	LimitCount          int     `json:"limitCount"`
	LimitCountMember    int     `json:"limitCountMember"`
	IsExclusive         int     `json:"isExclusive"`
	IsMember            int     `json:"isMember"`
	IsMemberStatus      int     `json:"isMemberStatus"`
	IsAgent             int     `json:"isAgent"`
	IsShowCode          int     `json:"isShowCode"`
	QrCodeShowTime      string  `json:"qrCodeShowTime"`
	AdvanceMinuteMember int     `json:"advanceMinuteMember"`
	TotalGeneral        int     `json:"totalGeneral"`
	TotalExchange       int     `json:"totalExchange"`
}

type TicketTimesList struct {
	Mid           int          `json:"mid"`
	TicketTimesID string       `json:"ticketTimesId"`
	Title         string       `json:"title"`
	Type          int          `json:"type"`
	PerformanceID string       `json:"performanceId"`
	TimeID        string       `json:"timeId"`
	UseStart      string       `json:"useStart"`
	UseEnd        string       `json:"useEnd"`
	TicketList    []TicketList `json:"ticketList"`
}

type PerformancesInfo struct {
	CityName  string `json:"city_name"`
	Title     string `json:"title"`
	AppStatus int    `json:"appStatus"`
	FieldName string `json:"field_name"`
}

type TicketListData struct {
	TicketTimesList  []TicketTimesList `json:"ticketTimesList"`
	PerformancesInfo PerformancesInfo  `json:"performancesInfo"`
}

type GetMobileCapturaResp struct {
	Code    string      `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

type GetEntersPeopleResp struct {
	Code    string                `json:"code"`
	Message interface{}           `json:"message"`
	Data    []GetEntersPeopleItem `json:"data"`
	Success bool                  `json:"success"`
}

type GetEntersPeopleItem struct {
	EntersID  string      `json:"entersId"`
	Type      int         `json:"type"`
	Name      string      `json:"name"`
	Mobile    string      `json:"mobile"`
	IDCard    string      `json:"idCard"`
	IsDefault bool        `json:"isDefault"`
	State     int         `json:"state"`
	CreatedAt string      `json:"createdAt"`
	UpdatedAt interface{} `json:"updatedAt"` // 可能为 nil
}

type SmsLoginResp struct {
	Code    string       `json:"code"`
	Message interface{}  `json:"message"`
	Data    SmsLoginData `json:"data"`
	Success bool         `json:"success"`
}

type SmsLoginData struct {
	Token         string      `json:"token"`
	UserInfo      UserInfo    `json:"userInfo"`
	UserMemberVo  interface{} `json:"userMemberVo"`
	WechatOpenid  interface{} `json:"wechatOpenid"`
	WechatUnionid interface{} `json:"wechatUnionid"`
}
type UserInfo struct {
	UID            string      `json:"uid"`
	Mobile         string      `json:"mobile"`
	Passwd         interface{} `json:"passwd"`
	Nickname       string      `json:"nickname"`
	State          int         `json:"state"`
	Sex            interface{} `json:"sex"`
	Birthday       interface{} `json:"birthday"`
	Area           interface{} `json:"area"`
	Signature      interface{} `json:"signature"`
	Avatar         string      `json:"avatar"`
	Background     string      `json:"background"`
	TagMe          interface{} `json:"tagMe"`
	CreateAt       string      `json:"createAt"`
	UpdatedAt      string      `json:"updatedAt"`
	ClosedAt       interface{} `json:"closedAt"`
	IsComplete     int         `json:"isComplete"`
	RongCloudToken interface{} `json:"rongCloudToken"`
	QrCode         string      `json:"qrCode"`
	StageMarker    int         `json:"stageMarker"`
	Province       string      `json:"province"`
	City           string      `json:"city"`
	County         interface{} `json:"county"`
}

type GetUserInfoResp struct {
	Code    string          `json:"code"`
	Message interface{}     `json:"message"`
	Data    GetUserInfoData `json:"data"`
	Success bool            `json:"success"`
}

type GetUserInfoData struct {
	UserInfo       UserInfo       `json:"userInfo"`
	RealNameInfo   interface{}    `json:"realNameInfo"`
	ThirdPartInfo  interface{}    `json:"thirdPartInfo"`
	MemberVo       MemberVo       `json:"memberVo"`
	UserMemberVo   interface{}    `json:"userMemberVo"`
	MemberJoinusVo MemberJoinusVo `json:"memberJoinusVo"`
}

type MemberVo struct {
	MemberID string `json:"memberId"`
	Name     string `json:"name"`
	Title    string `json:"title"`
	SubTitle string `json:"subTitle"`
	Icon     string `json:"icon"`
}
type MemberJoinusVo struct {
	QrCode          interface{} `json:"qrCode"`
	Title           string      `json:"title"`
	SubTitle        string      `json:"subTitle"`
	Cardface        string      `json:"cardface"`
	Type            int         `json:"type"`
	State           int         `json:"state"`
	InterestsDetail string      `json:"interestsDetail"`
	MemberNo        interface{} `json:"memberNo"`
}

type GetAddrListResp struct {
	Code    string            `json:"code"`
	Message interface{}       `json:"message"`
	Data    []GetAddrListData `json:"data"`
	Success bool              `json:"success"`
}

type GetAddrListData struct {
	AddressesID string      `json:"addressesId"`
	Name        string      `json:"name"`
	Phone       string      `json:"phone"`
	Zipcode     interface{} `json:"zipcode"`
	Province    string      `json:"province"`
	ProvinceID  string      `json:"provinceId"`
	City        string      `json:"city"`
	CityID      string      `json:"cityId"`
	County      string      `json:"county"`
	CountyID    string      `json:"countyId"`
	Address     string      `json:"address"`
	IsDefault   bool        `json:"isDefault"`
	State       int         `json:"state"`
	CreatedAt   string      `json:"createdAt"`
	UpdatedAt   string      `json:"updatedAt"`
}
