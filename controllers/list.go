package controllers

type List struct {
	ActionName string
	Method     string
	HttpMethod string
}

func NewList() *List {
	return &List{
		ActionName: "List",
		Method:     "list",
		HttpMethod: "GET",
	}
}

func (l *List) GetActionName() string {
	//TODO implement me
	panic("implement me")
}

func (l *List) GetMethod() string {
	//TODO implement me
	panic("implement me")
}

func (l *List) GetHttpMethod() string {
	//TODO implement me
	panic("implement me")
}

func (l *List) Do(args []string) (map[string]any, map[string]any) {
	//TODO implement me
	panic("implement me")
}
