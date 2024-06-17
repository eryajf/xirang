package response

type DashboardList struct {
	DataType  string `json:"dataType"`
	DataName  string `json:"dataName"`
	DataCount int64  `json:"dataCount"`
	Icon      string `json:"icon"`
	Path      string `json:"path"`
}

// UserLoginRsp 用户登录响应
type UserLoginRsp struct {
	Username     string   `json:"username"`
	Nickname     string   `json:"nickname"`
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	Expires      string   `json:"expires"`
	Roles        []string `json:"roles"`
}

// UserRefreshTokenRsp 用户刷新token响应
type UserRefreshTokenRsp struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Expires      string `json:"expires"`
}
