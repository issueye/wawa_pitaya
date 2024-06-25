package frm_logger

import (
	"fmt"
	"pitaya/internal/config"
	"strings"

	"github.com/ying32/govcl/vcl"
)

// ::private::
type TFrm_logger_settingsFields struct {
}

func (f *TFrm_logger_settings) OnFormCreate(sender vcl.IObject) {
	f.SetValue()

	// 事件
	f.Btn_cancel.SetOnClick(f.OnCancelClick)
	f.Btn_save.SetOnClick(f.OnSaveClick)
}

func (f *TFrm_logger_settings) SetValue() {
	path := config.GetParam(config.CfgLogPath, "").String()
	path = strings.TrimPrefix(path, "runtime/")
	f.Edt_log_path.SetText(path)
	f.Cmb_level.SetText(config.GetParam(config.CfgLogLevel, "").String())
	f.Edt_age.SetValue(int32(config.GetParam(config.CfgLogMaxAge, "30").Int()))
	f.Edt_backups.SetValue(int32(config.GetParam(config.CfgLogMaxBackups, "30").Int()))
	f.Ckb_compress.SetChecked(config.GetParam(config.CfgLogCompress, "true").Bool())
}

func (f *TFrm_logger_settings) GetValue() {
	path := fmt.Sprintf("%s/%s", f.Edt_runtime.Text(), f.Edt_log_path.Text())
	config.SetParam(config.CfgLogPath, path, "日志存放路径")
	config.SetParam(config.CfgLogLevel, f.Cmb_level.Text(), "日志等级")
	config.SetParam(config.CfgLogMaxAge, f.Edt_age.Text(), "保存天数")
	config.SetParam(config.CfgLogMaxBackups, f.Edt_backups.Text(), "日志备份数")

	if f.Ckb_compress.Checked() {
		config.SetParam(config.CfgLogCompress, "true", "是否压缩")
	} else {
		config.SetParam(config.CfgLogCompress, "false", "是否压缩")
	}
}

func (f *TFrm_logger_settings) OnCancelClick(sender vcl.IObject) {
	// 事件
	f.Close()
}

func (f *TFrm_logger_settings) OnSaveClick(sender vcl.IObject) {
	f.GetValue()
	// 事件
	f.Close()
}
