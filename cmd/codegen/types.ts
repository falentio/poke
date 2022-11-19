import handlebars from "npm:handlebars"
import { pascalCase } from "npm:change-case"
import { pokeapi, write } from "./helpers.ts"

const enumsTmpl = handlebars.compile(`package types

const (
	None TypeBit = 0

	{{#each types}}
	{{this}}{{#if @first}} TypeBit = 1 << iota >> 1{{/if}}
	{{/each}}
)
`)

const typesTmpl = handlebars.compile(`package types

func init() {
	Register({{type}}Data)
}

var {{type}}Data TypeData = TypeData{
	Bit: {{type}},

	DoubleDamageTaken: None{{#each doubleDamageTaken}} | {{this}}{{/each}},
	HalfDamageTaken:   None{{#each halfDamageTaken}} | {{this}}{{/each}},
	ZeroDamageTaken:   None{{#each zeroDamageTaken}} | {{this}}{{/each}},
}
`)

let { results } = await pokeapi("/api/v2/type")
results = results.filter(i => !["shadow", "unknown"].includes(i.name))

const code = enumsTmpl({
	types: results
		.map(i => pascalCase(i.name))
})
await write(`types/enum.go`, code)
await Promise.all(
	results
		.map(async i => {
			const res = await pokeapi(i.url)
			const code = typesTmpl({
				type: pascalCase(i.name),
				doubleDamageTaken: res.damage_relations.double_damage_from
					.map(i => pascalCase(i.name))
					.sort((a, b) => a.localeCompare(b)),
				halfDamageTaken: res.damage_relations.half_damage_from
					.map(i => pascalCase(i.name))
					.sort((a, b) => a.localeCompare(b)),
				zeroDamageTaken: res.damage_relations.no_damage_from
					.map(i => pascalCase(i.name))
					.sort((a, b) => a.localeCompare(b)),
			})
			await write(`types/${i.name}.go`, code)
		})
)