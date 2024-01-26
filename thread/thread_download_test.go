package thread

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestThreadDownload(t *testing.T) {
	start := time.Now()
	res := GetImgContentBusiness([]string{"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/7263ceb2eff523936870961b07278556_1698822856773.jpeg",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png"}, "listId_1234")
	end := time.Since(start)
	fmt.Println(len(res), "end:", end)
	fmt.Println(res)
}

func TestNoThread(t *testing.T) {
	arr := []string{"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/7263ceb2eff523936870961b07278556_1698822856773.jpeg",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png",
		"http://imghash-list-1251671073.cos.ap-beijing.myqcloud.com/20231101/19f135dd8409014e8591b47dfb21d0d2_1698823093090.png"}
	busList := []string{}
	for _, v := range arr {
		files, _ := ioutil.TempFile("", "tmpfile")
		start := time.Now()
		_ = WgetFile(v, files.Name())
		end := time.Since(start)
		fmt.Println("下载耗时：", fmt.Sprintf("%v", end))
		//fmt.Println("filName:", files.Name())
		start = time.Now()
		base64Num, _ := ImageFileBase64(files.Name())
		end = time.Since(start)
		fmt.Println("计算耗时:", fmt.Sprintf("%v", end))
		if base64Num != "" {
			busList = append(busList, GetStringMd5(base64Num+"listId_1234"))
		}
		_ = os.Remove(files.Name())
	}
}
// 2^k - 1 = n
//
func TestFile(t *testing.T) {
	result := []string{}
	b, err := ioutil.ReadFile("/Users/zhaodeman/Downloads/广告词1.txt")
	if err != nil {
		fmt.Println(err)
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		result = append(result, lineStr)
	}
	fmt.Println(result)
	for _,v := range result {
		fmt.Println(v)
	}
	return
}