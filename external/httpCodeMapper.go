package SharedPackages

import (
	"encoding/json"
	"net/http"
)

type HttpCodeMapper struct {
	httpCodeMap map[int]int
}

type HttpResult[T any] struct {
	httpCode          int
	applicationResult *ApplicationResult[T]
}

func (hr *HttpResult[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{

		"Result": hr.applicationResult,
	})
}

func (httpresult *HttpResult[any]) HttpCode() int {
	return httpresult.httpCode
}

func (httpresult *HttpResult[any]) Body() *ApplicationResult[any] {
	return httpresult.applicationResult
}

func NewHttpCodeMapper(httpCodeMap map[int]int) *HttpCodeMapper {
	return &HttpCodeMapper{httpCodeMap: httpCodeMap}
}

func (codemapper *HttpCodeMapper) getHttpCode(applicationcode int) int {
	httpcode, ok := codemapper.httpCodeMap[applicationcode]
	if !ok {
		return http.StatusInternalServerError
	}
	return httpcode
}

func GetHttpResponse[T any](s *HttpCodeMapper, applicationResult *ApplicationResult[T]) HttpResult[T] {
	httpcode := s.getHttpCode(applicationResult.StatusCode)
	if httpcode == http.StatusInternalServerError {
		return HttpResult[T]{httpCode: httpcode, applicationResult: nil}
	} else {
		return HttpResult[T]{httpCode: httpcode, applicationResult: applicationResult}
	}
}

func Example() {

	httpcodemap := NewHttpCodeMapper(map[int]int{1: 405})

	sstp := NewApplicationResult("omar", 1, "ASdas")

	GetHttpResponse[string](httpcodemap, &sstp)

}
