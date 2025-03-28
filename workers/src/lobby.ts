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
// github:kevindamm/q-party/workers/src/lobby.ts

import ROOM_FORM_HTML from '../htmx/room_form.html'

import { DurableObject } from "cloudflare:workers"
import { WorkerContext } from "../types"

// Stub representation for the time being.
export class LobbyServer extends DurableObject {
  // TODO properties of the runtime state

  constructor(state: DurableObjectState, env: WorkerContext) {
    super(state, env)
    // TODO state initialization
  }

  async fetch(request: Request): Promise<Response> {
    // TODO construct message
    return new Response()
  }

  // TODO getters and setters for persisted storage
}

/**
 * /join
 * @method GET
 */
export async function RoomForm(c: WorkerContext): Promise<Response> {
  // TODO check headers for resuming ongoing session

  return new Response(ROOM_FORM_HTML)
}

/**
 * /join
 * @method POST
 * @param jsonBody with { username, roomid, CSRF nonce }
 */
export async function JoinRoom(c: WorkerContext): Promise<Response> {
  // TODO validate room name and add user to the websocket/buzzer/UI-test

  return new Response('Room does not exist.', { status: 404 })
}

/**
 * /lobby/:roomid 
 * @method POST
 * @param jsonBody with { username, message, CSRF nonce }
 */
export async function PostMessage(c: WorkerContext): Promise<Response> {
  // TODO handle message via HTTP, open SSE for streaming messages
  return new Response()
}

/**
 * /lobby/:roomid
 * @method DELETE
 * @params formdata with username, CSRF nonce
 */
export async function LeaveRoom(c: WorkerContext): Promise<Response> {
  // TODO 
  return new Response()
}
