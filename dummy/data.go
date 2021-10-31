package dummy

type Event struct {
	ID          string
	Title       string
	Description string
}

type AllEvents []Event

var Events = AllEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
	{
		ID:          "2",
		Title:       "Introduction to Javascript",
		Description: "Come join us for a chance to learn how Javascript works and get to eventually try it out",
	},
	{
		ID:          "3",
		Title:       "Introduction to PHP",
		Description: "Come join us for a chance to learn how PHP works and get to eventually try it out",
	},
}
