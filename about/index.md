---
# https://vitepress.dev/reference/default-theme-home-page
layout: home
title: Documentation

hero:
  name: "About ?-Party"
  tagline: "a trivial side project"

prev: false
next: false
---

TODO system diagram at the top
(workers, durable objects, D1 (data), R2 (media), AI (whisper),  )

## Environment and Dependencies

`?-Party` has a TypeScript implementation (using Hono) and a Go implementation
is being developed (using echo) for self-hosting and local-first playability.

There is a database of trivia questions and gameplay history, some of which
has been vetted and some which was vetted a long time ago and may be inaccurate.
The challenge-response and challenge-media details are given quality ratings
based on trusted peers doing fact-checks while they play.

There is a websocket server, facilitated by Cloudflare's Durable Objects and
implemented on the client side with htmx and its websocket extension.
The self-hosted version does not implement websockets, it is meant for the host
to manage communication and connect to bespoke buzzer circuits.

TODO more details

## System Design

Check back later!

TODO system diagram
