package hypergryph

import (
	"fmt"

	"github.com/starudream/go-lib/resty/v2"
)

const (
	Addr      = "https://as.hypergryph.com"
	UserAgent = "okhttp/4.11.0"

	AppCodeSKLAND = "4ca99fa6b56cc2ba"
)

type M map[string]any

type BaseResp[T any] struct {
	StatusCode *int   `json:"statusCode"`
	Error      string `json:"error"`
	Message    string `json:"message"`

	Status *int   `json:"status"`
	Type   string `json:"type"`
	Msg    string `json:"msg"`
	Data   T      `json:"data,omitempty"`
}

func (t *BaseResp[T]) IsSuccess() bool {
	return t != nil && t.Status != nil && *t.Status == 0
}

func (t *BaseResp[T]) String() string {
	if t == nil {
		return "<nil>"
	}
	if t.StatusCode != nil {
		return fmt.Sprintf("status: %d, error: %s, message: %s", *t.StatusCode, t.Error, t.Message)
	} else if t.Status != nil {
		return fmt.Sprintf("status: %d, type: %s, msg: %s", *t.Status, t.Type, t.Msg)
	}
	return "<nil>"
}

func R() *resty.Request {
	return resty.R().SetHeader("User-Agent", UserAgent).SetHeader("Accept-Encoding", "gzip")
}

func Exec[T any](r *resty.Request, method, path string) (t T, _ error) {
	res, err := resty.ParseResp[*BaseResp[any], *BaseResp[T]](
		r.SetError(&BaseResp[any]{}).SetResult(&BaseResp[T]{}).Execute(method, Addr+path),
	)
	if err != nil {
		return t, fmt.Errorf("[hypergryph] %w", err)
	}
	return res.Data, nil
}
