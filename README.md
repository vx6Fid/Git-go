# ggit

A minimal, Git-like version control system written in Go.  
This project is inspired by [ugit](https://www.leshenko.net/p/ugit/) and serves as a learning tool to understand the inner workings of Git from the ground up.

## Features

- Initialize a new repository
- Hash and store file contents as objects
- Read and write objects (blobs, trees, commits)
- Create commits and track history
- Basic object inspection (`cat-file`, `hash-object`, etc.)

> More features coming soon as development progresses!

## Installation

Make sure you have Go installed (version 1.18+).

```bash
git clone https://github.com/yourusername/ggit.git
cd ggit
go build -o ggit .
```

Now you can run it as 
```bash
./ggit init
```

## Usage
```bash
ggit init                    # Initialize a new repository
ggit hash-object <file>     # Hash and store a file as a blob
ggit cat-file -p <hash>        # View the content of a stored object
ggit write-tree             # Write a tree object from the working directory
ggit commit-tree <tree>     # Commit a tree with a message
```
I am implementing more of these commands.

## Project Structure
```bash
.git/       # Internal data storage (like .git/)
  objects/       # Stores object files
  refs/          # References to commits (heads, tags)
  HEAD           # Points to the current branch ref
```

## Credits
Inspired by [ugit](https://www.leshenko.net/p/ugit/) and Git's own internals.

