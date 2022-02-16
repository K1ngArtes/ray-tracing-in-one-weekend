package trace

type List struct {
	hitList []Hittable
}

func NewList(hitList ...Hittable) *List {
	return &List{hitList: hitList}
}

func (l *List) Add(hl ...Hittable) int {
	l.hitList = append(l.hitList, hl...)
	return len(l.hitList)
}

func (l *List) Hittables() []Hittable {
	return l.hitList
}

func (l *List) Hit(r Ray, tMin float64, tMax float64, hitRecord *Hit) bool {
	var tempHitRec Hit
	isHit := false
	closestSoFar := tMax

	for _, o := range l.hitList {
		if o.Hit(r, tMin, closestSoFar, &tempHitRec) {
			isHit = true
			closestSoFar = tempHitRec.T
			hitRecord = &tempHitRec 
		}
	}

	return isHit
}