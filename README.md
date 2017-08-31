# plp

[Pl](#)uto [P](#)ackage Manager.

An easy way to install, remove, and update Pluto packages.

```
plp +package    # installs a package
plp -package    # removes a package
plp ^package    # updates a package
```

Operations can be chained (e.g. `plp +foo -bar`.)

You can list all packages available for installation:

```
plp list
```

Packages can be registered in the repository `pluto-language/packages` -
just submit a PR.
