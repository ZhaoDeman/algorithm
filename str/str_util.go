package str

import "fmt"

func IsTranslation(start, target string) bool {
	i := 0
	j := 1
	fmt.Println('R' == 'L' == (i < j))
	return j == len(target)
}
