package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

func main() {
	baidu()
}

func baidu() {
	var page int = 0
	for {
		url := `http://www.baidu.com/s?ie=utf-8&mod=1&isbd=1&isid=bdbbaf61002a79be&wd=%E8%B4%B4%E6%A0%87%E6%9C%BA&pn=` + strconv.Itoa(page*10) + `&oq=%E8%B4%B4%E6%A0%87%E6%9C%BA&tn=baiduhome_pg&ie=utf-8&usm=2&rsv_idx=2`
		html := HttpGet(url)
		//ioutil.WriteFile("./detail.html", []byte(html), 0666)
		reg := regexp.MustCompile(`<h3[\s\S]*?href="(.*?)"[\s\S]*?<\/h3>`)
		hrefs := reg.FindAllStringSubmatch(html, -1)
		for _, v := range hrefs {
			appendToFile("./baidu_list.txt", v[1]+"\n")
		}
		//os.Exit(1)
		page++
		//百度贴标机总共76页
		if page > 76 {
			os.Exit(1)
		}
	}

}

func appendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
		// 查找文件末尾的偏移量
		n, _ := f.Seek(0, os.SEEK_END)
		// 从末尾的偏移量开始写入内容
		_, err = f.WriteAt([]byte(content), n)
	}
	defer f.Close()
	return err
}

func HttpGet(url string) string {
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	response, err := client.Get(url)
	//response, err := http.Get(url)
	check(err)
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(len(string(body)))
	return string(body)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
