package repositories

type EdgeRepository interface {
	Save(source int64, target int64) (int64, error)
}
