package types

func NewOkResp() BaseResp {
	return NewBaseResp(0, "ok")
}

func NewOkRespWithMsg(msg string) BaseResp {
	return NewBaseResp(0, msg)
}

func NewBaseResp(c int, msg string) BaseResp {
	return BaseResp{
		Code: c,
		Msg:  msg,
	}
}
