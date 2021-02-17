package model

type User struct {
	UserId      int64	`json:"userId" gorm:"primary_key;not_null;auto_increment"`
	UserName    string	`json:"userName" gorm:"unique;not_null" from:"userName" binding:"required"`
	UserPwd     string	`json:"userPwd"`
	IsOk int64	`json:"isOk" gorm:"default 1"`
	UserAvatar  string	`json:"userAvatar"`
	UserRole string	`json:"userRole"`
}


type UserPathModel struct {
	PathId int64	`json:"pathId" gorm:"primary_key;not_null;auto_increment"`
	BaseUrl string	`json:"baseUrl" gorm:"unique;not_null"`
	BaseUrlName string	`json:"baseUrlName"`
	IsOk int64	`json:"isOk" gorm:"default 1"`
	BaseIcon string	`json:"baseIcon"`
	OneUrlList []OneUrlList	`json:"oneUrlList" gorm:"ForeignKey:UserPathId"`
}

type OneUrlList struct {
	UrlId int64	`json:"urlId" gorm:"primary_key;not_null;auto_increment"`
	UserPathId int64	`json:"userPathId"`
	OneUrl string	`json:"oneUrl"`
	IsOk int64	`json:"isOk" gorm:"default 1"`
	OneUrlName string	`json:"oneUrlName"`
	OneUrlIcon string	`json:"oneUrlIcon"`
}

type AdminPathModel struct {
	PathId int64	`json:"pathId" gorm:"primary_key;not_null;auto_increment"`
	BaseUrl string	`json:"baseUrl" gorm:"unique;not_null"`
	IsOk int64	`json:"isOk" gorm:"default 1"`
	BaseUrlName string	`json:"baseUrlName"`
	BaseIcon string	`json:"baseIcon"`
	AdminOne []AdminOneUrlList	`json:"adminOne" gorm:"foreignKey:PathId"`
}

type AdminOneUrlList struct {
	UrlId int64	`json:"urlId" gorm:"primary_key;not_null;auto_increment"`
	PathId int64	`json:"PathId"`
	IsOk int64	`json:"isOk" gorm:"default 1"`
	OneUrl string	`json:"oneUrl"`
	OneUrlName string	`json:"oneUrlName"`
	OneUrlIcon string	`json:"oneUrlIcon"`
}