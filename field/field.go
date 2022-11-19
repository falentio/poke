package field

type Weather uint8

const (
	ClearSkies Weather = iota
	HarshSunlight
	ExtremelyHarshSunlight
	Rain
	HeavyRain
	Sandstorm
	Hail
	ShadowyAura
	Fog
	StrongWinds
)

type Terrain uint8

const (
	ElectricTerrain Terrain = iota
	GrassyTerrain
	MistyTerrin
	PyshicTerrain
)
