package qparty

#AllSeasons: {
  version?: [...int]
  seasons: [...#SeasonMetadata]
  episodes: [...#EpisodeMetadata]
}

#SeasonMetadata: {
  ident: string
  name: string

  aired: {
    from: #ShowDate
    until: #ShowDate
  }
  episode_count: int
}
