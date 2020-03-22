package mobile

import (
	"github.com/gopub/gox/geo"
	"github.com/gopub/types"
)

type Point types.Point

func NewPoint() *Point {
	return new(Point)
}

type Place types.Place

func NewPlace() *Place {
	return new(Place)
}

func (p *Place) SetLocation(location *Point) {
	p.Location = (*types.Point)(location)
}

func (p *Place) GetLocation() *Point {
	return (*Point)(p.Location)
}

type Country types.Country

type CountryList struct {
	List []*Country
}

func NewCountryList() *CountryList {
	return &CountryList{}
}

func (l *CountryList) Size() int {
	return len(l.List)
}

func (l *CountryList) Get(i int) *Country {
	return l.List[i]
}

func GetCountryList() *CountryList {
	l := geo.GetCountries()
	ll := NewCountryList()
	ll.List = make([]*Country, len(l))
	for i, c := range l {
		ll.List[i] = (*Country)(c)
	}
	return ll
}
