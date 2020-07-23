---
title: "vim"
date: 2020-07-23T17:59:39+02:00
draft: false
---
### In VIM how to map save to Ctrl+s

Great to avoid freezing the terminal by mistake

in .vimrc
```
noremap <silent> <C-S>          :update<CR>
vnoremap <silent> <C-S>         <C-C>:update<CR>
inoremap <silent> <C-S>         <C-O>:update<CR>
```

in .bashrc or .zshrc
```
stty -ixon
```

### Help VIM frozen, don't accept input, can't close

Probably you did Ctrl+S by mistake - it freezes your terminal, type Ctrl+Q to get it going again.

