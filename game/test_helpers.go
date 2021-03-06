// Copyright 2017 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Helper methods for use in tests in this package and others.

package game

func TestScore1() *Score {
	fouls := []Foul{{Rule{"G22", false}, 25, 25.2}, {Rule{"G18", true}, 25, 150},
		{Rule{"G20", true}, 1868, 0}}
	return &Score{0, 1, 2, 20, 1, 12, 55, 1, fouls, false}
}

func TestScore2() *Score {
	return &Score{2, 2, 10, 0, 2, 65, 24, 3, []Foul{}, false}
}

func TestRanking1() *Ranking {
	return &Ranking{254, 1, RankingFields{20, 625, 90, 554, 10, 50, 0.254, 3, 2, 1, 0, 10}}
}

func TestRanking2() *Ranking {
	return &Ranking{1114, 2, RankingFields{18, 700, 625, 90, 554, 9, 0.1114, 1, 3, 2, 0, 10}}
}
