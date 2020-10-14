# Wiki

![Go](https://github.com/jasongauvin/wikiPattern/workflows/Go/badge.svg?branch=master)

This repository is a small training on design patterns, this iteration is about the MVC design pattern.

The project comes with a front-end that you can access through : `localhost:8000`

>Some features are presented but are not working yet: Thumbnails, Popular articles list, Likes, Tags, and SearchBar.

>Working features: HomePage, Article list, CRUD for articles, Comments on articles and export articles on CSV and XLSX formats.
---
## Install

With docker: 

Refer to [Docker-Setup](#docker-setup) to install with docker.


Without docker: 

If not, to run this project, you will need to install the following dependencies on your system:

- [go](https://golang.org/doc/install)

## Docker-Setup

![build instructions](assets/build.png)

## Branch naming convention

You branch should have a name that reflects it's purpose.

It should use the same guidelines as [COMMIT_CONVENTIONS](COMMIT_CONVENTIONS.md) (`feat`, `fix`, `build`, `perf`, `docs`), followed by an underscore (`_`) and a very quick summary of the subject in [kebab case][1].

Example: `feat_add-image-tag-database-relation`.

## Pull requests (PR)

Pull requests in this project follow two conventions, you will need to use the templates available in the [ISSUE_TEMPLATE](.github/ISSUE_TEMPLATE) folder :

- Adding a new feature should use the [FEATURE_REQUEST](.github/ISSUE_TEMPLATE/FEATURE_REQUEST.md) template.
- Reporting a bug should use the [BUG_REPORT](.github/ISSUE_TEMPLATE/BUG_REPORT.md) template.

If your pull request is still work in progress, please add "WIP: " (Work In Progress) in front of the title, therefor you inform the maintainers that your work is not done, and we can't merge it.

The naming of the PR should follow the same rules as the [COMMIT_CONVENTIONS](COMMIT_CONVENTIONS.md)

## Contributors

<table>
  <tr>
    <td align="center">
    <a href="https://github.com/jasongauvin">
      <img src="https://avatars1.githubusercontent.com/u/41618366?s=400&u=b970ed03cbb921ce1312ef86b39093e4fa0be7e3&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Jason Gauvin</b></sub>
    </a>
    </td>
    <td align="center">
    <a href="https://github.com/JackMaarek/">
      <img src="https://avatars3.githubusercontent.com/u/28316928?s=400&u=3cdfb5b0683245ad333a39cfca3a5251f3829824&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Jacques Maarek</b></sub>
    </a>
    </td>
    <td align="center">
    <a href="https://github.com/SteakBarbare">
      <img src="https://avatars2.githubusercontent.com/u/25483831?s=400&u=5316e2018489cb088c6120940df7e0b5d8d0f374&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Corto Dufour</b></sub>
    </a>
    </td>
    <td align="center">
    <a href="https://github.com/edwinvautier">
      <img src="https://avatars3.githubusercontent.com/u/35581502?s=460&u=d9096f90151f35552d9adcd57bacaee366f0aaef&v=4" width="100px;" alt=""/>
      <br />
      <sub><b>Edwin Vautier</b></sub>
    </a>
    </td>
  </tr>
</table>