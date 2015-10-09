# hafen RFC

## Scenario: Installing some tool

When I run this:

```bash
hafen install docker-compose
```

Then the service finds information about `docker-compose` in registry, such as:

- Where and how to download the code (`git` url for now);

And the service builds the docker image for `docker-compose` of last released
version by using `docker` context at `{https://github.com/heimathafen/library;
/docker-compose/build}` and pushes it to `hub.docker.com` as
`heimathafen/docker-compose:{{version}}` unless it was already built for this
version (or commit SHA);

And the image gets downloaded to my machine;

And the script to run the final docker image gets installed at
`$HOME/.hafen/lib/docker-compose/{{version}}/bin/docker-compose` from
`{https://github.com/heimathafen/library; /docker-compose/run}` and linked to
`$HOME/.hafen/bin/docker-compose`,

When I run `docker-compose --version`;

Then I see a latest released version of `docker-compose`, version of `hafen`,
version of `docker-compose` `hafen` build configuration and links to the
project and issue tracker.
