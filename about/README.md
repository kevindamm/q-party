# About ?-Party (source code for site)

See the [about site](https://q-party.kevindamm.com/about) for the rendered view.

This is the source code for that page, which I suppose makes this README
an about about the about/.

## Static Site Generation

Vitepress works so well as a documentation format, with its mix of markdown
and mermaid and source snippets and Vue.

## Cloudflare Pages as host

They have a very generous free tier and that makes it easy to stand up
continuous-integration and other build automation.  I would also consider
GitHub pages but the ?-Party site itself is already being hosted on Cloudflare

## HTTPS but on a different auth stack than the gameplay

These pages could have been generated into a dist-adjacent output and bundled
with the rest of the ?-Party site, but the resources on the About page are heavy
enough, and unprotected from the primary authn/authz flow;
I wanted to avoid the easy DoS vector this presents.

Besides routing these pages through different servers, it also facilitates
serving them from a path-oriented but SPA-compatible router (cf Pages) whereas
the main application can still serve from the process- and service-oriented
router (via Workers).

## Inline references to source code

I see some instructive potential here but I will wait to say more about it.
The near-term goal is to have a basic description of the architecture and flow
control -- there are quite a few moving parts,
so that much is interesting on its own.  I have a side project that describes a
lot of common recipes for building a SaaS business and this may get folded into
that project, or be a proving ground for some of the embedded interactivity made
possible with this kind of documentation site.
