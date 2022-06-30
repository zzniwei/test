---
# rclone batchcopy

BatchCopy dir from source dir to dest dir with edited filelist.

## Synopsis


Batch Copy the source dir to the dest dir with edited filelist. the 
content of the filelist is pair like 'sourdir,destdir' line by line 
in the filelist, and terminated by '\n', seperated by ',' character,
the space can not be supported  in the filelist. 


For example

    rclone batchcopy test.txt

Let's say there are two items in test.txt

    sourcepath1,destpath1
    sourcepath2,destpath2

This copies all contents in sourcepath1 to destpath1 and sourcepath2 to destpath2


```
rclone copy filelist [flags]
```

## Options

```
      --batch-copy-limit  this flag is a number limit batchcopy jobs per time, defaut value is 5000
  -h, --help                    help for batchcopy
```

See the [global flags page](/flags/) for global options not listed here.

## SEE ALSO

* [rclone](/commands/rclone/)	 - Show help for rclone commands, flags and backends.

