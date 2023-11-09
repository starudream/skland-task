package skland

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/utils/structutil"
	"github.com/starudream/go-lib/resty/v2"
)

const (
	Addr      = "https://zonai.skland.com"
	UserAgent = "Skland/1.1.0 (com.hypergryph.skland; build:100100047; Android 33; ) Okhttp/4.11.0"
	Platform  = "1"
	VName     = "1.1.0"
	DId       = "743a446c83032899"

	GameCodeArknights = "1" // 明日方舟
	GameCodeExastris  = "2" // 来自星尘
	GameCodeEndfleld  = "3" // 终末地
	GameCodePopucom   = "4" // 泡姆泡姆

	GameAppCodeArknights = "arknights"
)

var (
	GameCodeByAppCode = map[string]string{
		GameAppCodeArknights: GameCodeArknights,
	}
	GameNameByAppCode = map[string]string{
		GameAppCodeArknights: "明日方舟",
	}
)

type M map[string]any

type BaseResp[T any] struct {
	Code    *int   `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func (t *BaseResp[T]) GetCodeMsg() (int, string) {
	if t == nil || t.Code == nil {
		return 999999, t.Message
	}
	return *t.Code, t.Message
}

func (t *BaseResp[T]) IsSuccess() bool {
	return t != nil && t.Code != nil && *t.Code == 0
}

func (t *BaseResp[T]) String() string {
	return fmt.Sprintf("code: %d, message: %s", *t.Code, t.Message)
}

func IsCode(err error, code int, msg string) bool {
	if err == nil {
		return false
	}
	e, ok1 := resty.AsRespErr(err)
	if ok1 {
		t, ok2 := e.Result().(interface{ GetCodeMsg() (int, string) })
		if ok2 {
			c, m := t.GetCodeMsg()
			return c == code && (msg == "" || m == msg)
		}
	}
	return false
}

func R() *resty.Request {
	return resty.R().SetHeader("User-Agent", UserAgent).SetHeader("Accept-Encoding", "gzip")
}

func Exec[T any](r *resty.Request, method, path string, token ...string) (t T, _ error) {
	if len(token) > 0 && token[0] != "" {
		AddSign(r, method, path, token[0])
	}
	res, err := resty.ParseResp[*BaseResp[any], *BaseResp[T]](
		r.SetError(&BaseResp[any]{}).SetResult(&BaseResp[T]{}).Execute(method, Addr+path),
	)
	if err != nil {
		return t, fmt.Errorf("[skland] %w", err)
	}
	return res.Data, nil
}

type signHeaders struct {
	Platform  string `json:"platform"`
	Timestamp string `json:"timestamp"`
	DId       string `json:"dId"`
	VName     string `json:"vName"`
}

func AddSign(r *resty.Request, method, path, token string) {
	ts := strconv.FormatInt(time.Now().Unix(), 10)

	// use struct to fix the order of headers
	headers := signHeaders{Platform: Platform, Timestamp: ts, DId: DId, VName: VName}

	r.SetHeaders(tom(headers))

	_, signature := sign(headers, method, path, token, r.QueryParam, r.Body)

	r.SetHeader("sign", signature)
}

func sign(headers signHeaders, method, path, token string, query url.Values, body any) (string, string) {
	str := query.Encode()
	if method != "GET" {
		str = json.MustMarshalString(body)
	}

	content := path + str + headers.Timestamp + json.MustMarshalString(headers)

	b1 := hmac256(token, content)
	s1 := hex.EncodeToString(b1)
	b2 := md5.Sum([]byte(s1))
	s2 := hex.EncodeToString(b2[:])

	return content, s2
}

func tom(s any) map[string]string {
	t := structutil.New(s)
	t.TagName = "json"
	m := map[string]string{}
	for k, v := range t.Map() {
		m[k] = v.(string)
	}
	return m
}

func hmac256(key, content string) []byte {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(content))
	return h.Sum(nil)
}

func IsUnauthorized(err error) bool {
	re, ok := resty.AsRespErr(err)
	if ok {
		return re.StatusCode() == 401
	}
	return false
}
