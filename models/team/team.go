package team

type Model struct {
	ID           int    `gorm:"column:id;primaryKey;autoIncrement"`
	Abbreviation string `gorm:"column:abbreviation;type:varchar(10);not null;unique"`
	City         string `gorm:"column:city;type:varchar(50);not null"`
	Conference   string `gorm:"column:conference;type:varchar(20);not null"`
	Division     string `gorm:"column:division;type:varchar(20);not null"`
	FullName     string `gorm:"column:full_name;type:varchar(100);not null;unique"`
	Name         string `gorm:"column:name;type:varchar(50);not null"`
}
