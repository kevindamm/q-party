# ?-Party Server Binary

Web server with sqlite database to host trivia games with friends.

## Building the server binary

```console
$ go build cmd/server
$ ./server -help
```

## Setting up the service dependencies

The server assumes that the environment has been initialized, needing at least
the database tables initialized and the metadata index populated, as well as
having public and private keys generated for serving over TLS.


### Initialize database tables

```console
$ go run cmd/jarchive -embed-init -write_db=jarchive.sqlite
```

See more details in the [jarchive README](../jarchive/README.md)


### Create TLS private and public keys

If you're planning on routing directly to the server, you can use Lets Encrypt,
or AutoTLS with a slight change to the server setup.  The current implementation
(at q-party.kevindamm.com) is routed through CloudFlare's proxy network, so it
builds a public key chain that includes CF's public key, using its origin certs.

Details of setting this up will vary by your environment, so documentation is
outside the scope of this README.

The server code assumes the certificates are at server.key and server.crt,
relative to the path specified by -cert_path (defaults to the current path).
