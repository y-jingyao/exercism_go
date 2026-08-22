package highscores

import "sort"

type HighScores struct {
	scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	copied := make([]int, len(scores))
	copy(copied, scores)
	return &HighScores{scores: copied}
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	out := make([]int, len(s.scores))
	copy(out, s.scores)
	return out
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.scores[len(s.scores)-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	maxh := s.scores[0]
	for _, v := range s.scores {
		if v > maxh {
			maxh = v
		}
	}
	return maxh
}

func (s *HighScores) TopThree() []int {
	cp := make([]int, len(s.scores))
	copy(cp, s.scores)
	sort.Slice(cp, func(i, j int) bool {
		return cp[i] > cp[j]
	})
	var topVals []int
	for i, num := range cp {
		if len(topVals) >= 3 || i >= len(s.scores) {
			break
		}
		topVals = append(topVals, num)
	}
	
	return topVals
}
