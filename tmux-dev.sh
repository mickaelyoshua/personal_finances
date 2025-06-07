#!/bin/bash

# This script sets up a tmux session for development with a specific layout and commands.
# Check if tmux is installed
if ! command -v tmux &> /dev/null; then
	echo "tmux is not installed. Please install it first."
	exit 1
fi

# Create a new tmux session
SESSION_NAME="dev"
tmux new-session -d -s $SESSION_NAME

# Split the window into panes
tmux split-window -h
tmux split-window -v
tmux select-pane -t 0

# Send commands to each pane
tmux send-keys -t 1 "make run_db" C-m
tmux send-keys -t 2 "cd app && air" C-m

# Attach to the session
tmux attach -t $SESSION_NAME