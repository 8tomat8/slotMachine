package atkinsDiet

type Symbol uint8

func (s Symbol) String() string {
	return symbolsName[s]
}

const (
	// DO NOT CHANGE ORDER!
	Wild Symbol = iota
	Steak
	Ham
	BuffaloWings
	Sausage
	Eggs
	Butter
	Cheese
	Bacon
	Mayonnaise
	Scatter
)

var Symbols = [11]Symbol{Wild, Steak, Ham, BuffaloWings, Sausage, Eggs, Butter, Cheese, Bacon, Mayonnaise, Scatter}

var symbolsName = map[Symbol]string{
	Wild:         "Wild",
	Steak:        "Steak",
	Ham:          "Ham",
	BuffaloWings: "BuffaloWings",
	Sausage:      "Sausage",
	Eggs:         "Eggs",
	Butter:       "Butter",
	Cheese:       "Cheese",
	Bacon:        "Bacon",
	Mayonnaise:   "Mayonnaise",
	Scatter:      "Scatter",
}

type Reel [32]Symbol

var Reals = [5]Reel{
	{Scatter, Mayonnaise, Ham, Sausage, Bacon, Eggs, Cheese, Mayonnaise, Sausage, Butter, BuffaloWings, Bacon, Eggs, Mayonnaise, Steak, BuffaloWings, Butter, Cheese, Eggs, Wild, Bacon, Mayonnaise, Ham, Cheese, Eggs, Scatter, Butter, Bacon, Sausage, BuffaloWings, Steak, Butter},
	{Mayonnaise, BuffaloWings, Steak, Sausage, Cheese, Mayonnaise, Ham, Butter, Bacon, Steak, Sausage, Mayonnaise, Ham, Wild, Butter, Eggs, Cheese, Bacon, Sausage, BuffaloWings, Scatter, Mayonnaise, Butter, Cheese, Bacon, Eggs, BuffaloWings, Mayonnaise, Steak, Ham, Cheese, Bacon},
	{Ham, Butter, Eggs, Scatter, Cheese, Mayonnaise, Butter, Ham, Sausage, Bacon, Steak, BuffaloWings, Butter, Mayonnaise, Cheese, Sausage, Eggs, Bacon, Mayonnaise, BuffaloWings, Ham, Sausage, Bacon, Cheese, Eggs, Wild, BuffaloWings, Bacon, Butter, Cheese, Mayonnaise, Steak},
	{Ham, Cheese, Wild, Scatter, Butter, Bacon, Cheese, Sausage, Steak, Eggs, Bacon, Mayonnaise, Sausage, Cheese, Butter, Ham, Mayonnaise, Bacon, BuffaloWings, Sausage, Cheese, Eggs, Butter, BuffaloWings, Bacon, Mayonnaise, Eggs, Ham, Sausage, Steak, Mayonnaise, Bacon},
	{Bacon, Scatter, Steak, Ham, Cheese, Sausage, Butter, Bacon, BuffaloWings, Cheese, Sausage, Ham, Butter, Steak, Mayonnaise, Eggs, Sausage, Ham, Wild, Butter, BuffaloWings, Mayonnaise, Eggs, Ham, Bacon, Butter, Steak, Mayonnaise, Sausage, Eggs, Cheese, BuffaloWings},
}

type SpinKind string

const (
	Main SpinKind = "main"
	Free SpinKind = "free"
)

type SpinResult struct {
	Type  SpinKind
	Total uint64
	Stops [5]uint8
}

var lines = [20][5]uint8{
	{1, 1, 1, 1, 1}, // 1
	{0, 0, 0, 0, 0}, // 2
	{2, 2, 2, 2, 2}, // 3
	{0, 1, 2, 1, 0}, // 4
	{2, 1, 0, 1, 2}, // 5
	{1, 0, 0, 0, 1}, // 6
	{1, 2, 2, 2, 1}, // 7
	{0, 0, 1, 2, 2}, // 8
	{2, 2, 1, 0, 0}, // 9
	{1, 0, 1, 2, 1}, // 10
	{1, 2, 1, 0, 1}, // 11
	{0, 1, 1, 1, 0}, // 12
	{2, 1, 1, 1, 2}, // 13
	{0, 1, 0, 1, 0}, // 14
	{2, 1, 2, 1, 2}, // 15
	{1, 1, 0, 1, 1}, // 16
	{1, 1, 2, 1, 1}, // 17
	{0, 0, 2, 0, 0}, // 18
	{2, 2, 0, 2, 2}, // 19
	{0, 2, 2, 2, 0}, // 20
}

// payment for the scaters is calculated separately
var payTable = [len(Symbols)][6]uint16{
	//0,1, 2, 3,   4,    5
	{0, 0, 5, 50, 500, 5000}, // 0 - Wild
	{0, 0, 3, 40, 200, 1000}, // 1 - Steak
	{0, 0, 2, 30, 150, 500},  // 2 - Ham
	{0, 0, 2, 25, 100, 300},  // 3 - BuffaloWings
	{0, 0, 0, 20, 75, 200},   // 4 - Sausage
	{0, 0, 0, 20, 75, 200},   // 5 - Eggs
	{0, 0, 0, 15, 50, 100},   // 6 - Butter
	{0, 0, 0, 15, 50, 100},   // 7 - Cheese
	{0, 0, 0, 10, 25, 50},    // 8 - Bacon
	{0, 0, 0, 10, 25, 50},    // 9 - Mayonnaise
	{0, 0, 0, 5, 25, 100},    // 10 - Scatter
}
