package storage

type IndicatorStorage struct {
}

func NewIndicatorStorage() Indicatorer {
	return &IndicatorStorage{}
}

// Задел на будущее. Необходимо будет добавить адаптерс
