// Code generated by goctl. DO NOT EDIT.
package types

type UserReq struct {
	Id string `form:"id"`
}

type UserResp struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}