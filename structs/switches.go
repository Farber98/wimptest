package structs

/* type SwitchWrapper struct {
	Switch Switches
}
*/
type Switches struct {
	Mac string  `bson:"mac" json:"mac,omitempty"`
	Lat float64 `bson:"lat" json:"lat,omitempty"`
	Lng float64 `bson:"lng" json:"lng,omitempty"`
}
