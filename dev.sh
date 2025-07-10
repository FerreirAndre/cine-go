#!/bin/bash

session="dev"

tmux new-session -d -s $session

tmux rename-window -t 1 'nvim'
tmux send-keys -t 'nvim' 'nvim .' C-m

tmux new-window -t $session:2 -n 'server'
tmux new-window -t $session:3 -n 'terminal'

#tmux send-keys -t 'server' 'dotnet run --project Cine.Resenha.Api' C-m
#tmux send-keys -t 'server-front' 'cd Cine.Resenha.UI' C-m 'npm run dev -- --open' C-m


tmux attach-session -t $session:1
