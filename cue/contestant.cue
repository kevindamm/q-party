package qparty

// Uniquely identifies a contestant across episodes.
#ContestantID: {
  id!: int
  name?: string
  ...
}

// Additional details about the contestant.
#Contestant: #ContestantID & {
  name!: string
  bio: string

  episodes?: [...#EpisodeID]
}

// An appearance is the joining of a contestant and an episode.
#Appearance: #ContestantID & #EpisodeID
