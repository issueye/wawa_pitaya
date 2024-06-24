package home

import (
	"fmt"
	"pitaya/pages/about"
	"pitaya/pages/message"
	logger_settings "pitaya/pages/settings/logger"
	mqtt_settings "pitaya/pages/settings/mqtt"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

// ::private::
type TFrmPitayaFields struct {
	IsServerRun  bool
	ShowLogCount int32
	IsTrueClose  bool
	msg          *message.TFrmMessage
}

func (f *TFrmPitaya) OnFormCreate(sender vcl.IObject) {
	f.LblServerName.SetCaption("服务名称 火龙果后台服务")
	f.IsServerRun = false
	// 后续修改为读取配置文件
	f.ShowLogCount = 50
	f.MenuItemServerSettings.SetCaption("日志设置")

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
	f.MenuItemServerSettings.SetOnClick(f.OnCronClick)
	f.Menu_mqtt_settings.SetOnClick(f.OnMqttSettingsClick)
	f.BtnServerRun.SetOnClick(f.OnRunServerClick)
	f.BtnClearLog.SetOnClick(f.OnClearLogOnClick)
	f.TForm.SetOnClose(f.OnFormClose)
	f.PopMClose.SetOnClick(f.OnAppCloseClick)
	f.PopMShow.SetOnClick(f.OnAppShowClick)
}

func (f *TFrmPitaya) OnFormClose(sender vcl.IObject, action *types.TCloseAction) {
	f.Hide()
	if !f.IsTrueClose {
		*action = types.CaHide
	}
}

func (f *TFrmPitaya) OnMqttSettingsClick(sender vcl.IObject) {
	mqtt_settings.NewFrm_mqtt_settings(f).ShowModal()
}

func (f *TFrmPitaya) OnAppCloseClick(sender vcl.IObject) {
	f.IsTrueClose = true
	f.Close()
}

func (f *TFrmPitaya) OnAppShowClick(sender vcl.IObject) {
	f.Show()
}

func (f *TFrmPitaya) OnCronClick(sender vcl.IObject) {
	pitaya := logger_settings.NewFrm_logger_settings(nil)
	pitaya.SetPosition(types.PoMainFormCenter)
	pitaya.ShowModal()
}

func (f *TFrmPitaya) OnAboutClick(sender vcl.IObject) {
	about := about.NewFrmAbout(nil)
	about.SetPosition(types.PoMainFormCenter)
	about.ShowModal()
}

func (f *TFrmPitaya) OnRunServerClick(sender vcl.IObject) {
	panels := f.StatusBar.Panels()
	f.TFrmPitayaFields.IsServerRun = !f.TFrmPitayaFields.IsServerRun
	if f.TFrmPitayaFields.IsServerRun {
		f.BtnServerRun.SetCaption("开启服务")
		panels.Items(1).SetText("关闭")
		f.addLog("关闭服务")
	} else {
		f.BtnServerRun.SetCaption("关闭服务")

		now := time.Now().Format("2006-01-02 15:04:05")
		f.LblServerStartTime.SetCaption(fmt.Sprintf("开启时间 %s", now))
		panels.Items(1).SetText("开启")
		f.addLog("开启服务")
	}

	f.msg.SetMessage("测试内容")
	f.msg.Show()

	go func() {
		time.Sleep(time.Second * 5)
		f.msg.Hide()
	}()
}

func (f *TFrmPitaya) OnTimer(sender vcl.IObject) {
	item := f.StatusBar.Panels().Items(3)
	now := time.Now().Format("2006-01-02 15:04:05")
	item.SetText(now)
}

func (f *TFrmPitaya) OnClearLogOnClick(sender vcl.IObject) {
	f.MmoRunLog.Lines().Clear()
}

// 添加日志到主窗口中
func (f *TFrmPitaya) addLog(msg string) {
	now := time.Now().Format("2006-01-02 15:04:05")
	f.MmoRunLog.Lines().Add(fmt.Sprintf("%s %s", msg, now))

	if f.MmoRunLog.Lines().Count() > f.ShowLogCount {
		f.MmoRunLog.Lines().Delete(0)
	}
}
