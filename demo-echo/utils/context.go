package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CustomeContext struct {
	echo.Context
	*FilterInfo
}

type FilterInfo struct {
	Type       string            `json:"type"`
	Conditions [][]string        `json:"conditions"`
	Sorting    map[string]string `json:"sorting"`
	PageSize   int               `json:"page_size"`
	Page       int               `json:"page"`
}

type PagingResponse struct {
	Page       int         `json:"page"`
	PageSize   int         `json:"pageSize"`
	TotalCount int         `json:"totalCount"`
	Results    interface{} `json:"results"`
}

type MyResponse struct {
	HTTPCode     int         `json:"-"`
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	Success      bool        `json:"success"`
	Data         interface{} `json:"data"`
	ResourceName string      `json:"resource_name,omitempty"`
	ResourceID   string      `json:"resource_id,omitempty"`
}

func (response MyResponse) Error() string {
	return response.Render()
}

func (response MyResponse) Render() string {
	resp, err := json.Marshal(response)
	if err != nil {
		return fmt.Sprintf(`
		{
			"message": "%s",
			"code": 500,
			"success": false
		}
		`, err)
	}
	return string(resp)
}

func NewErrorResponse(code int, message string) MyResponse {
	return MyResponse{
		Code:    code,
		Message: message,
	}
}

func NewResponse(response interface{}) MyResponse {
	return MyResponse{
		Data:    response,
		Code:    200,
		Success: true,
	}
}

func NewPagingResponse(page, pageSize, total int, response interface{}) MyResponse {

	return MyResponse{
		Data: PagingResponse{
			Page:       page,
			PageSize:   pageSize,
			TotalCount: total,
			Results:    response,
		},
		Code:    200,
		Success: true,
	}
}

func NewResponseWithLog(code int, message string, success bool, response interface{}, resourceName string, resourceID string) MyResponse {
	return MyResponse{
		Code:         code,
		Message:      message,
		Success:      success,
		Data:         response,
		ResourceID:   resourceID,
		ResourceName: resourceName,
	}
}

func New500ErrResp(message string) MyResponse {
	if message == "" {
		message = "服务器内部错误"
	}
	return MyResponse{
		Code:    http.StatusInternalServerError,
		Success: false,
		Message: message,
	}
}
