package meander

type j struct {
	Name        string
	PlacesTypes []string
}

var Journeys = []any{
	&j{Name: "ロマンティック", PlacesTypes: []string{"park", "bar", "movie_theater", "restaurant", "florist", "taxi_stand"}},
	&j{Name: "ショッピング", PlacesTypes: []string{"department_store", "cafe", "clothng_store", "jewely_store", "shoe_store"}},
	&j{Name: "ロマンティック", PlacesTypes: []string{"bar", "casino", "food", "bar", "night_club"}},
}
