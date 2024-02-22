package SharedPackages

import "net/http"

type HttpCodeMapper struct {
	httpCodeMap map[int]int
}

type HttpResult struct {
	httpCode          int
	applicationResult *ApplicationResult[any]
}

func (httpresult *HttpResult) HttpCode() int {
	return httpresult.httpCode
}

func (httpresult *HttpResult) Body() *ApplicationResult[any] {
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

func (codemapper *HttpCodeMapper) GetHttpResponse(applicationResult *ApplicationResult[any]) HttpResult {
	httpcode := codemapper.getHttpCode(applicationResult.StatusCode)
	if httpcode == http.StatusInternalServerError {
		return HttpResult{httpCode: httpcode, applicationResult: nil}
	} else {
		return HttpResult{httpCode: httpcode, applicationResult: applicationResult}
	}
}
