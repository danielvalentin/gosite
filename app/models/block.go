package models

import (
	"fmt"
	"html/template"
)

type Block struct {
	Base
	Id						int64
	Content_id				int64
	Blocktype_id			int64
	Parent					int64
	Value					string
	Collapsed				int64
	Order					int64
}

func (block *Block) Render() template.HTML {
	return template.HTML(block.Value)
}

func (block *Block) GetByContentID(contentID int64) []Block {
	
	dbmap := block.InitDb()
	defer dbmap.Db.Close()
	
	var blocks []Block
	_, err := dbmap.Select(&blocks, "SELECT * FROM `blocks` WHERE `content_id` = ?", contentID)
	
	if err != nil {
		fmt.Printf("Error selecting blocks. Err: %s",err.Error())
		panic("Argh")
		return nil
	}
	
	return blocks
	
}
