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

import { DurableObject } from 'cloudflare:workers'
import { WorkerContext } from '../types'

class Bucket {
  count: number

  constructor(
    readonly timestamp_s: number,
    readonly window_size: number,
    count: number) {
    this.count = count || 0
  }

  Contains = (ts: number) => ts < (this.timestamp_s + this.window_size);

  AddCount = (delta: number) => { this.count += delta }
}

export class RateLimiter extends DurableObject {
  // It's ok if this value occasionally gets reset, it won't happen often if the buckets are being filled
  buckets: Bucket[] = []
  capacity: number = 1000
  lockUntil: number

  static readonly millis_per_request = 5;
  static readonly millis_of_grace = 5000;
  
  constructor(state: DurableObjectState, env: WorkerContext) {
    super(state, env)
    this.lockUntil = 0
  }

  async getMillisUntilNextRequest(): Promise<number> {
    const now = Date.now()

    this.lockUntil = Math.max(now, this.lockUntil)
    this.lockUntil += RateLimiter.millis_per_request

    const value = Math.max(0,
      this.lockUntil - now - RateLimiter.millis_of_grace)
    return value
  }
}
