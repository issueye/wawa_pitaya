package mqtt_settings

import (
	"pitaya/internal/config"

	"github.com/ying32/govcl/vcl"
)

// ::private::
type TFrm_mqtt_settingsFields struct {
}

func (f *TFrm_mqtt_settings) OnFormCreate(sender vcl.IObject) {
	// 设置数据
	f.SetValue()

	f.Btn_save.SetOnClick(f.OnBtn_saveClick)
	f.Btn_cancel.SetOnClick(f.OnBtn_cancelClick)
}

func (f *TFrm_mqtt_settings) SetValue() {
	f.Edt_host.SetText(config.GetParam("mqtt_host", "").String())
	f.Edt_port.SetText(config.GetParam("mqtt_port", "").String())
	f.Edt_user.SetText(config.GetParam("mqtt_user", "").String())
	f.Edt_pwd.SetText(config.GetParam("mqtt_password", "").String())
}

func (f *TFrm_mqtt_settings) GetValue() {
	// 获取数据
	config.SetParam("mqtt_host", f.Edt_host.Text(), "服务地址")
	config.SetParam("mqtt_port", f.Edt_port.Text(), "服务端口")
	config.SetParam("mqtt_user", f.Edt_user.Text(), "用户名")
	config.SetParam("mqtt_password", f.Edt_pwd.Text(), "密码")
}

func (f *TFrm_mqtt_settings) OnBtn_saveClick(sender vcl.IObject) {
	f.GetValue()
	f.Close()
}

func (f *TFrm_mqtt_settings) OnBtn_cancelClick(sender vcl.IObject) {
	f.Close()
}
