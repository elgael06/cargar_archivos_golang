package funtions

import (
	"github.com/elgael06/curso_1/data"
	"github.com/elgael06/curso_1/structs"
)

func GetAllTask() []structs.Task {
	return data.AllTasks
}
