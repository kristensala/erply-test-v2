package models

type ApiResponse struct {
    IsSuccess bool `json:"isSuccess" default:"true"`
    Data interface{} `json:"data"`
    Message string `json:"message" default:""`
}
