# pkg
**pkg** is a simple package management library written in go. It comes with **pkgfetch**, a universal package management utility. The library was originally designed to deploy mods/content for games but can be used for any purpose.

## Features
- Add repositories
- Synchronize the local package index with the repositories
- Install packages
- Upgrade packages
- Search the local package index
- List the packages to upgrade
- List the dependencies of a package

Uninstalling a package is not possible, because **pkg** does not keep track of files which have been installed by the packages.

## Repository format
The repositories are json documents, served via http(s). Below is an example of a repository index file:

```json
{
    "Packages": [
        {
            "Name": "another_tool",
            "Version": 1,
            "File": "http:\/\/localhost\/repo\/another_tool.zip",
            "Dependencies": [
                "fancy_tool"
            ]
        },
        {
            "Name": "fancy_tool",
            "Version": 1,
            "File": "http:\/\/localhost\/repo\/fancy_tool.zip",
            "Dependencies": []
        }
    ]
}
```

When updating the local package index, the package entry of a repository is ignored if there is a newer version in the local index.
