// Copyright (c) 2025, Kevin Damm
// All Rights Reserved.
// BSD 3-Clause License
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
// github:kevindamm/q-party/workers/types/bindings.d.ts

import { Context } from 'hono'
import { RateLimiter } from './ratelimiter'
import { LobbyServer } from './lobby'
import { GameplayServer } from './gameplay'

export interface WorkerEnv {
  // Trivia challenges and categories, game and match histories, etc.
  DB: D1Database

  // Performs speech-to-text via Workers AI using whisper model.
  WHISPER: Ai

  // Durable objects which facilitate websocket connections.
  THROTTLE: DurableObjectNamespace<RateLimiter>
  LOBBIES: DurableObjectNamespace<LobbyServer>
  GAMEPLAY: DurableObjectNamespace<GameplayServer>

  // For storing and retrieving media such as audio and video,
  // used in supplementing the text of a trivia challenge.
  MEDIA: R2Bucket

  // A simplification, for now... one valid room, one working token.
  SECRET_ROOM: string
  VALID_TOKEN: string
}

export type WorkerContext = Context<{ Bindings: WorkerEnv}>
