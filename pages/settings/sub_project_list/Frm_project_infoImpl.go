package sub_project_list

import (
	"fmt"
	"pitaya/internal/global"
	"pitaya/internal/model"
	"pitaya/internal/service"
	"unsafe"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

// ::private::
type TFrm_project_infoFields struct {
	srv *service.ProjectList

	tableData []*model.ProjectList
	fresh     chan struct{}

	ButtonList []*vcl.TButton
}

func (f *TFrm_project_info) OnFormCreate(sender vcl.IObject) {
	f.srv = &service.ProjectList{}
	f.tableData = make([]*model.ProjectList, 0)
	f.ButtonList = make([]*vcl.TButton, 0)
	f.Btn_add.SetOnClick(f.OnBtn_addClick)
	f.Lv_data_table.SetOnCustomDrawSubItem(f.OnLv_data_tableCustomDrawSubItem)
	f.GetData()
	f.fresh = make(chan struct{})
	f.Monitor()
}

// 添加按钮事件
func (f *TFrm_project_info) OnBtn_addClick(sender vcl.IObject) {
	// 添加数据到数据库

	if f.Edt_topic.Text() != "" {
		err := f.srv.Add(&model.ProjectList{
			Topic:       f.Edt_topic.Text(),
			Description: f.Mmo_description.Text(),
		})

		if err != nil {
			global.Logger.Errorf("添加数据到数据库失败：%v", err)
			vcl.ShowMessage("添加数据到数据库失败：" + err.Error())
			return
		}

		f.Edt_topic.Clear()
		f.Mmo_description.Clear()

		f.GetData()
	}
}

// 列表绘制事件
func (f *TFrm_project_info) OnLv_data_tableCustomDrawSubItem(sender *vcl.TListView, item *vcl.TListItem, subItem int32, state types.TCustomDrawState, defaultDraw *bool) {
	// 如果是操作列
	if subItem == 2 {
		// 添加一个删除按钮
		if item.Data() == nil {
			*defaultDraw = false

			if int(item.Index()) > len(f.tableData) {
				return
			}

			btn := vcl.NewButton(sender)
			btn.SetParent(sender)
			btn.SetCaption("删除")
			btn.SetVisible(true)
			btn.SetTag(int(item.Index()))
			btn.SetLeft(item.Left() + 295)
			btn.SetTop(item.Top())
			btn.SetWidth(48)
			btn.SetHeight(20)
			btn.Font().SetSize(8)
			btn.SetOnClick(func(sender vcl.IObject) {
				fmt.Println("item.Index()", item.Index())

				data := f.tableData[int(item.Index())]
				err := f.srv.Delete(uint(data.Id))
				if err != nil {
					global.Logger.Errorf("删除数据失败：%v", err)
					return
				}

				f.fresh <- struct{}{}
			})
			item.SetData((unsafe.Pointer)(&btn))

			f.ButtonList = append(f.ButtonList, btn)
		}
	}
}

func (f *TFrm_project_info) Monitor() {
	go func() {
		for {
			select {
			case <-f.fresh:
				vcl.ThreadSync(func() {
					f.GetData()
					// vcl.ShowMessage("刷新成功")
				})
			}
		}
	}()
}

func (f *TFrm_project_info) GetData() {
	var err error
	f.tableData = make([]*model.ProjectList, 0)

	f.tableData, err = f.srv.List()
	if err != nil {
		global.Logger.Errorf("获取列表失败：%v", err)
		vcl.ShowMessage("获取列表失败：" + err.Error())
		return
	}

	fmt.Println("f.tableData", len(f.tableData))

	f.Lv_data_table.Items().BeginUpdate()
	f.Lv_data_table.Items().Clear()

	// 释放所有按钮
	for _, btn := range f.ButtonList {
		btn.SetVisible(false)
		btn.SetParent(nil)
		btn.Free()
	}

	f.ButtonList = make([]*vcl.TButton, 0)

	for _, data := range f.tableData {
		item := f.Lv_data_table.Items().Add()
		item.SetCaption(data.Topic)
		item.SubItems().Add(data.Description)
	}
	f.Lv_data_table.Items().EndUpdate()
}
