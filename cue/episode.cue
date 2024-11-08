package qparty

// Episode identifier and the show number (based on airing order).
#EpisodeID: {
  id!: int
  show_number?: int
  ...
}

// Identifiers and statistics for each episode.
#EpisodeMetadata: #EpisodeID & {
  season: #SeasonID
  aired?: #ShowDate

  contestant_ids?: [int, int, int]
  comments?: string
  media?: [...#Media]

  single_clues?: int
  double_clues?: int
  final_category?: string
  triple_stumpers?: int
}
