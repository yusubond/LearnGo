package subond

import (
	"testing"
)

func TestNetParseIP(t *testing.T) {
	if NetParseIP("202.192.168.14") == true {
		t.Error("mynet.NetParseIP() is ok")
	}
}
