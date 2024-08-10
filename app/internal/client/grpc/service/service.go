package service

type DataRepository[T any] interface {
	GetAllData() ([]T, error)
	SaveData(data T) (T, error)
	GetData(entryID string) (T, error)
	EditData(data T) (T, error)
	DeleteData(entryID string) (T, error)
}

type DataItem interface {
	GetID() string
}
