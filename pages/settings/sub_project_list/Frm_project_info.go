// Automatically generated by the res2go, do not edit.

package sub_project_list

import (
	"github.com/ying32/govcl/vcl"
)

type TFrm_project_info struct {
	*vcl.TForm
	Edt_topic       *vcl.TLabeledEdit
	Btn_add         *vcl.TBitBtn
	Lv_data_table   *vcl.TListView
	Mmo_description *vcl.TMemo
	Lbl_description *vcl.TLabel

	// ::private::
	TFrm_project_infoFields
}

var Frm_project_info *TFrm_project_info

// Loaded in bytes.
// vcl.Application.CreateForm(&Frm_project_info)

func NewFrm_project_info(owner vcl.IComponent) (root *TFrm_project_info) {
	vcl.CreateResForm(owner, &root)
	return
}

var Frm_project_infoBytes = []byte("\x54\x50\x46\x30\x0B\x54\x44\x65\x73\x69\x67\x6E\x46\x6F\x72\x6D\x10\x46\x72\x6D\x5F\x70\x72\x6F\x6A\x65\x63\x74\x5F\x69\x6E\x66\x6F\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x03\xE9\x01\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\x7E\x01\x0B\x42\x6F\x72\x64\x65\x72\x53\x74\x79\x6C\x65\x07\x08\x62\x73\x44\x69\x61\x6C\x6F\x67\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0C\xE8\xAE\xA2\xE9\x98\x85\xE4\xBF\xA1\xE6\x81\xAF\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xE9\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x7E\x01\x00\x0C\x54\x4C\x61\x62\x65\x6C\x65\x64\x45\x64\x69\x74\x09\x45\x64\x74\x5F\x74\x6F\x70\x69\x63\x04\x4C\x65\x66\x74\x02\x25\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x10\x05\x57\x69\x64\x74\x68\x03\xFC\x00\x10\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x48\x65\x69\x67\x68\x74\x02\x11\x0F\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x57\x69\x64\x74\x68\x02\x18\x11\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE4\xB8\xBB\xE9\xA2\x98\x15\x45\x64\x69\x74\x4C\x61\x62\x65\x6C\x2E\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x0D\x4C\x61\x62\x65\x6C\x50\x6F\x73\x69\x74\x69\x6F\x6E\x07\x06\x6C\x70\x4C\x65\x66\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x69\x74\x42\x74\x6E\x07\x42\x74\x6E\x5F\x61\x64\x64\x04\x4C\x65\x66\x74\x03\x29\x01\x06\x48\x65\x69\x67\x68\x74\x02\x1E\x03\x54\x6F\x70\x02\x0E\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\xB7\xBB\xE5\x8A\xA0\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x09\x54\x4C\x69\x73\x74\x56\x69\x65\x77\x0D\x4C\x76\x5F\x64\x61\x74\x61\x5F\x74\x61\x62\x6C\x65\x04\x4C\x65\x66\x74\x02\x0A\x06\x48\x65\x69\x67\x68\x74\x03\x5D\x01\x03\x54\x6F\x70\x03\x82\x00\x05\x57\x69\x64\x74\x68\x03\x69\x01\x07\x43\x6F\x6C\x75\x6D\x6E\x73\x0E\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE4\xB8\xBB\xE9\xA2\x98\x05\x57\x69\x64\x74\x68\x03\x96\x00\x00\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\x8F\x8F\xE8\xBF\xB0\x05\x57\x69\x64\x74\x68\x03\x96\x00\x00\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\x93\x8D\xE4\xBD\x9C\x00\x00\x09\x47\x72\x69\x64\x4C\x69\x6E\x65\x73\x09\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x09\x56\x69\x65\x77\x53\x74\x79\x6C\x65\x07\x08\x76\x73\x52\x65\x70\x6F\x72\x74\x00\x00\x05\x54\x4D\x65\x6D\x6F\x0F\x4D\x6D\x6F\x5F\x64\x65\x73\x63\x72\x69\x70\x74\x69\x6F\x6E\x04\x4C\x65\x66\x74\x02\x26\x06\x48\x65\x69\x67\x68\x74\x02\x3B\x03\x54\x6F\x70\x02\x38\x05\x57\x69\x64\x74\x68\x03\xFA\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x06\x54\x4C\x61\x62\x65\x6C\x0F\x4C\x62\x6C\x5F\x64\x65\x73\x63\x72\x69\x70\x74\x69\x6F\x6E\x04\x4C\x65\x66\x74\x02\x0B\x06\x48\x65\x69\x67\x68\x74\x02\x11\x03\x54\x6F\x70\x02\x39\x05\x57\x69\x64\x74\x68\x02\x18\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\xE6\x8F\x8F\xE8\xBF\xB0\x0B\x50\x61\x72\x65\x6E\x74\x43\x6F\x6C\x6F\x72\x08\x00\x00\x00")

// 注册窗口资源
var _ = vcl.RegisterFormResource(Frm_project_info, &Frm_project_infoBytes)
