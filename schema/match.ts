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
// github:kevindamm/q-party/schema/match.ts

import * as z from "@zod/mini"
import { SeasonSlug } from "./season"
import { ShowDate } from "./show_date"
import { ContestantID } from "./contestant"
import { MediaRef } from "./challenge"
import { BoardPosition } from "./round"

export const MatchNumber = z.int64()
    .check(z.positive())
    .brand("MatchNumber")

export const MatchID = z.object({
  match: MatchNumber,
  show_title: z.string().check(z.minLength(3)),
  season: SeasonSlug,
})

export const MatchMetadata = z.extend(MatchID, {
  aired: z.optional(ShowDate),
  taped: z.optional(ShowDate),

  contestants: z.optional(z.array(ContestantID)),
  media: z.optional(z.set(MediaRef)),
  comments: z.optional(z.string()),
})

export const MatchIndex = z.map(MatchNumber, MatchMetadata)

export const MatchStats = z.extend(MatchMetadata, {
  single_count: z.optional(z.int()),
  double_count: z.optional(z.int()),

  triple_stumpers: z.optional(z.set(BoardPosition)),
})
