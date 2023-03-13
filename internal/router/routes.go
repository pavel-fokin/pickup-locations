package router

type Location struct {
	Lat  float64
	Long float64
}

type Route struct {
	Destination Location
	Duration    float64
	Distance    float64
}

type ByTimeAndDistance []Route

func (a ByTimeAndDistance) Len() int      { return len(a) }
func (a ByTimeAndDistance) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTimeAndDistance) Less(i, j int) bool {
	return a[i].Duration < a[j].Duration || a[i].Distance < a[j].Distance
}
