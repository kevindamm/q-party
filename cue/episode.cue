package qparty

#EpisodeMetadata: {
  season: string
  show_number: int
  aired: #ShowDate
  comments: string
  media: [...#Media]

  contestants: [int, int, int]

  single_clues: int
  double_clues: int
  final_category: string
  triple_stumpers: int
}
