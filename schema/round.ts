// Copyright (c) 2025 Kevin Damm
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//
// github:kevindamm/q-party/schema/round.ts

import * as z from "@zod/mini"

import { CategoryMetadata } from "./category"
import { ChallengeMetadata, Value } from "./challenge"
import { ContestantID } from "./contestant"
import { MatchNumber } from "./match"

export const RoundEnum = z.enum([
    "[UNKNOWN]",
    "Single!",
    "Double!",
    "Final!",
    "Tiebreaker!!",
    "[other]",
  ])

export const RoundID = z.object({
  episode: MatchNumber,
  round: RoundEnum,
})

export const BoardPosition = z.required(z.object({
  column: z.int().check(z.positive()),
  index: z.int().check(z.positive()),
}))

export const Board = z.extend(RoundID, {
  columns: z.array(CategoryMetadata),
  missing: z.set(BoardPosition),
})

export const BoardSelection = z.extend(ContestantID,
  z.extend(ChallengeMetadata, BoardPosition))

export const SelectionOutcome = z.extend(BoardSelection, {
  correct: z.boolean(),
  delta: Value,
})

export const BoardState = z.extend(Board, {
  history: z.array(SelectionOutcome),
})
