package message

import (
	"github.com/ying32/govcl/vcl"
)

// ::private::
type TFrmMessageFields struct {
}

func (f *TFrmMessage) OnFormCreate(sender vcl.IObject) {
	f.ImgClose.SetOnClick(f.OnImgCloseClick)

	f.Timer.SetEnabled(false)
	f.Timer.SetOnTimer(f.OnTimerTimer)
}

func (f *TFrmMessage) OnImgCloseClick(sender vcl.IObject) {
	f.Close()
}

func (f *TFrmMessage) OnTimerTimer(sender vcl.IObject) {
	f.Timer.SetEnabled(false)
	f.Close()
}

func (f *TFrmMessage) SetMessage(msg string) {
	f.LblMessage.SetCaption(msg)
}
