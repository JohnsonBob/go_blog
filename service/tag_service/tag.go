package tag_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"go_blog/models"
	"go_blog/pkg/excel"
	"go_blog/pkg/file"
	"go_blog/pkg/gredis"
	"go_blog/pkg/util"
	"go_blog/service/cache_service"
	"mime/multipart"
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
	index := xlsFile.NewSheet(excel.Sheet)
	xlsFile.SetActiveSheet(index)

	err = xlsFile.SetCellValue(excel.Sheet, "A1", "ID")
	err = xlsFile.SetCellValue(excel.Sheet, "B1", "名称")
	err = xlsFile.SetCellValue(excel.Sheet, "C1", "创建人")
	err = xlsFile.SetCellValue(excel.Sheet, "D1", "创建时间")
	err = xlsFile.SetCellValue(excel.Sheet, "E1", "修改人")
	err = xlsFile.SetCellValue(excel.Sheet, "F1", "修改时间")

	for index, value := range *tags {
		line := 2 + index
		err = xlsFile.SetCellValue(excel.Sheet, fmt.Sprintf("%s%d", "A", line), value.ID)
		err = xlsFile.SetCellValue(excel.Sheet, fmt.Sprintf("%s%d", "B", line), value.Name)
		err = xlsFile.SetCellValue(excel.Sheet, fmt.Sprintf("%s%d", "C", line), value.CreatedBy)
		err = xlsFile.SetCellValue(excel.Sheet, fmt.Sprintf("%s%d", "D", line), value.CreatedOn)
		err = xlsFile.SetCellValue(excel.Sheet, fmt.Sprintf("%s%d", "E", line), value.ModifiedBy)
		err = xlsFile.SetCellValue(excel.Sheet, fmt.Sprintf("%s%d", "F", line), value.ModifiedOn)
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

func ImportTag(file multipart.File) error {
	defer func() { _ = file.Close() }()
	xlsFile, err := excelize.OpenReader(file)
	if err != nil {
		return err
	}
	rows, err := xlsFile.GetRows(excel.Sheet)
	if err != nil {
		return err
	}
	for index, value := range rows {
		if index > 0 {
			if len(value) < 5 {
				return errors.New("导入模板错误")
			}
			models.AddTag(&models.Tag{Name: value[1], CreatedBy: value[2], ModifiedBy: value[4], State: 1})
		}
	}
	return nil
}
