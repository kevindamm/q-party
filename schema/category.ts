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
// github:kevindamm/q-party/schema/category.ts

import * as z from "@zod/mini"
import { ChallengeMetadata, MediaRef } from "./challenge"
import { ShowDate } from "./show_date"

export const CategoryID = z.string().brand("CategoryID")

export const CategoryMetadata = z.required(z.object({
  catID: CategoryID,
  title: z.string(),
}))

export const CategoryIndex = z.map(CategoryID, CategoryMetadata)

export const Category = z.extend(CategoryMetadata, {
  challenges: z.array(ChallengeMetadata),
  media: z.optional(z.set(MediaRef)),
  comments: z.optional(z.string()),
})

export const CategoryAired = z.extend(CategoryMetadata, {
  aired: ShowDate,
})
