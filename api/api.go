package api

type CreateURLReq struct {
	URL string `json:"url"`
}

type AccessURLReq struct {
	URL string `json:"url"`
}

type ServiceResponse struct {
	code int32  `json:"code"`
	msg  string `json:"msg"`
}
