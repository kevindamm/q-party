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
// github:kevindamm/q-party/workers/src/router.ts

import { Hono } from 'hono'
import { fromHono } from 'chanfana'
import { downgrade_protection } from './middleware/tls'
import { auth_required } from './middleware/auth'
import { WorkerEnv } from '../types'
import { logger } from 'hono/logger'

import { ChallengeForm, InitChallenge, JoinGame } from './gameplay'
import { RoomForm, JoinRoom, PostMessage, LeaveRoom } from './lobby'
import { AudioUI, TranscribeAudio } from './transcribe'
import { CategoryIndex, DateRangeCategoryList, DescribeCategory, SeasonCategoryList } from './category'

export { RateLimiter } from './ratelimiter'
export { LobbyServer } from './lobby'
export { GameplayServer } from './gameplay'

const app = new Hono<{Bindings: WorkerEnv}>()
const api = fromHono(app)

app.use(logger())
app.use(downgrade_protection)
app.use(auth_required)

// DEBUGGING [
app.get('/speak', AudioUI)
app.post('/speak', TranscribeAudio)
// ]

// Gameplay routes

app.get('/join', RoomForm)
app.put('/join/:userid', JoinRoom)
app.post('/lobby/:roomid', PostMessage)
app.delete('/lobby/:roomid/:userid', LeaveRoom)

app.get('/challenge', ChallengeForm)
app.post('/challenge', InitChallenge)
app.post('/play/:roomid', JoinGame)

// API methods

api.get('/category', CategoryIndex)
api.get('/category/:catname', DescribeCategory)

api.get('/catseas/', SeasonCategoryList)
api.get('/catseas/:season', SeasonCategoryList)

api.get('/catwhen', DateRangeCategoryList)
api.get('/catwhen/:year', DateRangeCategoryList)
api.get('/catwhen/:year/:month', DateRangeCategoryList)
api.get('/catwhen/:year/:month/:day', DateRangeCategoryList)

export default api
