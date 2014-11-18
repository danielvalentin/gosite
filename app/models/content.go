package models

import (
	"fmt"
)

type Content struct {
	Base
	Id						int64
	User_id					int64
	Contenttype_id			int64
	Contenttypetype_id		int64
	Parent					int64
	Splittest				int64
	Splittest_title			string
	Splittest_parent_hits	int64
	Title					string
	Slug					string
	Guid					string
	Image_id				int64
	Pagetitle				string
	Metadescription			string
	Metakeywords			string
	Canonical				string
	Og_title				string
	Og_description			string
	Og_type					string
	Og_image				int64
	Og_url					string
	Status					string
	Hits					int64
	Created					int64
	Published				int64
	Blocks					[]Block
}


func (content *Content) GetAll() []Content {
	
	dbmap := content.InitDb()
	defer dbmap.Db.Close()
	
	var contents []Content
	_, err := dbmap.Select(&contents, "SELECT * FROM contents")
	
	if err != nil {
		fmt.Printf("Error selecting. Err: %s",err.Error())
		panic("Argh")
		return nil
	}
	
	return contents
	
}

func (content *Content) GetByGuid(slug string) Content {
	
	dbmap := content.InitDb()
	defer dbmap.Db.Close()
	
	var cont Content
	err := dbmap.SelectOne(&cont, "SELECT * FROM contents WHERE `guid` = ?", slug)
	
	if err != nil {
		fmt.Printf("Error in GetBySlug: SelectOne contents. Err: %s",err.Error())
		panic("Argh")
	}
	
	cont.Blocks = new(Block).GetByContentID(cont.Id)
	
	return cont
	
}
