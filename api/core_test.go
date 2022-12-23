package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

////////// CoreDirectorTester

type CoreDirectorTester struct{}

func (tester *CoreDirectorTester) start(t *testing.T, proxy *OchanocoProxy, proxyServ *httptest.Server, testServ *httptest.Server) {
}

func (tester *CoreDirectorTester) director(t *testing.T, url string) OchanocoDirector {
	return makesSimpleDirector(t, url)
}
func (tester *CoreDirectorTester) modifyResp(t *testing.T) OchanocoModifyResponse {
	return makeEmptyModifyResp()
}
func (tester *CoreDirectorTester) testServers(t *testing.T) (*httptest.Server, *httptest.Server, *httptest.Server) {
	return makeSimpleServers()
}

func (tester *CoreDirectorTester) request(t *testing.T, url string) *http.Response {
	return requestGetforTest(t, url)
}
func (tester *CoreDirectorTester) check(t *testing.T, resp *http.Response) {
	checkResponseWithBody(t, resp, TEST_RESP_BODY1)
}

func TestCoreDirector(t *testing.T) {
	tester := CoreDirectorTester{}
	runCommonTest(t, &tester, "core/director")
}

// //////// CoreModifyResponseTester
type CoreModifyResponseTester struct{}

func (tester *CoreModifyResponseTester) start(t *testing.T, proxy *OchanocoProxy, proxyServ *httptest.Server, testServ *httptest.Server) {
}

func (tester *CoreModifyResponseTester) director(t *testing.T, url string) OchanocoDirector {
	return makesSimpleDirector(t, url)
}
func (tester *CoreModifyResponseTester) modifyResp(t *testing.T) OchanocoModifyResponse {
	return makesSimpleModifyResp()
}
func (tester *CoreModifyResponseTester) testServers(t *testing.T) (*httptest.Server, *httptest.Server, *httptest.Server) {
	return makeSimpleServers()
}

func (tester *CoreModifyResponseTester) request(t *testing.T, url string) *http.Response {
	return requestGetforTest(t, url)
}
func (tester *CoreModifyResponseTester) check(t *testing.T, resp *http.Response) {
	checkResponseWithBody(t, resp, TEST_RESP_BODY2)
}

func TestCoreModifyResp(t *testing.T) {
	tester := CoreModifyResponseTester{}
	runCommonTest(t, &tester, "core/modify_resp")
}
