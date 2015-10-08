# hafen RFC

## Scenario: Installing some tool

When I run this:

```bash
hafen install docker-compose
```

Then the service finds information about `docker-compose` in registry, such as:

- Where and how to download the code (`tar`, `zip` and `git` url templates with
  free `version` variable);
- How to build code into a fully self-contained `docker` container (Github link
  of form: `{https://github.com/heimathafen/recipes; /docker-compose}`);
- How to run the final docker image (basically a script, that would imitate an
  actual tool and abstract the fact that it is run inside of docker container,
  i.e.: Github link of form `{https://github.com/heimathafen/recipes;
  /docker-compose/run.sh}`);

And the service builds the docker image for `docker-compose` of last released
version and pushes it to `hub.docker.com` as
`heimathafen/docker-compose:{{version}}` unless it was already built for this
version (or commit SHA);

And the image gets downloaded to my machine;

And the script to run the final docker image gets installed at
`/opt/hafen/docker-compose/bin/docker-compose` and linked to
`/usr/local/bin/docker-compose`,

- Given I already have something in `/usr/local/bin/docker-compose`;
- And it is not a link to `/opt/hafen/*/bin/*`;
- Then `hafen` asks me a (yes/no) question if I want to override it;

When I run `docker-compose --version`;

Then I see a latest released version of `docker-compose`.
