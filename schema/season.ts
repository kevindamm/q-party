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
// github:kevindamm/q-party/schema/season.ts

import * as z from "@zod/mini"

import { CategoryIndex } from "./category"
import { MatchIndex } from "./match"
import { ShowDateRange } from "./show_date"
import { MaybePositiveInt } from "./util"

export const SeasonSlug = z.string()
  .check(z.regex(/^[a-zA-Z][0-9a-zA-Z_-]*$/))
  .brand("SeasonSlug")

export const SeasonID = z.object({
  slug: SeasonSlug,
  title: z.optional(z.string().check(z.trim(), z.minLength(1))),
})

export const SeasonMetadata = z.extend(SeasonID, {
  aired: ShowDateRange,

  episode_count: MaybePositiveInt,
  category_count: MaybePositiveInt,
  challenge_count: MaybePositiveInt,
  tripstump_count: MaybePositiveInt,
})

export const SeasonIndex = z.map(SeasonSlug, SeasonMetadata)

export const SeasonDirectory = z.object({
  version: z.array(z.number()),
  seasons: SeasonIndex,
})

export const Season = z.extend(SeasonMetadata, {
  episodes: MatchIndex,
  categories: CategoryIndex,
})
