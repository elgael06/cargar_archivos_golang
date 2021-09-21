package data

import "curso_1/src/structs"

type allTask []structs.task

var tasks = allTask{
	{
		ID:      1,
		Name:    "task one",
		Content: "some content",
	},
}
