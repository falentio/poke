package field

import "github.com/falentio/poke/pokemon"

type Weather uint8
const (
	WeatherClearSkies Weather = iota
	WeatherHarshSunlight
	WeatherExtremelyHarshSunlight
	WeatherRain
	WeatherHeavyRain
	WeatherSandstorm
	WeatherHail
	WeatherShadowyAura
	WeatherFog
	WeatherStrongWinds
)

type Terrain uint8
const (
	TerrainElectric Terrain = iota
	TerrainGrassy
	TerrainMisty
	TerrainPyshic
)

type Player struct{
	Pokemon []*pokemon.Pokemon
	Item []any
}

type Sides struct{
	Red []*Player
	Blue []*Player
}

type Field struct{
	Sides *Sides

	Weather Weather
	Terrain Terrain
}