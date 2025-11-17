package core

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
)

// translate 翻译成 HTTP 请求
func (config *Config) translate(c context.Context, req *Request) (*http.Request, error) {
	body := req.Body

	// 获取请求体
	contentType, rawBody, err := payload(body, config.serializable)
	if err != nil {
		return nil, err
	}

	// 获取路径参数
	var pathes []string
	for _, path := range strings.Split(req.ApiPath, "/") {
		// 替换路径参数
		if strings.Index(path, ":") == 0 {
			name := path[1:]
			val := req.PathParams.Get(name)

			if val == "" {
				return nil, fmt.Errorf("http path:%s, name: %s, value is empty", req.ApiPath, name)
			}

			pathes = append(pathes, url.PathEscape(val))
			continue
		}

		// 添加路径
		pathes = append(pathes, path)
	}

	// 拼接路径
	newPath := strings.Join(pathes, "/")
	if strings.Index(newPath, "http") != 0 {
		newPath = fmt.Sprintf("%s%s", config.BaseUrl, newPath)
	}

	// 拼接查询参数
	queryPath := req.QueryParams.Encode()
	if queryPath != "" {
		newPath = fmt.Sprintf("%s?%s", newPath, queryPath)
	}

	// 创建 HTTP 请求
	return config.newHTTPRequest(c, req.AccessTokenType(), req.HttpMethod, newPath, contentType, rawBody)
}

// newHTTPRequest 创建 HTTP 请求
func (config *Config) newHTTPRequest(
	c context.Context,
	accessTokenType AccessTokenType,
	httpMethod, url, contentType string, body []byte) (*http.Request, error) {

	// 创建 HTTP 请求
	httpRequest, err := http.NewRequestWithContext(c, httpMethod, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// 添加请求头
	for k, vs := range config.Header {
		for _, v := range vs {
			httpRequest.Header.Add(k, v)
		}
	}

	// 设置用户代理
	httpRequest.Header.Set(userAgentHeader, userAgent())

	// 设置内容类型
	if contentType != "" {
		httpRequest.Header.Set(httpHeaderContentType, contentType)
	}

	var accessToken string
	switch accessTokenType {

	// 应用访问令牌
	case AccessTokenTypeApp:
		if accessToken, err = config.getAppAccessToken(c); err != nil {
			return nil, err
		}

	// 租户访问令牌
	case AccessTokenTypeTenant:
		if accessToken, err = config.getTenantAccessToken(c); err != nil {
			return nil, err
		}

	// 用户访问令牌
	case AccessTokenTypeUser:
		return nil, fmt.Errorf("user access token not supported")

	}
	authorizationToHeader(httpRequest, accessToken)

	return httpRequest, nil
}

// payload 获取请求体
func payload(body any, serializable Serializable) (string, []byte, error) {
	// 如果为空，返回默认内容类型和空内容
	if body == nil {
		return defaultContentType, nil, nil
	}

	// 如果为表单数据，返回表单数据
	if form, ok := body.(*Formdata); ok {
		return form.content()
	}

	// 序列化 body
	bytes, err := serializable.Serialize(body)
	return httpHeaderContentTypeJson, bytes, err
}

type Formdata struct {
	fields map[string]any
	data   *struct {
		content     []byte
		contentType string
	}
}

// content 获取表单数据
func (form *Formdata) content() (contentType string, content []byte, err error) {
	// 如果数据不为空，则返回数据类型和内容
	if form.data != nil {
		return form.data.contentType, form.data.content, nil
	}

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	for key, val := range form.fields {
		if r, ok := val.(io.Reader); ok {
			// 创建表单文件
			part, err := writer.CreateFormFile(key, "unknown-file")
			if err != nil {
				return "", nil, err
			}

			if _, err = io.Copy(part, r); err != nil {
				return "", nil, err
			}
		} else {
			// 创建表单字段
			writer.WriteField(key, fmt.Sprintf("%v", val))
		}
	}

	// 获取表单内容类型
	contentType = writer.FormDataContentType()
	if err := writer.Close(); err != nil {
		return "", nil, err
	}

	form.data = &struct {
		content     []byte
		contentType string
	}{content: buf.Bytes(), contentType: contentType}
	return form.data.contentType, form.data.content, nil
}
