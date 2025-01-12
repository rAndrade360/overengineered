package domain

type Player struct {
	ID                   int
	Name                 string
	Firstname            string
	Lastname             string
	Age                  int
	BirthDate            string
	BirthPlace           string
	BirthCountry         string
	Nationality          string
	Height               string
	Weight               string
	Injured              string
	Photo                string
	TeamID               int
	TeamName             string
	TeamLogo             string
	LeagueID             int
	LeagueName           string
	LeagueCountry        string
	LeagueLogo           string
	LeagueFlag           string
	LeagueSeason         int
	GamesAppearences     int
	GamesLineups         int
	GamesMinutes         int
	GamesNumber          string
	GamesPosition        string
	GamesRating          float64
	GamesCaptain         string
	SubstitutesIn        int
	SubstitutesOut       int
	SubstitutesBench     int
	ShotsTotal           int
	ShotsOn              int
	GoalsTotal           int
	GoalsConceded        int
	GoalsAssists         int
	GoalsSaves           int
	PassesTotal          int
	PassesKey            int
	PassesAccuracy       int
	TacklesTotal         int
	TacklesBlocks        int
	TacklesInterceptions int
	DuelsTotal           int
	DuelsWon             int
	DribblesAttempts     int
	DribblesSuccess      int
	DribblesPast         string
	FoulsDrawn           int
	FoulsCommitted       int
	CardsYellow          int
	CardsYellowred       int
	CardsRed             int
	PenaltyWon           int
	PenaltyCommited      int
	PenaltyScored        int
	PenaltyMissed        int
	PenaltySaved         int
}
