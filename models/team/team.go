package team

type Model struct {
	ID           string `gorm:"column:id;primaryKey;autoIncrement"`
	Abbreviation string `gorm:"column:abbreviation;type:varchar(10)"`
	City         string `gorm:"column:city;type:varchar(50)"`
	Conference   string `gorm:"column:conference;type:varchar(20)"`
	Division     string `gorm:"column:division;type:varchar(20)"`
	FullName     string `gorm:"column:full_name;type:varchar(100)"`
	Name         string `gorm:"column:name;type:varchar(50)"`
}
