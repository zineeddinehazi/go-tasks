package pkg

type Task struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	IsDone  bool   `json:"isdone"`
}

var TaskList = []Task{
	{ID: "1", Content: "Build CRUD app using Go", IsDone: false},
	{ID: "2", Content: "Prepare for my university exams", IsDone: false},
	{ID: "3", Content: "Drink 30 cups of coffee", IsDone: false},
	{ID: "4", Content: "Complete my lab reports", IsDone: false},
}
