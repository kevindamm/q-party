package qparty

// A board state, includes the minimum needed information for starting play.
#Board: #EpisodeID & {
  round: int & >=0 & <len(round_names)
  round_name: round_names[round]

  columns: [...#Category]
  missing?: [...#Position]
  history?: [...#Position]
}

// Display strings for the different rounds.
round_names: [...string] & [
  "[UNKNOWN]",
	"Single!",
	"Double!",
	"Final!",
	"Tiebreaker!!",
	"[printed media]",
]

// A category instance must have a title and
// may have any number of challenges (typically five).
#Category: {
  title!: string
  commentary?: string
  challenges: [...#ChallengeID]
}

// A board position is identified by its column and (row) index.
#Position: {
  column!: int & >=0 & <6
  index!: int & >=0 & <5
}

// Sentinel representation for any blank board cell.
UnknownChallenge: #ChallengeID & { id: 0 }
