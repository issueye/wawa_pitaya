package settings

import (
	"github.com/ying32/govcl/vcl"
)

// ::private::
type TFrmPitayaSettingsFields struct {
}

func (f *TFrmPitayaSettings) OnFormCreate(sender vcl.IObject) {
	f.EdtShowLogCount.SetMinValue(20)

	// 事件
	f.BtnCancel.SetOnClick(f.OnCancelClick)
	f.BtnSave.SetOnClick(f.OnSaveClick)
	f.EdtShowLogCount.SetOnChange(f.OnEdtShowLogCountChange)
}

func (f *TFrmPitayaSettings) OnCancelClick(sender vcl.IObject) {
	// 事件
	f.Close()
}

func (f *TFrmPitayaSettings) OnSaveClick(sender vcl.IObject) {
	// 事件

}

func (f *TFrmPitayaSettings) OnEdtShowLogCountChange(sender vcl.IObject) {
	minValue := f.EdtShowLogCount.MinValue()
	// 事件
	if f.EdtShowLogCount.Value() < minValue {
		f.EdtShowLogCount.SetValue(minValue)
	}
}
