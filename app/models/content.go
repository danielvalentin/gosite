package models

type Content struct {
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

func (content Content) GetAll() []Content {
	conn := getDB()
	defer conn.Db.Close()
	
	var contents []Content
	_, err := conn.Select(&contents, "SELECT * FROM `contents`")
	if err != nil {
		panic(err.Error())
	}
	
	return contents
	
}

func (content Content) GetByGuid(slug string) Content {
	conn := getDB()
	defer conn.Db.Close()
	var cont Content
	err := conn.SelectOne(&cont, "SELECT * FROM contents WHERE `guid` = ?", slug)
	if err != nil {
		print("Slug was: "+slug)
		panic(err.Error())
	}
	cont.Blocks = new(Block).GetByContentID(cont.Id)
	return cont
}

