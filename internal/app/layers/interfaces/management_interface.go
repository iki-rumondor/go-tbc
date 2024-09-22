package interfaces

type ManagementInterface interface {
	CreateModel(pointerModel interface{}) error
}
