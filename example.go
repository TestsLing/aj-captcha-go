package main

import (
	"encoding/json"
	"errors"
	config2 "github.com/TestsLing/aj-captcha-go/config"
	"github.com/TestsLing/aj-captcha-go/const"
	"github.com/TestsLing/aj-captcha-go/service"
	"io/ioutil"
	"log"
	"net/http"
)

type clientParams struct {
	Token       string `json:"token"`
	PointJson   string `json:"pointJson"`
	CaptchaType string `json:"captchaType"`
}

var config = config2.NewConfig()
var factory = service.NewCaptchaServiceFactory(config)

func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")                                                                      // 可将将 * 替换为指定的域名
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization,x-requested-with") //你想放行的header也可以在后面自行添加
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")                                                    //我自己只使用 get post 所以只放行它
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// 放行所有OPTIONS方法
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		// 处理请求
		f(w, r)
	}
}

func getCaptcha(writer http.ResponseWriter, request *http.Request) {
	params, err := getParams(request)
	if err != nil {
		res, _ := json.Marshal(errorRes(err))
		writer.Write(res)
		return
	}
	if params.CaptchaType == "" {
		res, _ := json.Marshal(errorRes(errors.New("参数传递不完整")))
		writer.Write(res)
		return
	}

	ser := factory.GetService(params.CaptchaType)

	data := ser.Get()

	res, _ := json.Marshal(successRes(data))
	request.Body.Close()
	writer.Write(res)
}

func main() {

	factory.RegisterCache(constant.MemCacheKey, service.NewMemCacheService(20))
	factory.RegisterService(constant.ClickWordCaptcha, service.NewClickWordCaptchaService(factory))
	factory.RegisterService(constant.BlockPuzzleCaptcha, service.NewBlockPuzzleCaptchaService(factory))

	http.HandleFunc("/captcha/get", cors(getCaptcha))
	http.HandleFunc("/captcha/check", cors(checkCaptcha))

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func checkCaptcha(writer http.ResponseWriter, request *http.Request) {
	params, err := getParams(request)

	if params.Token == "" || params.PointJson == "" || params.CaptchaType == "" {
		res, _ := json.Marshal(errorRes(errors.New("参数传递不完整")))
		writer.Write(res)
		return
	}

	if err != nil {
		res, _ := json.Marshal(errorRes(err))
		writer.Write(res)
		return
	}

	ser := factory.GetService(params.CaptchaType)

	err = ser.Check(params.Token, params.PointJson)
	if err != nil {
		res, _ := json.Marshal(errorRes(err))
		writer.Write(res)
		return
	}

	res, _ := json.Marshal(successRes(nil))
	writer.Write(res)
}

func successRes(data interface{}) map[string]interface{} {
	ret := make(map[string]interface{})
	ret["error"] = false
	ret["repCode"] = "0000"
	ret["repData"] = data
	ret["repMsg"] = nil
	ret["successRes"] = true

	return ret
}

func getParams(request *http.Request) (*clientParams, error) {
	params := &clientParams{}
	all, _ := ioutil.ReadAll(request.Body)

	if len(all) <= 0 {
		return nil, errors.New("未获取到客户端数据")
	}

	err := json.Unmarshal(all, params)
	if err != nil {
		return nil, err
	}

	return params, nil
}

func errorRes(err error) map[string]interface{} {
	ret := make(map[string]interface{})
	ret["error"] = true
	ret["repCode"] = "0001"
	ret["repData"] = nil
	ret["repMsg"] = err.Error()
	ret["successRes"] = false
	return ret
}
