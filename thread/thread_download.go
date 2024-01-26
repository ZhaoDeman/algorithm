package thread

import (
	"crypto/md5"
	crand "crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

type ImgContentBusiness struct {
	BusinessList []string `json:"businessList"`
}

func GetImgContentBusiness(contents []string, listId string) (ans []string) {
	num := 1
	resBusiness := make(chan *ImgContentBusiness, 1)
	closeCh := make(chan struct{}, 1)
	index := 0

	for index+num < len(contents) {
		go downloadListContentUrl(contents[index:index+num], listId, resBusiness)
		index += num
	}
	go downloadListContentUrl( contents[index:], listId, resBusiness)
	go func() {
		for {
			select {
			case ts, ok := <-resBusiness:
				if ok {
					ans = append(ans,ts.BusinessList...)
				} else {
					fmt.Println("管道已经关闭，close")
					closeCh <- struct{}{}
					return
				}
			}
		}
	}()
	time.Sleep(5*time.Second)
	return ans
}

func downloadListContentUrl( contents []string, listId string, bus chan *ImgContentBusiness) {
	defer DeferPanicToLog()
	var busList []string
	for _, v := range contents {
		files, _ := ioutil.TempFile("", "tmpfile")
		_ = WgetFile(v, files.Name())
		//fmt.Println("filName:", files.Name())
		base64Num, _ := ImageFileBase64(files.Name())
		if base64Num != "" {
			busList = append(busList, GetStringMd5(base64Num+listId))
		}
		_ = os.Remove(files.Name())
	}
	bus <- &ImgContentBusiness{
		BusinessList: busList,
	}
}

func WgetFile(url string, tmpPath string) (err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	f, err := os.Create(tmpPath)
	defer f.Close()
	if err != nil {
		return
	}
	io.Copy(f, res.Body)
	return
}
func DeferPanicToLog(opts ...interface{}) {
	if err := recover(); err != nil {
		fmt.Println(err, string(debug.Stack()))
	}
}

func mergeBusList(ch chan *ImgContentBusiness, closeCh chan<- struct{}, business *ImgContentBusiness) {
	defer DeferPanicToLog()
	for {
		select {
		case ts, ok := <-ch:
			if ok {
				business.BusinessList = append(business.BusinessList, ts.BusinessList...)
			} else {
				fmt.Println("管道已经关闭，close")
				closeCh <- struct{}{}
				return
			}
		}
	}
}

func DownloadImgUrl(fileURL string) (fileName string) {
	// 发起HTTP GET请求获取文件
	response, err := http.Get(fileURL)
	if err != nil {
		fmt.Println("HTTP GET请求失败:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("HTTP请求失败，状态码: ", response.StatusCode)
		return
	}

	// 创建本地文件保存下载的文件
	fmt.Println("fileUrl:", getFileNameFromURL(fileURL))
	fileName = "/Users/zhaodeman/" + GetUUID() + getFileNameFromURL(fileURL)
	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("创建本地文件失败:", err)
		return
	}
	defer outputFile.Close()

	// 将HTTP响应的文件数据保存到本地文件
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		fmt.Println("保存文件失败:", err.Error())
		return
	}
	return
}

// 从URL中提取文件名
func getFileNameFromURL(url string) string {
	parts := strings.Split(url, "/")
	return parts[len(parts)-1]
}

func GetUUID() string {
	b := make([]byte, 48)
	// var b [48]byte
	if _, err := io.ReadFull(crand.Reader, b); err != nil {
		return ""
	}
	return GetStringMd5(base64.URLEncoding.EncodeToString(b))
}

func GetStringMd5(passwd string) string {
	passwddata := []byte(passwd)
	has := md5.Sum(passwddata)
	md5pwd := fmt.Sprintf("%x", has) // 将[]byte转成16进制
	return md5pwd
}

func ImageFileBase64(path string) (imgbase64 string, err error) {
	file, err := os.Open(path)
	defer file.Close()
	imgByte, _ := ioutil.ReadAll(file)
	//mimeType := http.DetectContentType(imgByte)
	imgbase64 = base64.StdEncoding.EncodeToString(imgByte)
	return
}
