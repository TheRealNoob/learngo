package apis

type user struct {
	first string
	last  string
}

func (u *user) changeName(newFirst string, newLast string) {
	u.first = newFirst
	u.last = newLast
}
