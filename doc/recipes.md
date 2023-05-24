```bash
#!/bin/bash

source $HOME/.envrc

if [ -f "$CAP_FILE" ]; then
  exit 1
fi

cap.dedup-links

fd . $CAP_DIR | tail -n 20 | hcat > $CAP_FILE

nvim $CAP_FILE -c 'cd ~/personal/00-capture/cap' -c '$'

cap.sync
```
