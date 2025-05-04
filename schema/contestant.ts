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
// github:kevindamm/q-party/schema/contestant.ts

import * as z from "@zod/mini"
import { MatchID } from "./match"
import { MediaRef, Value } from "./challenge"

export const ContestantID = z.object({
  cid: z.uint64(),
  name: z.optional(z.string().check(z.minLength(4))),
})

export const Contestant = z.extend(ContestantID, {
  name: z.string(),
  occupation: z.optional(z.string()),
  residence: z.optional(z.string()),

  notes: z.optional(z.string()),
  media: z.set(MediaRef),
})

export const Appearance = z.extend(ContestantID, {
  match: MatchID,
})

export const Career = z.extend(ContestantID, {
  matches: z.array(MatchID),
  winnings: Value,
})

