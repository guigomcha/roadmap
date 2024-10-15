package models

type DefaultData struct {
	Updated int64 `gorm:"autoUpdateTime:milli"` // Use unix milli seconds as updating time
	Created int64 `gorm:"autoCreateTime"`       // Use unix seconds as creating time
}

type Spot struct {
	DefaultData
	ID          uint        `gorm:"primaryKey"`
	Location    GeoLocation `gorm:"unique,embedded"`
	Description string
	Name        string
	Attachments []string // foreign key??
	Tasks       []string // foreign key??

}

type GeoLocation struct {
	Long float64
	Lat  float64
}

type Attachment struct {
	DefaultData
	ID   uint `gorm:"primaryKey"`
	Name string
	Url  string
}

type Task struct {
	DefaultData
	ID uint `gorm:"primaryKey"`
}

// https://gorm.io/docs/associations.html#tags
