// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoReq struct {
	UserId int64 `json:"UserId"`
}

type UserInfoResp struct {
	UserId   int64  `json:"UserId"`
	Nickname string `json:"nickname"`
}
