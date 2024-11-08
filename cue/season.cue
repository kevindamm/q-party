package qparty

// Schema for [seasons.json] containing season and episode metadata.
#AllSeasons: {
  version?: [...int]
  seasons: [...#SeasonMetadata]
  episodes: [...#EpisodeMetadata]
}


#SeasonID: {
  id!: string
  name?: string
  ...
}

// Metadata for a single season, has identity and some statistics.
#SeasonMetadata: #SeasonID & {
  aired: {
    from: #ShowDate
    until: #ShowDate
  }
  episode_count: int
  challenge_count: int
}
