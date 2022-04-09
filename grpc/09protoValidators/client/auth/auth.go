package auth

import "context"

/*
@author RandySun
@create 2022-03-29-10:50
*/
// Token token认证
type Token struct {
	Value string
}

const headerAuthorize string = "authorization"

// GetRequestMetadata 获取当前请求认证所需的元数据
func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{headerAuthorize: t.Value}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 认证进行安全传输, 返回false无需TLS认证
func (t *Token) RequireTransportSecurity() bool {
	return true
}
