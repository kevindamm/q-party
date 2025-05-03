# Source Code for "About ?-Party" website

See the [about site](https://q-party.kevindamm.com/about), hosted on Cloudflare
Pages, for the rendered content.

This is the source code for that page, which I suppose makes this README
an **about about the about/**.

## Static Site Generation

Vitepress works so well as a documentation format.  I enjoy its mix of markdown
and mermaid, as well as source snippets and Vue components.  There are some
interactive diagrams and novel layouts that I want to try with this project.

## Cloudflare Pages as host

They have a very generous free tier and that makes it easy to stand up
continuous-integration and other build automation.  I would also consider
GitHub pages but the ?-Party site itself is already being hosted on Cloudflare
Workers -- besides it being easier to set up routing between Pages and Workers,
Cloudflare Pages doesn't have the restrictions that GitHub Pages has on content.

## HTTPS but on a different auth stack than the gameplay

These pages could have been generated into a dist-adjacent output and bundled
with the rest of the ?-Party site, but the resources on the About page are heavy
enough, and unprotected from the primary authn/authz flow;
I wanted to avoid the easy DoS vector that this site presents, at least from the
same workers that are serving the auth and games themselves.

Besides routing these pages through different servers, it also facilitates
serving them from a path-oriented but SPA-compatible router (cf Pages) whereas
the main application can still serve from the process- and service-oriented
router (via Workers).  For example, static pages should fall back on the Pages
router (main), but if they were both on Workers then that router is shared.

## Inline references to source code

I see some instructive potential here but I will wait to say more about it.
The near-term goal is to have a basic description of the architecture and flow
control -- there are quite a few moving parts --
so having that much in these pages is interesting on its own.

I have a side project that describes the common recipes I've encountered for
building a SaaS business and this may get folded into that project, or be a
proving ground for some of the embedded interactivity made possible with this
kind of documentation site.
