package controller

//go:generate easytags $GOFILE
type UserSubAddRequest struct {
	UserID int `json:"user_id"`
}

type UserSubDelRequest struct {
	UserID int `json:"user_id"`
}

type UserSubResponse struct {
	Success   bool `json:"success"`
	ErrorCode int  `json:"error_code,omitempty"`
}

type UserSubListResponse struct {
	SubUserIDs []int `json:"sub_user_ids"`
	Success    bool  `json:"success"`
	ErrorCode  int   `json:"error_code,omitempty"`
}
