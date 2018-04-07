package model

type BasicInfo struct {
	Authors   []string
	Chapters  string
	Directors []string
	Hour      string // {6, 7, 8, 9, 10, 11}pm
	Name      string
	URL       string
	Year      string
}

type Novela struct {
	BasicInfo
	Actors []string
}

func (n *Novela) AppendActors(a []string) {
	n.Actors = append(n.Actors, a...)
}
