![donuts-are-good's followers](https://img.shields.io/github/followers/donuts-are-good?&color=555&style=for-the-badge&label=followers) ![donuts-are-good's stars](https://img.shields.io/github/stars/donuts-are-good?affiliations=OWNER%2CCOLLABORATOR&color=555&style=for-the-badge) ![donuts-are-good's visitors](https://komarev.com/ghpvc/?username=donuts-are-good&color=555555&style=for-the-badge&label=visitors)

# timestamp

a tool to view and modify timestamps


## usage

you can view timestamps and/or rewrite them, so the `--modify` flag is optional

```bash
timestamp [--modify] /path/to/file
```

without the `--modify` flag, `timestamp` will display the timestamps of the specified file. with the `--modify` flag, you'll be asked interactively for new timestamps and it'll confirm before any changes are made.

## examples


1. show timestamps

to only show the timestamps of a file, run `timestamp` with the path to the file as an argument:

```shell
$ timestamp /path/to/file.txt
```
this will show the current, created, and modified timestamps of the file.

2. modify timestamps
to modify the timestamps of a file, use the `--modify` flag:

```bash
$ timestamp --modify /path/to/file.txt
```
this will start an interactive prompt where you can enter new timestamps. 

here's what the prompt process looks like:

```bash
enter the new created timestamp
(format: RFC3339, ex: '2006-01-02T15:04:05Z07:00')
or leave it blank to use the current time): <put your timestamp here>

enter the new modified timestamp
(format: RFC3339, ex: '2006-01-02T15:04:05Z07:00')
or leave it blank to use the current time): <put your other timestamp here>

You entered:
New created timestamp: 2023-05-12T10:30:00Z
New modified timestamp: 2023-05-12T11:30:00Z
Are you sure you want to proceed with these changes? (yes/no): <if you just press enter, it should quit>
```
enter the new timestamps in the example format or leave it blank to use the current time, then confirm it by typing `yes`.

that's it :) 

## note

i couldn't find a way to get file creation time in mac and unix systems. if you can do that, please pr it or make an issue if you have time.

## license

MIT License 2023 donuts-are-good, for more info see license.md
