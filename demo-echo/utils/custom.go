package utils

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo/v4"
)

func unescapeUrl(dest *string, src string) {
	s, err := url.QueryUnescape(src)
	if err != nil {
		logrus.Errorf("Unescape error origin:%s\n", *dest)
	}
	*dest = s
}

func CustomeMiddleware(h echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		filterInfo := FilterInfo{}

		filter := c.QueryParam("filter")
		if filter != "" {
			err := json.Unmarshal([]byte(filter), &filterInfo)
			if err != nil {
				filterInfo := FilterInfo{}
				err := json.Unmarshal([]byte(filter), &filterInfo)
				if err != nil {
					logrus.Errorf("Transfomat filter info error! filter info:%s\n", filter)
				}
			}
		}

		sorting := c.QueryParam("sorting")
		if sorting != "" {
			s := make(map[string]string)
			err := json.Unmarshal([]byte(sorting), &s)
			if err != nil {
				logrus.Errorf("Transfomat sorting  error! sorting:%s\n", sorting)
			} else {
				filterInfo.Sorting = s
			}
		}

		sortBy := c.QueryParam("sortby")
		order := c.QueryParam("order")

		if sortBy != "" && order != "" {
			sortMap := make(map[string]string)
			sortMap[sortBy] = order
			filterInfo.Sorting = sortMap
		}

		page := c.QueryParam("page")
		pageSize := c.QueryParam("pageSize")

		if page != "" && pageSize != "" {
			p, err := strconv.Atoi(page)
			if err != nil {
				logrus.Errorf("Transfomat Page  error! Page:%s\n", page)
				filterInfo.Page = 1
			} else {
				filterInfo.Page = p
			}
			ps, err := strconv.Atoi(pageSize)
			if err != nil {
				logrus.Errorf("Transfomat PageSize  error! PageSize:%s\n", pageSize)
				filterInfo.PageSize = 10
			} else {
				filterInfo.PageSize = ps
			}
		} else {
			// filterInfo.Page = 1
			// filterInfo.PageSize = 10
		}

		cc := &CustomeContext{
			Context:    c,
			FilterInfo: &filterInfo,
		}

		return h(cc)
	}
}

func CustomeResponse(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		resp := next(c)
		myResp, ok := resp.(MyResponse)
		if !ok {
			return resp
		}
		if myResp.HTTPCode == 0 {
			myResp.HTTPCode = http.StatusOK
		}
		return c.Blob(myResp.HTTPCode, echo.MIMEApplicationJSONCharsetUTF8, []byte(myResp.Render()))
	}
}
