## A text file with a parseable date or time prefix

```
2012-09-01-chores.md
```

## A Heading

 \#\# filename.md


GTD so the names of things will be odd

Open every text file in a given dir
assumes that the files are prefixed with dates so that fd sorts asc
hcats the last 20 files

opens text editor and when text editor is closed it splits them back out


```bash
#!/bin/bash

source $HOME/.envrc

if [ -f "$CAP_FILE" ]; then
  exit 1
fi

fd . $CAP_DIR | tail -n 20 | hcat > $CAP_FILE

nvim $CAP_FILE -c 'cd ~/personal/00-capture/cap' -c '$'

cap.sync
```

splits the contents of a file with headings

```
#!/bin/bash

source $HOME/.envrc

if [ -f "$CAP_FILE" ];
then
  cat $CAP_FILE | hsplit $CAP_DIR && rm $CAP_FILE
fi
```

