package localinfo

import (
	"fmt"
	"testing"
)

func TestIp(t *testing.T) {
	fmt.Println(GetLocalIPDotFormat())
}
