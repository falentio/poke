package move

type Flag uint32

// Reference: "https://github.com/smogon/pokenmon-showdown/tree/master/sim/dex-moves.ts#L26-L49"
const (
	FlagNone Flag = 1 << iota >> 1
	// Ignores a target's substitute.
	FlagBypassSubsitute
	// Power is multiplied by 1.5 when used by a Pokemon with the Ability Strong Jaw.
	FlagBite
	// Has no effect on Pokemon with the Ability Bulletproof.
	FlagBullet
	// The user is unable to make a move between turns.
	FlagCharge
	// Makes contact.
	FlagContact
	// When used by a Pokemon, other Pokemon with the Ability Dancer can attempt to execute the same move.
	FlagDance
	// Thaws the user if executed successfully while the user is frozen.
	FlagDefrost
	// Can target a Pokemon positioned anywhere in a Triple Battle.
	FlagDistance
	// Prevented from being executed or selected during Gravity's effect.
	FlagGravity
	// Prevented from being executed or selected during Heal Block's effect.
	FlagHeal
	// Can be copied by Mirror Move.
	FlagMirror
	// The move has an animation when used on an ally.
	FlagAllyAnimation
	// Prevented from being executed or selected in a Sky Battle.
	FlagNonSky
	// Has no effect on Pokemon which are Grass-type, have the Ability Overcoat, or hold Safety Goggles.
	FlagPowder
	// Blocked by Detect, Protect, Spiky Shield, and if not a Status move, King's Shield.
	FlagProtect
	// Power is multiplied by 1.5 when used by a Pokemon with the Ability Mega Launcher.
	FlagPulse
	// Power is multiplied by 1.2 when used by a Pokemon with the Ability Iron Fist.
	FlagPunch
	// If this move is successful, the user must recharge on the following turn and cannot make a move.
	FlagRecharge
	// Bounced back to the original user by Magic Coat or the Ability Magic Bounce.
	FlagReflectable
	// Can be stolen from the original user and instead used by another Pokemon using Snatch.
	FlagSnatch
	// Has no effect on Pokemon with the Ability Soundproof.
	FlagSound
)
