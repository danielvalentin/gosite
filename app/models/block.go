package models

import (
	"html/template"
)

type Block struct {
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
	conn := getDB()
	defer conn.Db.Close()
	var blocks []Block
	_, err := conn.Select(&blocks, "SELECT * FROM `blocks` WHERE `content_id` = ?", contentID)
	if err != nil {
		panic(err.Error())
	}
	return blocks
}
