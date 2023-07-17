# git-identity

This CLI tool is a complementary `git extension` to allow better
management and integration of multiple git user profiles.

The git command line is pain to use, hard to remember, and easy
to get wrong. That's why this tool exists.


## Features

- `git identity create <alias>` creates a new identity/alias with username, email address, and SSH key pairs.
- `git identity import <alias>` creates a new identity/alias from existing global SSH/git config.
- `git identity config <alias>` configures the current repository to use the given alias.
- `git identity clone <alias> git@github.com:private/repo.git` uses the given alias SSH key pairs to clone a private repository.
- `git identity show-key <alias>` prints the public key for easy copy/paste into your web app of choice.


## Installation

```bash
# clone the repository
git clone https://github.com/cookiengineer/git-identity;

# builds and installs the binary to /usr/bin/git-identity
cd ./git-identity;
bash install.sh;
```

## License

WTFPL

