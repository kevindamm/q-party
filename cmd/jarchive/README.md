# ?-Party command-line interpreter

Initializes the environment for serving from a local DB, building on existing
index tables or creating all the tables for the service to function.  This
includes user tables and session tables, from a hard-coded list of lobby names.


## Build the CLI binary

```console
$ go build cmd/jarchive
$ ./jarchive -help
jarchive [list|play]
  where
    *list* lists season or episode info
    *play* starts an interactive session

list takes arguments 'season' or 'episode' or no additional argument
  with no argument, it lists all seasons and count of available episodes
  with 'season' argument and season-ID it prints information about that season
  with 'episode' argument and 'season/show' ID, prints episode information

play takes arguments [season/show ID] (playing a selected show)
  or 'episode' or 'category' or 'challenge' (playing a random selection)

  -data-path string
        path where converted and created games are written (default "./.data")
  -db-path
        path of the SQLite database file, relative to -data-path
        (do not write to Db if "")
  -seed-db
        create & populate tables, and write all new category and challenge content to a local DB
```


## Example commands

Create the database tables and write the initial metadata.

```console
$ ./jarchive -seed-db -db-path=jarchive.sqlite
```

This will use the default data path (`./.data` from cwd) and will write the
embedded index to a sqlite file in that directory.  Further interactions may
require additional fetches, the command takes care of rate-limiting these
fetches to be considerate to the original hosts, and only the bare minimum of
external pages will be fetched to populate a local board or two.


### List available seasons

```console
$ ./jarchive list
Season 01 (s01) [164 episodes]
Season 02 (s02) [179 episodes]
...
```


### List episodes in a season

```console
$ ./jarchive list superjeopardy
```


## Local cache

To further reduce the load on the host, the local DB will cache the episode's
content (when -write_db is given a value).  The host must comply with the terms
of service for j-archive.com if persisting the contents of this database.

## Debug output

If making changes to the code, use `-log_path=debug.log` (with any file path)
to save the verbose log to a file.  Pass "-" as a filename to print to STDOUT.


