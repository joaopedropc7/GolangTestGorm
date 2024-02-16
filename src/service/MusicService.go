package service

import (
	"Routes/src/banco"
	"Routes/src/models"
	"Routes/src/repository"
)

func CreateMusic(music *models.MusicRequestVO, fileData []byte, usuario *models.Usuario) (*models.Music, error) {

	db, err := banco.Conectar()
	if err != nil {
		return nil, err
	}

	musicModel := &models.Music{
		Artist:  music.Artist,
		Album:   music.Album,
		Title:   music.Title,
		Usuario: usuario,
	}

	musicRepository := repository.NewMusicRepository(db)

	return musicRepository.CreateMusic(musicModel, fileData)
}

func FindMusicById(musicId int64) (*models.Music, error) {
	db, err := banco.Conectar()
	if err != nil {
		return nil, err
	}

	musicRepository := repository.NewMusicRepository(db)

	music, err := musicRepository.GetMusicById(musicId)
	if err != nil {
		return nil, err
	}

	return &music, nil
}
