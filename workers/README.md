# ?-Party as a Cloudflare Worker

Using Workers and Durable Objects with Hibernation for handling realtime updates
(and listening for buzzer actions) over websockets.  It retrieves board layouts
and challenge (Q/A) details through D1, it handles various roles (Host,
Contestant and Spectators) and provides a tailored interface for each.  The
media files for any questions with audio and video clues are served from
R2 where bandwidth is especially generous.

<!-- TODO diagram basic system architecture -->
