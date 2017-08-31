# plp

[Pl](#)uto [P](#)ackage Manager.

Doesn't really do much at the moment, but in the future, you will be
able to install, remove, and update, Pluto packages.

```
plp +package    # installs a package
plp -package    # removes a package
plp ^package    # updates a package
```

Operations can be chained (e.g. `plp +foo -bar`.)

Packages can be registered in the repository `pluto-language/packages` -
just submit a PR.
