package qparty

#ChallengeID: {
  id!: int
  value?: int
  ...
}


#Challenge: #ChallengeID & {
  value!: int
  clue!: string
  media?: [...#Media]

  is_wager?: bool
  category?: string
  comments?: string
}

#HostChallenge: #ChallengeID & {
  correct: string
}

#PlayerWager: #ChallengeID & {
  comments: string
}

#PlayerResponse: #ChallengeID & {
  response?: string
}

UnknownChallenge: #ChallengeID & {
  id: 0
  value: 0
}