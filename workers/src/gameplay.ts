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
// github:kevindamm/q-party/workers/src/gameplay.ts

import CHALLENGE_FORM_HTML from '../htmx/challenge_form.html'

import { DurableObject } from 'cloudflare:workers'
import { WorkerContext, WorkerEnv } from "../types"

// Stub representation for the time being.
export class GameplayServer extends DurableObject {
  host = [] as WebSocket[]
  players = [] as WebSocket[]
  viewers = [] as WebSocket[]

  constructor(state: DurableObjectState, env: WorkerContext) {
    super(state, env)

    state.blockConcurrencyWhile(async () => {
      // The websockets are tagged with their roles in the game.
      this.host = state.getWebSockets('host')
      this.players = state.getWebSockets('player')
      this.viewers = state.getWebSockets('viewer')
    })
  }

  async webSocketMessage(ws: WebSocket, message: ArrayBuffer | string) {
    ws.send(
      `{\n  "ack": ${message},\n  hosted: ${this.host.length > 0},\n  count_players: ${this.players.length},\n  count_viewers: ${this.players.length}\n}`,
    );
  }

  async webSocketClose(
    ws: WebSocket,
    code: number,
    /* reason: string, */
    /* wasClean: boolean, */
  ) {
    // If the client closes the connection, the runtime will invoke the webSocketClose() handler.
    ws.close(code, "Durable Object is closing WebSocket");
  }
}

export async function ChallengeForm(c: WorkerContext): Promise<Response> {
  // TODO base page for picking out categories & constraints (date range, difficulty, etc.)
  return new Response(CHALLENGE_FORM_HTML)
}

export async function InitChallenge(c: WorkerContext): Promise<Response> {
  // TODO /join/:roomid or /host/:roomid depending on whether a host exists yet.
  return new Response()
}

export async function JoinGame(c: WorkerContext): Promise<Response> {
  const roomname = c.req.param('roomid')
  if (c.env.SECRET_ROOM !== roomname) {
    return new Response('Game does not exist', { status: 404 })
  }

  return new Response('TODO game board grid as HTML response, opaque')
}

export async function HostGame(c: WorkerContext): Promise<Response> {
  const roomname = c.req.param('roomid')
  if (c.env.SECRET_ROOM !== roomname) {
    return new Response('Game does not exist', { status: 404 })
  }

  return new Response('TODO game board grid plus host controls')
}

export async function WatchGame(c: WorkerContext): Promise<Response> {
  // TODO check valid room
  // TODO stream updates with SSE, no websockets

  return new Response('TODO game board with simplified controls (mute/fullscreen), no buzzer')
}

export async function ResignGame(c: WorkerContext): Promise<Response> {
  // TODO check valid room
  // TODO check user in room
  // TODO check user is a contestant

  return new Response('TODO game over, forward to spectate & summary')
}
