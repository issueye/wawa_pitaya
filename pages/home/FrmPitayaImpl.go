package home

import (
	"fmt"
	"pitaya/pages/about"
	"pitaya/pages/message"
	logger_settings "pitaya/pages/settings/logger"
	mqtt_settings "pitaya/pages/settings/mqtt"
	sub_project_list "pitaya/pages/settings/sub_project_list"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

// ::private::
type TFrm_pitayaFields struct {
	IsServerRun  bool
	ShowLogCount int32
	IsTrueClose  bool
	msg          *message.TFrmMessage
}

func (f *TFrm_pitaya) OnFormCreate(sender vcl.IObject) {
	f.LblServerName.SetCaption("服务名称 火龙果后台服务")
	f.IsServerRun = false
	// 后续修改为读取配置文件
	f.ShowLogCount = 50
	f.Menu_logger_settings.SetCaption("日志设置")

	f.TForm.SetPosition(types.PoScreenCenter)
	f.TrayIcon.SetVisible(true)

	f.msg = message.NewFrmMessage(nil)
	f.msg.Hide()
	f.msg.SetPosition(types.PoDesigned)
	f.msg.SetTop(10)
	sx, _ := robotgo.GetScreenSize()
	f.msg.SetLeft(int32(sx) - f.msg.Width() - 10)
	fmt.Println("left", f.msg.Left())

	f.Timer.SetOnTimer(f.OnTimer)
	f.MenuItemAbout.SetOnClick(f.OnAboutClick)
	f.Menu_logger_settings.SetOnClick(f.OnCronClick)
	f.Menu_mqtt_settings.SetOnClick(f.OnMqttSettingsClick)
	f.Btn_server_run.SetOnClick(f.OnRunServerClick)
	f.Menu_project_list.SetOnClick(f.OnProjectListClick)
	f.TForm.SetOnClose(f.OnFormClose)
	f.PopMClose.SetOnClick(f.OnAppCloseClick)
	f.PopMShow.SetOnClick(f.OnAppShowClick)
}

func (f *TFrm_pitaya) OnFormClose(sender vcl.IObject, action *types.TCloseAction) {
	f.Hide()
	if !f.IsTrueClose {
		*action = types.CaHide
	}
}

func (f *TFrm_pitaya) OnMqttSettingsClick(sender vcl.IObject) {
	mqtt_settings.NewFrm_mqtt_settings(f).ShowModal()
}

func (f *TFrm_pitaya) OnAppCloseClick(sender vcl.IObject) {
	f.IsTrueClose = true
	f.Close()
}

func (f *TFrm_pitaya) OnAppShowClick(sender vcl.IObject) {
	f.Show()
}

func (f *TFrm_pitaya) OnCronClick(sender vcl.IObject) {
	pitaya := logger_settings.NewFrm_logger_settings(nil)
	pitaya.SetPosition(types.PoMainFormCenter)
	pitaya.ShowModal()
}

func (f *TFrm_pitaya) OnAboutClick(sender vcl.IObject) {
	about := about.NewFrmAbout(nil)
	about.SetPosition(types.PoMainFormCenter)
	about.ShowModal()
}

func (f *TFrm_pitaya) OnProjectListClick(sender vcl.IObject) {
	proList := sub_project_list.NewFrm_project_info(nil)
	proList.SetPosition(types.PoMainFormCenter)
	proList.ShowModal()
}

func (f *TFrm_pitaya) OnRunServerClick(sender vcl.IObject) {
	panels := f.StatusBar.Panels()
	f.TFrm_pitayaFields.IsServerRun = !f.TFrm_pitayaFields.IsServerRun
	if f.TFrm_pitayaFields.IsServerRun {
		f.Btn_server_run.SetCaption("开启")
		panels.Items(1).SetText("关闭")
	} else {
		f.Btn_server_run.SetCaption("关闭")
		panels.Items(1).SetText("开启")
	}

	f.msg.SetMessage("测试内容")
	f.msg.Show()

	go func() {
		time.Sleep(time.Second * 5)
		f.msg.Hide()
	}()
}

func (f *TFrm_pitaya) OnTimer(sender vcl.IObject) {
	item := f.StatusBar.Panels().Items(3)
	now := time.Now().Format("2006-01-02 15:04:05")
	item.SetText(now)
}
