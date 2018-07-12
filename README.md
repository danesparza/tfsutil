# tfsutil [![CircleCI](https://circleci.com/gh/danesparza/tfsutil.svg?style=shield)](https://circleci.com/gh/danesparza/tfsutil)
Utilities for working with [TFS (Team Foundation Server)](https://docs.microsoft.com/en-us/vsts/user-guide/?view=tfs-2018)

## Getting started
- Grab the [latest release](https://github.com/danesparza/tfsutil/releases/latest) and unzip `tfsutil` to a location in your [path](https://en.wikipedia.org/wiki/PATH_(variable)).
- **Create a config file** using the command `tfsutil config create`.  Put the generated text in a file named `tfsutil.yml` in the directory that %userprofile% points to on windows (probably the root of your user directory), or to your home directory (if on a unix/linux based platform), or in the same directory as the binary.
- **Update the TFS url** in the `tfsutil.yml` with your server information.  Also update the default collection and project you'd like to use.
- **Create a personal access token** and set it in the tfsutil.yml config file.  (Need help? [See the guide on Microsoft's site](https://docs.microsoft.com/en-us/vsts/accounts/use-personal-access-tokens-to-authenticate?view=vsts).)

### Listing variable groups
To list variable groups, execute the command:

```
tfsutil vg list
```

All variable groups for the current collection and project will be listed, along with the count of the variables in each group.

### Copying a variable group
To copy a variable group, execute the command: 

```
tfsutil vg copy "Special unicorn variables"
```

Where 'Special unicorn variables' is the name of the variable group you want to copy.  Note:  The variable group name should be surrounded with quotes.
