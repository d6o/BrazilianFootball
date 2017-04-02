package main

import "flag"

func main() {

	team := flag.String("team", "", "Filter by Team")
	championship := flag.String("championship", "", "Filter by Championship")
	status := flag.String("status", "", "Filter by Status")
	date := flag.String("date", "", "Filter by Date")
	dayOfWeek := flag.String("dayOfWeek", "", "Filter by Day of week")
	phase := flag.String("phase", "", "Filter by Round phase")
	location := flag.String("location", "", "Filter by Location")
	hour := flag.String("hour", "", "Filter by Hour")
	today := flag.Bool("today", false, "Filter by games that will ocour Today")
	upcoming := flag.Bool("upcoming", false, "Hide finished matches")

	limit := flag.Int("limit", 5, "Limit output lines")

	flag.Parse()

	bfData := NewBFData()
	bfData.Team(*team)
	bfData.Championship(*championship)
	bfData.Status(*status)
	bfData.Date(*date)
	bfData.DayOfWeek(*dayOfWeek)
	bfData.Phase(*phase)
	bfData.Location(*location)
	bfData.Hour(*hour)
	bfData.Today(*today)
	bfData.Upcoming(*upcoming)

	games, err := bfData.End()
	if err != nil {
		panic(err)
	}
	if len(games) > *limit {
		games = games[:*limit]
	}

	for _, game := range games {
		game.Print()
	}
}
