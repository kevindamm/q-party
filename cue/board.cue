package qparty

#Board: {
  round: string
  columns: [...#Category]
  history?: [...#Position]
}

#Category: {
  title: string
  commentary?: string
  challenges: [...#Challenge]
}

#Position: {
  column: int & >=0 & <6
  index: int & >=0 & <5
}
