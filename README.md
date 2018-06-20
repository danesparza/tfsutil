# tfsutil
Utilities for working with TFS (Team Foundation Server)

## Getting started
- Grab the [latest release](https://github.com/danesparza/tfsutil/releases/latest) and unzip `tfsutil` to a location in your path.
- Create a config file using the command `tfsutil config create`.  Save the generated text to tfsutil.yml in the same directory as the binary (or to your home directory if on a unix/linux based platform)
- Update the TFS url in the tfsutil.yml with your server information, including the default collection and project you'd like to use.  **Be sure to leave the /_apis at the end of the url**
- Create a personal access token and set it in the tfsutil.yml config file.  (Need help? [See the guide on Microsoft's site, here](https://docs.microsoft.com/en-us/vsts/accounts/use-personal-access-tokens-to-authenticate?view=vsts).)
