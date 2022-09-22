package gen

import (
	"fmt"
	"testing"
)

func TestGetInvCodeByUID(t *testing.T) {
	fmt.Println(GetInvCodeByUIDUniqueNew(1, 6)) // SVPZ52
	fmt.Println(GetInvCodeByUIDUniqueNew(2, 6)) // VC3AB5
}

func TestGetInvCodeByUID1(t *testing.T) {
	fmt.Println(GetInvCodeByUID(1, 6)) // 3H9V38
	fmt.Println(GetInvCodeByUID(2, 6)) // 4CWSAG
}
