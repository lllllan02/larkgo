package core

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"strings"
)

// DoRequest 发送请求
func (config *Config) DoRequest(c context.Context, req *Request) (res *Response, err error) {
	for i := 0; i < 2; i++ {
		// 翻译请求
		httpRequest, err := config.translate(c, req)
		if err != nil {
			return nil, err
		}
		if config.LogReqAtDebug {
			config.logger.Debugf("req:%v", Prettify(httpRequest))
		} else {
			config.logger.Debugf("req:%s,%s", req.HttpMethod, req.ApiPath)
		}

		// 发送请求
		res, err = config.doSend(httpRequest)
		if config.LogReqAtDebug {
			config.logger.Debugf("resp:%v", Prettify(res))
		}
		if err != nil {
			if _, isDialError := err.(*DialFailedError); isDialError {
				continue
			}
			return nil, err
		}

		// 如果响应体不是 JSON 格式，则不进行重试
		if !strings.Contains(res.Header.Get(httpHeaderContentType), httpHeaderContentTypeJson) {
			break
		}

		// 解析响应体
		codeError := &CodeError{}
		if err = config.serializable.Deserialize(res.RawBody, codeError); err != nil {
			return nil, err
		}

		// 如果错误码不是访问令牌无效错误码，则不进行重试
		code := codeError.Code
		if code != errCodeAccessTokenInvalid &&
			code != errCodeAppAccessTokenInvalid &&
			code != errCodeTenantAccessTokenInvalid {
			break
		}

		// 如果访问令牌类型为空，则不进行重试
		if req.AccessTokenType() == AccessTokenTypeNone {
			break
		}
	}

	return res, err
}

func (config *Config) doSend(req *http.Request) (*Response, error) {
	res, err := config.httpClient.Do(req)
	if err != nil {
		if er, ok := err.(*url.Error); ok {
			// 客户端超时错误
			if er.Timeout() {
				return nil, &ClientTimeoutError{msg: er.Error()}
			}

			// 连接失败错误
			if e, ok := er.Err.(*net.OpError); ok && e.Op == "dial" {
				return nil, &DialFailedError{msg: er.Error()}
			}
		}
		return nil, err
	}

	// 服务器超时错误
	if res.StatusCode == http.StatusGatewayTimeout {
		logID := res.Header.Get(httpHeaderLogId)
		if logID == "" {
			logID = res.Header.Get(httpHeaderRequestId)
		}
		config.logger.Infof("req path:%s, server time out,requestId:%s", req.URL.RequestURI(), logID)
		return nil, &ServerTimeoutError{msg: "server time out error"}
	}

	// 读取响应体
	body, err := readResponse(res)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: res.StatusCode,
		Header:     res.Header,
		RawBody:    body,
	}, nil
}
