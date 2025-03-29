// Copyright (c) 2025, Kevin Damm
// All Rights Reserved.
// BSD 3-Clause License:
// 
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
// 
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
// 
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
// 
// 3. Neither the name of the copyright holder nor the names of its
//    contributors may be used to endorse or promote products derived from
//    this software without specific prior written permission.
// 
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.
// 
// github:kevindamm/q-party/workers/src/category.ts

import { z } from 'zod'
import { OpenAPIRoute } from 'chanfana'
import { WorkerContext } from '../types'

/**
 * GET /category
 */ 
export class CategoryIndex extends OpenAPIRoute {
  schema = {
    tags: ['Category'],
    summary: 'Get a summary of the entire category index',

    response: {
      "200": {
        description: 'all GET requests should succeed with category listing',
        content: {
          "text/html": z.string(),
        }
      }
    }
  }

  async handle(c: WorkerContext) {
    // list of seasons and list of recent category names are small enough
    // to fetch on first request and cache for ~ hours
  }
}

/**
 * GET /category/:catname
 */
export class DescribeCategory extends OpenAPIRoute {
  schema = {
    request: {
      params: z.object({
        catname: z.string().nonempty(),
      })},

    responses: {
      "200": { description: 'details for the category being named' },
      "404": { description: 'a category by that name was not found' },
    }
  }

  async handle(c: WorkerContext) {
    const { catname } = (await this.getValidatedData<typeof this.schema>()).params

    // TODO read from cache the category/:catname page

    // TODO lookup category named by the url path
  }
}

/**
 * GET /catseas/:season
 */
export class SeasonCategoryList extends OpenAPIRoute {
  schema = {
    request: {
      params: z.object({
        season: z.string().nonempty(),
      })},

    responses: {
      "200": { description: 'listing of categories that appeared during a season' },
      "404": { description: 'a season by that name was not found' }
    }
  }

  async handle(c: WorkerContext) {
    const { season } = (await this.getValidatedData<typeof this.schema>()).params

    // TODO read from cache when this season's categories have been fetched recently

    // TODO lookup season's category index
  }
}

/**
 * (locally) /catwhen#${from_YYYY[MM[DD]]}-${until_YYYY[MM[DD]]}
 * GET /catwhen/
 * GET /catwhen/:year
 * GET /catwhen/:year/:month
 */
export class DateRangeCategoryList extends OpenAPIRoute {

}
