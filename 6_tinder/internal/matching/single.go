package matching

type Single struct {
	Gender      string `json:"gender"`
	Height      int    `json:"height"`
	WantedDates int    `json:"wantedDates"`
}

func (s *Single) IsValidGender() bool {
	return s.IsBoy() || s.IsGirl()
}

func (s *Single) IsBoy() bool {
	return "BOY" == s.Gender
}

func (s *Single) IsGirl() bool {
	return "GIRL" == s.Gender
}
