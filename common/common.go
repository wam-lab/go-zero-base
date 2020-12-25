package common

import "github/yguilai/timetable-micro/gateway/internal/types"

func NewOkResp() types.BaseResp {
	return NewBaseResp(0, "ok")
}

func NewOkRespWithMsg(msg string) types.BaseResp {
	return NewBaseResp(0, msg)
}

func NewBaseResp(c int, msg string) types.BaseResp {
	return types.BaseResp{
		Code: c,
		Msg:  msg,
	}
}
