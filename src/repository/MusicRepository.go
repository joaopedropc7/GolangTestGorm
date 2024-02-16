package repository

import (
	"Routes/src/models"
	"errors"
	"github.com/tcolgate/mp3"
	"gorm.io/gorm"
	"io"
	"os"
)

type MusicRepository struct {
	DB *gorm.DB
}

func NewMusicRepository(db *gorm.DB) *MusicRepository {
	return &MusicRepository{db}
}

func (r *MusicRepository) CreateMusic(music *models.Music, fileData []byte) (*models.Music, error) {

	err := r.saveFile(music, fileData)
	if err != nil {
		return nil, err
	}

	if err := r.DB.Create(music).Error; err != nil {
		return nil, err
	}

	fileInDb, err := os.Open(music.Path)
	if err != nil {
		return nil, err
	}
	defer fileInDb.Close()

	duration, err := getMP3Duration(fileInDb)
	if err != nil {
		return nil, err
	}

	music.Duration = duration

	if err := r.DB.Model(&models.Music{}).Where("music_id = ?", music.MusicId).Update("duration", duration).Error; err != nil {
		return nil, err
	}

	return music, nil
}

// saveFile salva o arquivo da música no sistema de arquivos
func (r *MusicRepository) saveFile(music *models.Music, fileData []byte) error {
	uploadPath := "./uploads/"
	filePath := uploadPath + music.Title + ".mp3"

	_, err := os.Stat(filePath)
	if err == nil {
		// O arquivo já existe, retorne um erro ou tome outra ação apropriada
		return errors.New("Já existe um arquivo com o mesmo título")
	} else if !os.IsNotExist(err) {
		// Outro erro ocorreu ao verificar a existência do arquivo
		return err
	}

	errMkDir := os.MkdirAll(uploadPath, os.ModePerm)
	if errMkDir != nil {
		return errMkDir
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(fileData)
	if err != nil {
		return err
	}

	music.Path = filePath
	return nil
}

func (r *MusicRepository) GetMusicById(musicId int64) (models.Music, error) {
	var music models.Music
	if err := r.DB.First(&music, musicId).Error; err != nil {
		return models.Music{}, errors.New("não foi encontrado nenhum registro com este ID")
	}
	return music, nil
}

func getMP3Duration(file *os.File) (float64, error) {

	d := mp3.NewDecoder(file)
	var f mp3.Frame
	skipped := 0

	t := 0.0

	for {

		if err := d.Decode(&f, &skipped); err != nil {
			if err == io.EOF {
				break
			}

			return 0, err
		}

		t = t + f.Duration().Seconds()
	}

	return t, nil
}
