# git-identity

This CLI tool is a complementary `git extension` to allow better
management and integration of multiple git user profiles.

The git command line is pain to use, hard to remember, and easy
to get wrong. That's why this tool exists.


## Syntax

| Action             | Description                                         |
|:-------------------|:----------------------------------------------------|
| `create <alias>`   | creates a new alias                                 |
| `import <alias>`   | imports global git user config as a new alias       |
| `clone <alias>`    | clones a private repository with SSH key of alias   |
| `config <alias>`   | modifies current git repository to always use alias |
| `show`             | prints out available aliases                        |
| `show-key <alias>` | prints out the public SSH key for given alias       |


## Examples

```bash
# import global git/ssh config as new alias
git identity import cookiengineer;

# create an alias
git identity create another-alias;

# clone a repo via an aliases' SSH key
git identity clone cookiengineer git@github.com:cookiengineer/private-repository.git ./private-repo;

# modify alias usage inside a repo
cd ./private-repo;
git identity config another-alias;
git push origin master; # uses another-alias automatically!

# easy copy/paste of public key to web apps
git identity show-key another-alias;
```


## Installation

```bash
# clone the repository
git clone https://github.com/cookiengineer/git-identity;

# builds and installs the binary to /usr/bin/git-identity
cd ./git-identity;
bash install.sh;
```


## Filesystem Backup / Config Folders

- The compiled [main.go](/main.go) binary has to be installed as `/usr/bin/git-identity`
- `$XDG_CONFIG_HOME/git-identity` is the preferred config folder
- `$HOME/.config/git-identity` is the first fallback config folder
- `/home/$USER/.config/git-identity` is the second fallback config folder


## License

WTFPL

