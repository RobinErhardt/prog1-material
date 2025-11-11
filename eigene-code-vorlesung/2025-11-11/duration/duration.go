package duration

type Duration int

func FromSeconds(s int) Duration {
	return Duration(s)
}

func FromMinutes(m int) Duration {
	return Duration(m * 60)
}

func FromHours(h int) Duration {
	return Duration(h * 3600)
}

func (t Duration) Seconds() int {
	return int(t)
}

func (t Duration) Minutes() int {
	return t.Seconds() / 60
}

func (t Duration) Hours() int {
	return t.Minutes() / 60
}