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
// github:kevindamm/q-party/schema/scanner.ts

import * as z from "@zod/mini"
import { SeasonDirectory, SeasonMetadata } from "./season"
import { MatchMetadata } from "./match"

type SeasonDirectoryType = z.infer<typeof SeasonDirectory>

function LoadSeasonDir(): SeasonDirectoryType {
  // TODO load season directory from parameter or local file
  return {
    version: [0, 0, 0, 20250404],
    seasons: new Map(),
  }
}

type SeasonMetadataType = z.infer<typeof SeasonMetadata>

function LoadSeason(season_slug: string): SeasonMetadataType {
  // TODO load season info
  return {aired: {}}
}

type MatchMetadataType = z.infer<typeof MatchMetadata>

function LoadMatch(match: number): MatchMetadataType {
  // TODO load match info
  return {}
}
