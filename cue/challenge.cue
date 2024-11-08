package qparty

// A unique identifier for the challenge and (optionally) its ID.
// If value is undefined it did not have an associated monetary value.
#ChallengeID: {
  id!: int
  value?: int
  ...
}


// The challenge details, sans the correct answer.
#Challenge: #ChallengeID & {
  value!: int
  clue!: string
  media?: [...#Media]

  is_wager?: bool
  category?: string
  comments?: string
}

// The host may see the correct answer while the contestants cannot.
#HostChallenge: #ChallengeID & {
  correct: string
}

// Before answering, sometimes a player must provide a wager value first.
#PlayerWager: #ChallengeID & {
  value!: int
  comments?: string
}

// The player's response for a challenge.
#PlayerResponse: #ChallengeID & {
  response?: string
}
