import handlebars from "npm:handlebars"
import { pascalCase } from "npm:change-case"
import { pokeapi, write } from "./helpers.ts"

const enumsTmpl = handlebars.compile(`package types

const (
	None Type = iota
	{{#each types}}
	{{this}}
	{{/each}}
)
`)

const typesTmpl = handlebars.compile(`package types

func init() {
	Register({{type}}Data)
}

var {{type}}Data TypeData = TypeData{
	Type: {{type}},

	DoubleDamageTaken: []Type{
		{{#each doubleDamageTaken}}
		{{this}},
		{{/each}}
	},
	HalfDamageTaken  : []Type{
		{{#each halfDamageTaken}}
		{{this}},
		{{/each}}
	},
	ZeroDamageTaken  : []Type{
		{{#each zeroDamageTaken}}
		{{this}},
		{{/each}}
	},
}
`)

const { results } = await pokeapi("/api/v2/type")

const code = enumsTmpl({
	types: results
		.map(i => pascalCase(i.name))
		.filter(i => !["Shadow", "Unknown"].includes(i)),
})
await write(`types/enum.go`, code)
await Promise.all(
	results.map(async i => {
		const res = await pokeapi(i.url)
		const code = typesTmpl({
			type: pascalCase(i.name),
			doubleDamageTaken: res.damage_relations.double_damage_from.map(i => pascalCase(i.name)),
			halfDamageTaken: res.damage_relations.half_damage_from.map(i => pascalCase(i.name)),
			zeroDamageTaken: res.damage_relations.no_damage_from.map(i => pascalCase(i.name)),
		})
		await write(`types/${i.name}.go`, code)
	})
)