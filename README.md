# plp

[Pl]()uto [P]()ackage Manager.

Doesn't really do much at the moment (since there's no source), but
in the future, you will be able to install, remove, and update,
Pluto packages.

These functions will be subcommands of `plp`:

```
plp install <package>
plp remove <package>
plp update <package>
```

You will be able to fetch the packages directory with `plp fetch`,
which will download the list of registered packages.

Packages can be registered in the repository `pluto-language/packages` -
just submit a PR.
