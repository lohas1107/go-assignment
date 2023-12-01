package matching

const (
	Boy Gender = iota
	Girl
)

type Gender int

func (g Gender) String() string {
	switch g {
	case Boy:
		return "Boy"
	case Girl:
		return "Girl"
	default:
		return "Unknown"
	}
}

type Single struct {
	Gender Gender `json:"gender"`
	Height int    `json:"height"`
}

func (s *Single) IsBoy() bool {
	return Boy == s.Gender
}

func (s *Single) IsGirl() bool {
	return Girl == s.Gender
}
