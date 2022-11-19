import handlebars from "npm:handlebars";
import { pascalCase } from "npm:change-case";
import { pokeapi, write } from "./helpers.ts";

const enumsTmpl = handlebars.compile(`package types

const (
	None TypeBit = 1 << iota >> 1
	{{#each types}}
	{{this}}
	{{/each}}
)
`);

const typesTmpl = handlebars.compile(`package types

import (
	"github.com/falentio/poke/types"
)

func init() {
	types.Register({{type}}Data)
}

var {{type}}Data types.TypeData = types.TypeData{
	Bit: types.{{type}},

	DoubleDamageTaken: types.None{{#each doubleDamageTaken}} | types.{{this}}{{/each}},
	HalfDamageTaken:   types.None{{#each halfDamageTaken}} | types.{{this}}{{/each}},
	ZeroDamageTaken:   types.None{{#each zeroDamageTaken}} | types.{{this}}{{/each}},
}
`);

let { results } = await pokeapi("/api/v2/type");
results = results.filter((i) => !["shadow", "unknown"].includes(i.name));

const code = enumsTmpl({
	types: results
		.map((i) => pascalCase(i.name)),
});
await write(`types/enum.go`, code);
await Promise.all(
	results
		.map(async (i) => {
			const res = await pokeapi(i.url);
			const code = typesTmpl({
				type: pascalCase(i.name),
				doubleDamageTaken: res.damage_relations.double_damage_from
					.map((i) => pascalCase(i.name))
					.sort((a, b) => a.localeCompare(b)),
				halfDamageTaken: res.damage_relations.half_damage_from
					.map((i) => pascalCase(i.name))
					.sort((a, b) => a.localeCompare(b)),
				zeroDamageTaken: res.damage_relations.no_damage_from
					.map((i) => pascalCase(i.name))
					.sort((a, b) => a.localeCompare(b)),
			});
			await write(`types/data/${i.name}.go`, code.replace(/types.None \| /g, ""));
		}),
);
