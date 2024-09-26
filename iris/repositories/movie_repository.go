package repositories

import "imooc-iris/datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {

}

func NewMovieManager() MovieRepository  {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string  {
	movie:=&datamodels.Movie{Name:"xxxx"}
	return movie.Name
}
