---
title: "rclone genautocomplete zsh"
description: "Output zsh completion script for rclone."
slug: rclone_genautocomplete_zsh
url: /commands/rclone_genautocomplete_zsh/
# autogenerated - DO NOT EDIT, instead edit the source code in cmd/genautocomplete/zsh/ and as part of making a release run "make commanddocs"
---
# rclone genautocomplete zsh

Output zsh completion script for rclone.

## Synopsis


Generates a zsh autocompletion script for rclone.

This writes to /usr/share/zsh/vendor-completions/_rclone by default so will
probably need to be run with sudo or as root, e.g.

    sudo rclone genautocomplete zsh

Logout and login again to use the autocompletion scripts, or source
them directly

    autoload -U compinit && compinit

If you supply a command line argument the script will be written
there.

If output_file is "-", then the output will be written to stdout.


```
rclone genautocomplete zsh [output_file] [flags]
```

## Options

```
  -h, --help   help for zsh
```

See the [global flags page](/flags/) for global options not listed here.

## SEE ALSO

* [rclone genautocomplete](/commands/rclone_genautocomplete/)	 - Output completion script for a given shell.
