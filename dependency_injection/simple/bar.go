package simple

type BarReposiotry struct {
}

func NewBarReposiotry() *BarReposiotry {
	return &BarReposiotry{}
}

type BarService struct {
	*BarReposiotry
}

func NewBarService(barReposiotry *BarReposiotry) *BarService {
	return &BarService{BarReposiotry: barReposiotry}
}
