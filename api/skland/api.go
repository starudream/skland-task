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

	"github.com/go-viper/mapstructure/v2"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/utils/osutil"
	"github.com/starudream/go-lib/resty/v2"

	"github.com/starudream/skland-task/config"
)

const (
	Addr      = "https://zonai.skland.com"
	UserAgent = "Skland/1.5.1 (com.hypergryph.skland; build:100501001; Android 33; ) Okhttp/4.11.0"
	Platform  = "1"
	VName     = "1.5.1"
	DId       = "743a446c83032899"

	GameIdArknights = "1"   // 明日方舟
	GameIdExastris  = "2"   // 来自星尘
	GameIdEndfleld  = "3"   // 明日方舟：终末地
	GameIdPopucom   = "4"   // 泡姆泡姆
	GameIdNest      = "100" // 纳斯特港
	GameIdOpenup    = "101" // 开拓芯

	GameAppCodeArknights = "arknights"

	MessageForumHasSigned = "重复签到"
	MessageGameHasSigned  = "请勿重复签到！"
)

type BaseResp[T any] struct {
	Code    *int   `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

func (t *BaseResp[T]) GetMessage() string {
	if t == nil {
		return ""
	}
	return t.Message
}

func (t *BaseResp[T]) IsSuccess() bool {
	return t != nil && t.Code != nil && *t.Code == 0
}

func (t *BaseResp[T]) String() string {
	if t == nil || t.Code == nil {
		return "<nil>"
	}
	return fmt.Sprintf("code: %d, message: %s", *t.Code, t.Message)
}

func IsMessage(err error, msg string) bool {
	if err == nil {
		return false
	}
	e, ok1 := resty.AsRespErr(err)
	if ok1 {
		t1, ok2 := e.Response.Error().(interface{ GetMessage() string })
		t2, ok3 := e.Response.Result().(interface{ GetMessage() string })
		return (ok2 && t1.GetMessage() == msg) || (ok3 && t2.GetMessage() == msg)
	}
	return false
}

func IsUnauthorized(err error) bool {
	re, ok := resty.AsRespErr(err)
	if ok {
		return re.StatusCode() == 401
	}
	return false
}

func R() *resty.Request {
	return resty.R().SetHeader("User-Agent", UserAgent).SetHeader("Accept-Encoding", "gzip")
}

func Exec[T any](r *resty.Request, method, path string, vs ...any) (t T, _ error) {
	for i := 0; i < len(vs); i++ {
		switch v := vs[i].(type) {
		case config.AccountSkland:
			AddSign(r, method, path, v)
		}
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

func AddSign(r *resty.Request, method, path string, skland config.AccountSkland) {
	ts := strconv.FormatInt(time.Now().Unix(), 10)

	// use struct to fix the order of headers
	headers := signHeaders{Platform: Platform, Timestamp: ts, DId: DId, VName: VName}

	r.SetHeaders(tom(headers))

	_, signature := sign(headers, method, path, skland.Token, r.QueryParam, r.Body)

	r.SetHeader("cred", skland.Cred)
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
	out := map[string]string{}
	cfg := &mapstructure.DecoderConfig{
		Squash:  true,
		Result:  &out,
		TagName: "json",
	}
	decoder, err := mapstructure.NewDecoder(cfg)
	osutil.PanicErr(err)
	err = decoder.Decode(s)
	osutil.PanicErr(err)
	return out
}

func hmac256(key, content string) []byte {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(content))
	return h.Sum(nil)
}
