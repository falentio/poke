import handlebars from "npm:handlebars";
import { pascalCase, snakeCase } from "npm:change-case";
import { pokeapi, write } from "./helpers.ts";
import { Moves } from "./data/moves.ts"

const moveTmpl = handlebars.compile(`package move_data

import (
	"github.com/falentio/poke/move"
	"github.com/falentio/poke/types"
)

func init() {
	move.Register({{name}}Data)
}

var {{name}}Data move.Move = move.Move{
	Id: {{id}},
	Type: types.Type{{type}},
	Flag: move.FlagNone {{#each flag}} | move.Flag{{this}}{{/each}}
}

`)

for (const move of Object.values(Moves)) {
	const code = moveTmpl({
		name: move.name.replace(/\s/g, ""),
		id: move.id,
		type: move.type,
		target: pascalCase(move.target),
		flag: Object.keys(move.flags).filter(k => move[k] == 1).map(k => pascalCase(k)),
	})
		.replace("move.FlagNone | ", "")
}