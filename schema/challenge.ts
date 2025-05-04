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
// github:kevindamm/q-party/schema/challenge.ts

import * as z from "@zod/mini"

import { ContestantID } from "./contestant"

export const Value = z.int()
  .brand("Value")
export const Wager = z.int()
  .check(z.positive())
  .brand("Wager")

export const ChallengeMetadata = z.object({
  qid: z.uint64().check(z.positive()),
  value: z.optional(Value),
})

export const UnknownChallenge = { qid: 0 }

export const MediaRef = z.object({
  mime: z.enum([
    "image/jpeg",
    "image/png",
    "image/svg+xml",
    "audio/mpeg",
    "video/mp4",
    "video/quicktime",
  ]),
  url: z.string(),
})

export const ChallengeData = z.extend(ChallengeMetadata, {
  clue: z.string(),
  media: z.optional(z.array(MediaRef)),
  category: z.optional(z.string()),
  comments: z.optional(z.string()),
})

export const Challenge = z.extend(ChallengeData, {
  value: Value,
})

export const BiddingChallenge = z.extend(ChallengeData, {
  wager: Wager,
})

export const HostChallenge = z.extend(Challenge, {
  correct: z.set(z.string()),
  value: z.optional(Value),
  wager: z.optional(Wager),
})

export const PlayerWager = z.extend(ContestantID,
  z.extend(ChallengeMetadata, {
    wager: Wager,
    comments: z.optional(z.string()),
  }))

export const PlayerResponse = z.extend(ContestantID,
  z.extend(ChallengeMetadata, {
    response: z.optional(z.string()),
  }))
