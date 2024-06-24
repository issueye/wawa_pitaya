package about

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

// ::private::
type TFrmAboutFields struct {
}

func (f *TFrmAbout) OnFormCreate(sender vcl.IObject) {
	f.TForm.SetPosition(types.PoScreenCenter)
}
