package tag_service

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"go_blog/models"
	"go_blog/pkg/excel"
	"go_blog/pkg/file"
	"go_blog/pkg/gredis"
	"go_blog/pkg/util"
	"go_blog/service/cache_service"
	"strconv"
	"time"
)

func GetTags(tag *cache_service.Tag) (tags *[]models.Tag, err error) {
	key := tag.GetTagsKey()
	if gredis.IsExists(key) {
		var data []byte
		data, err = gredis.Get(key)
		if err != nil {
			return
		}
		err = json.Unmarshal(data, &tags)
		return
	}
	maps := make(map[string]interface{})
	if tag.Name != nil {
		maps["name"] = tag.Name
	}
	if tag.State != nil {
		maps["state"] = tag.State
	}
	all := models.GetTags(*tag.PageNum, *tag.PageSize, maps)
	tags = &all
	err = gredis.Set(key, tags, 1200)
	return
}

func Export(tag *cache_service.Tag) (fileName string, err error) {
	var tags *[]models.Tag
	tags, err = GetTags(tag)
	if err != nil {
		return
	}

	xlsFile := excelize.NewFile()
	// Create a new sheet.
	index := xlsFile.NewSheet("标签信息")
	xlsFile.SetActiveSheet(index)

	err = xlsFile.SetCellValue("标签信息", "A1", "ID")
	err = xlsFile.SetCellValue("标签信息", "B1", "名称")
	err = xlsFile.SetCellValue("标签信息", "C1", "创建人")
	err = xlsFile.SetCellValue("标签信息", "D1", "创建时间")
	err = xlsFile.SetCellValue("标签信息", "E1", "修改人")
	err = xlsFile.SetCellValue("标签信息", "F1", "修改时间")

	for index, value := range *tags {
		line := 2 + index
		err = xlsFile.SetCellValue("标签信息", fmt.Sprintf("%s%d", "A", line), value.ID)
		err = xlsFile.SetCellValue("标签信息", fmt.Sprintf("%s%d", "B", line), value.Name)
		err = xlsFile.SetCellValue("标签信息", fmt.Sprintf("%s%d", "C", line), value.CreatedBy)
		err = xlsFile.SetCellValue("标签信息", fmt.Sprintf("%s%d", "D", line), value.CreatedOn)
		err = xlsFile.SetCellValue("标签信息", fmt.Sprintf("%s%d", "E", line), value.ModifiedBy)
		err = xlsFile.SetCellValue("标签信息", fmt.Sprintf("%s%d", "F", line), value.ModifiedOn)
	}

	fullPath := excel.GetExcelFullPath()
	if err = file.IsNotExistMkDir(fullPath); err != nil {
		return
	}
	if isP := file.CheckPermission(fullPath); isP {
		return
	}

	timeNow := strconv.Itoa(int(time.Now().Unix()))
	fileName = "tags-" + timeNow + excel.EXT

	if err = xlsFile.SaveAs(fullPath + fileName); err != nil {
		util.Println(err)
	}
	return
}
