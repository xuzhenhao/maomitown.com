package user

// CreateRequest request消息体参数
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// CreateResponse response消息体参数
type CreateResponse struct {
	Username string `json:"username"`
}
