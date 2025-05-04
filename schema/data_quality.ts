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
// github:kevindamm/q-party/schema/data_quality.ts

import * as z from "@zod/mini"
import { ChallengeMetadata } from "./challenge"

export const DataQualityEnum = z.enum([
    "Needs Review",
    "Entirely Incorrect",
    "Recently Incorrect",
    "Suspected Outdated",
    "Needs Minor Change",
    "Disagreement",
    "Correct",
    "Confirmed Correct",
])

export const DataQuality = z.object({
  dqID: z.int().check(z.gte(0), z.lt(8)),
  quality: DataQualityEnum,
})

export const DataQualityJudgement = z.extend(ChallengeMetadata,
  z.extend(DataQuality, {
    comments: z.optional(z.string()),
  }))
