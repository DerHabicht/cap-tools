# CAP Tools

A clearinghouse of everything I've written for [Civil Air Patrol](https://gocivilairpatrol.com).

## Repository Overview

This is a massive monorepo of everything I've written for a CAP context. It contains three kinds of projects:

- Tools/utilities intended to be used locally on a workstation.
- Web apps that are useful for squadrons and Wing Activities.
- Libraries that define common functionality for the above.

A more comprehensive breakdown is given in the next section.

### Local Utility Projects

#### CAPMC

### Web App Projects

#### CAP/A5—Activities

#### CAP/A5—Units

### Library Projects

#### Members

#### Units

## Guidelines

The key words "MUST", "MUST NOT", "REQUIRED", "SHALL", "SHALL NOT", "SHOULD", "SHOULD NOT", "RECOMMENDED",  "MAY", and "
OPTIONAL" in this section are to be interpreted as described in
RFCs [2119](https://datatracker.ietf.org/doc/html/rfc2119) and [8174](https://datatracker.ietf.org/doc/html/rfc8174).

### Version Numbers

Each project is independently versioned IAW the [Semantic Versioning](https://semver.org/) specification. Yes, this can
be potentially thorny in a monorepo, but I've done this to make all the breaking changes to the library projects easily
discoverable through regression testing on their consumer projects.

### Commits

Commit messages in the `main` branch MUST follow
the [Conventional Commit](https://www.conventionalcommits.org/en/v1.0.0/#specification) specification. All commits in
`main` MUST be signed. Topic branches SHOULD be rebased against `main` regularly, and MUST be rebased before opening a 
PR. While it's probably not a bad idea to follow these rules within topic branches, deviations are OK since 

### Changelogs

Each project in this repo SHOULD have its own changelog which MUST follow
the [Keep a Changelog](https://keepachangelog.com/en/1.1.0/) specification.