package moves

type MoveFlags uint32

// Reference: "https://github.com/smogon/pokenmon-showdown/tree/master/sim/dex-moves.ts#L26-L49"
const (
	None = 1 << iota >> 1
	// Ignores a target's substitute.
	BypassSubsitute
	// Power is multiplied by 1.5 when used by a Pokemon with the Ability Strong Jaw.
	Bite
	// Has no effect on Pokemon with the Ability Bulletproof.
	Bullet
	// The user is unable to make a move between turns.
	Charge
	// Makes contact.
	Contact
	// When used by a Pokemon, other Pokemon with the Ability Dancer can attempt to execute the same move.
	Dance
	// Thaws the user if executed successfully while the user is frozen.
	Defrost
	// Can target a Pokemon positioned anywhere in a Triple Battle.
	Distance
	// Prevented from being executed or selected during Gravity's effect.
	Gravity
	// Prevented from being executed or selected during Heal Block's effect.
	Heal
	// Can be copied by Mirror Move.
	Mirror
	// The move has an animation when used on an ally.
	AllyAnimation
	// Prevented from being executed or selected in a Sky Battle.
	NonSky
	// Has no effect on Pokemon which are Grass-type, have the Ability Overcoat, or hold Safety Goggles.
	Powder
	// Blocked by Detect, Protect, Spiky Shield, and if not a Status move, King's Shield.
	Protect
	// Power is multiplied by 1.5 when used by a Pokemon with the Ability Mega Launcher.
	Pulse
	// Power is multiplied by 1.2 when used by a Pokemon with the Ability Iron Fist.
	Punch
	// If this move is successful, the user must recharge on the following turn and cannot make a move.
	Recharge
	// Bounced back to the original user by Magic Coat or the Ability Magic Bounce.
	Reflectable
	// Can be stolen from the original user and instead used by another Pokemon using Snatch.
	Snatch
	// Has no effect on Pokemon with the Ability Soundproof.
	Sound
)