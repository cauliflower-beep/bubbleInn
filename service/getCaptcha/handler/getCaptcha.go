package handler

import (
	"context"
	"encoding/json"
	"getCaptcha/model"
	"image/color"

	//log "github.com/micro/go-micro/v2/logger"
	"github.com/afocus/captcha"
	getCaptcha "getCaptcha/proto/getCaptcha"
)

type GetCaptcha struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetCaptcha) Call(ctx context.Context, req *getCaptcha.Request, rsp *getCaptcha.Response) error {
	cap := captcha.New()

	//设置字体
	cap.SetFont("./conf/comic.ttf")

	//设置验证码大小
	cap.SetSize(128,64)

	//设置干扰强度
	cap.SetDisturbance(captcha.NORMAL)

	//设置前景色
	cap.SetFrontColor(color.RGBA{128,255,0,128})

	//设置背景色
	cap.SetBkgColor(color.RGBA{100,5,255,128},color.RGBA{255,123,255,255})
	//生成字体
	//产生的str字符串，是作为传入参数的
	img,str := cap.Create(4,captcha.NUM)

	//存储图片验证码到redis中
	err := model.SaveImgCode(str,req.Uuid)
	if err != nil{
		return err
	}

	//将生成的图片验证码序列化
	imgBuf,err := json.Marshal(img)
	if err != nil{
		return err
	}
	//将imgBuf使用参数rsp传出
	rsp.Img = imgBuf
	return nil
}

