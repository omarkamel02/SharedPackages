package SharedPackages

import "net/http"

type HttpCodeMapper struct {
	httpCodeMap map[int]int
}

type HttpResult struct {
	HttpCode          int
	ApplicationResult *ApplicationResult[any]
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
		return HttpResult{HttpCode: httpcode, ApplicationResult: nil}
	} else {
		return HttpResult{HttpCode: httpcode, ApplicationResult: applicationResult}
	}
}
