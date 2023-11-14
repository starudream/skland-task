package skland

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/utils/testutil"
)

const PlayerUid = ""

func TestSign(t *testing.T) {
	headers := signHeaders{
		Platform:  Platform,
		Timestamp: "1699186712",
		DId:       DId,
		VName:     VName,
	}
	content, signature := sign(headers, "GET", "/api/v1/user", "4fe483e3b7d043afb51b3f0919f3fb2b", nil, nil)
	testutil.Equal(t, `/api/v1/user1699186712{"platform":"1","timestamp":"1699186712","dId":"743a446c83032899","vName":"1.1.0"}`, content)
	testutil.Equal(t, "2fb7d2969c7e4552f1387871ff2cc0d5", signature)
}
