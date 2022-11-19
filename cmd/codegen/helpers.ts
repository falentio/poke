import { resolve } from "https://deno.land/std/path/mod.ts";

export const pokeapiUrl = new URL("https://pokeapi.co/");
export const pokeapi = async (path: string) => {
	const url = new URL(path, pokeapiUrl).href;
	const res = await fetch(url);
	if (!res.ok) {
		console.error(await res.text());
		throw new Error("error encountered");
	}
	return res.json();
};

export const write = (path: string, data: string) => {
	const p = resolve(new URL(import.meta.url).pathname, "../../../", path);
	return Deno.writeTextFile(p, data);
};
