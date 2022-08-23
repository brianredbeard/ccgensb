package ccgensb

// Supported card types
const (
	AmericanExpress CardType = iota + 1
	DinersClub
	DinersClubUS
	Discover
	JCB
	Laser
	Maestro
	Mastercard
	Solo
	Unionpay
	Visa
	Mir
	ShadyBucks
)

// Holds valid card number length
var cardLength = map[CardType][]int{
	AmericanExpress: {15},
	DinersClub:      {14},
	DinersClubUS:    {16},
	Discover:        {16, 19},
	JCB:             {15, 16},
	Laser:           {16, 17, 18, 19},
	Maestro:         {12, 13, 14, 15, 16, 17, 18, 19},
	Mastercard:      {16},
	Solo:            {16, 18, 19},
	Unionpay:        {16, 17, 18, 19},
	Visa:            {13, 16, 19},
	Mir:             {13, 16},
	ShadyBucks:      {16},
}

// First N digits reserved by the card scheme
var cardPrefix = map[CardType][]string{
	AmericanExpress: {"34", "37"},
	DinersClub:      {"300", "301", "302", "303", "304", "305", "36"},
	DinersClubUS:    {"54", "55"},
	Discover: {"6011", "622126", "622127", "622128", "622129", "62213",
		"62214", "62215", "62216", "62217", "62218", "62219",
		"6222", "6223", "6224", "6225", "6226", "6227", "6228",
		"62290", "62291", "622920", "622921", "622922", "622923",
		"622924", "622925", "644", "645", "646", "647", "648",
		"649", "65"},
	JCB:   {"1800", "2131", "3528", "3529", "353", "354", "355", "356", "357", "358"},
	Laser: {"6304", "6706", "6771", "6709"},
	Maestro: {"5018", "5020", "5038", "6304", "6759", "6761", "6762", "6763",
		"6764", "6765", "6766", "6772"},
	Mastercard: {"2221", "2222", "2223", "2224", "2225", "2226", "2227", "2228", "2229",
		"223", "224", "225", "226", "227", "228", "229",
		"23", "24", "25", "26", "271", "2720",
		"51", "52", "53", "54", "55"},
	Solo: {"6334", "6767"},
	Unionpay: {"622126", "622127", "622128", "622129", "62213", "62214",
		"62215", "62216", "62217", "62218", "62219", "6222", "6223",
		"6224", "6225", "6226", "6227", "6228", "62290", "62291",
		"622920", "622921", "622922", "622923", "622924", "622925"},
	Visa:       {"4"},
	Mir:        {"2200", "2201", "2202", "2203", "2204"},
	ShadyBucks: {"89979866722"},
}
