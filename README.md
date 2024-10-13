
## Repo structure

- assets: Assets involved in the system as external git repos. Includes libraries, SDKs, repos for local reference...
- .github: CICD
- components: Source code of each custom component. Develop, install and test in isolation. 
  - docs: Low Level docs
  - install: dockerfile and scripts to install/build the component
  - tools: Custom tools
- deploy: Instructions to deploy components as part of the an E2E system including instructions to deploy opensource components which are required
- docs: High level docs
- external: Directory ignored to be able to host local git submodules of thirdparty software of interest but not extended by us.
- tools: Custom tools 

## High Level description



## Roadmap



## Contribution

A Gitflow methodology is implemented within this repo. Proceed as follows:

1. Open an Issue, label appropriately and assign to the next planned release.
2. Pull the latest develop and branch off of it with a branch or PR created from Github.
4. After the first commit, open a DRAFT PR to develop.
3. Commit and push frequently.
4. When ready, set PR as ready, request review and wait for approval.

## License

Personal Public License to ensure that the project grows and receives contributions.

```text
                    GNU GENERAL PUBLIC LICENSE
                       Version 3, 29 June 2007

 Copyright (C) 2007 Free Software Foundation, Inc. <https://fsf.org/>
 Everyone is permitted to copy and distribute verbatim copies
 of this license document, but changing it is not allowed.

```
