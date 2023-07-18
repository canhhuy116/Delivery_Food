package uploadmodel

type Upload struct {
	// gorm.Model
	ID   int    `json:"id" gorm:"column:id;"`
	Url  string `json:"url" gorm:"column:url;"`
	Name string `json:"name" gorm:"column:name;"`
}

func (Upload) TableName() string {
	return "uploads"
}

type UploadCreate struct {
	// gorm.Model
	ID   int    `json:"id" gorm:"column:id;"`
	Url  string `json:"url" gorm:"column:url;"`
	Name string `json:"name" gorm:"column:name;"`
}

func (*UploadCreate) TableName() string {
	return Upload{}.TableName()
}

type UploadUpdate struct {
	// gorm.Model
	ID   int    `json:"id" gorm:"column:id;"`
	Url  string `json:"url" gorm:"column:url;"`
	Name string `json:"name" gorm:"column:name;"`
}

func (*UploadUpdate) TableName() string {
	return Upload{}.TableName()
}

func ErrFileIsNotImage(err error) error {
	return err
}

func ErrCannotSaveFile(err error) error {
	return err
}
