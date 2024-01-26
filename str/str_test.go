package str

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	IsTranslation("11","22")
}


// // len(be-contact-model-new.textResult.url.DEFINITENESS)
func TestOperatorStr(t *testing.T) {
	str := "len(be-contact-model-new.textResult.url.DEFINITENESS)"
	c := 0
	hash := map[string]bool{}
	for i := 0;i<len(str);i++ {
		for j := i+1;j<len(str);j++ {
//			fmt.Println(str[i:j])
			if !hash[str[i:j]] {
				c++
				hash[str[i:j]] = true
			}
		}
	}
	fmt.Println(c)
}