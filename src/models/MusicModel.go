package models

type Music struct {
	MusicId  int64 `gorm:"primaryKey"`
	Title    string
	Artist   string
	Album    string
	Duration float64
	Path     string
	Likes    int64
}

type MusicRequestVO struct {
	Title  string
	Artist string
	Album  string
}
