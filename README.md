[![GitHub release](https://img.shields.io/github/release/Harry-Moore-dev/bump-tf-values.svg?style=flat-square)](https://github.com/Harry-Moore-dev/bump-tf-values/releases/latest)


# bump-tf-values

Github action to update (bump) the value for a Terraform local within a specified Terraform file formatted with the standard syntax.

## Usage

This action can be used by referencing the version tag in the workflow file.
```yaml
uses: Harry-Moore-dev/bump-tf-values@v0.1.0-alpha
```
Or using the preferred way for faster execution from the lightweight prebuilt docker image published with the same version tag or the 'latest' tag.
```yaml
uses: docker://ghcr.io/harry-moore-dev/bump-tf-values:v0.1.0-alpha
```

### Examples
#### Basic Usage

Example usage triggered when publishing a release to modify the local 'code_version' within the submodule 'module/main.tf' to be set as the value of the tag that has just been published. This example uses the 'actions/checkout' action to checkout the repository and then the 'peter-evans/create-pull-request' action to raise a pull request with the changes.

```yaml
name: bump-tf
on:
  release:
    types: [published]
jobs:
  bump-tf:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Source
        uses: actions/checkout@v3
      - name: Bump tf value
        uses: docker://ghcr.io/harry-moore-dev/bump-tf-values:latest
        with:
          filepath: module/main.tf
          varname: code_version
          value: "${{ github.ref_name }}"
      - name: Create Pull Request
        uses: peter-evans/create-pull-request@v5
```

#### Usage with Signed Commits

For a repository that enforces signed commits, I recommend creating a bot account and following the guidance in [this post](https://httgp.com/signing-commits-in-github-actions/) to create and store a GPG key as a repository secret. This can then be imported using the 'crazy-max/ghaction-import-gpg' action.

```yaml
name: bump-tf
on:
  release:
    types: [published]
jobs:
  bump-tf:
    env:
      FILEPATH: main.tf
      VARNAME: code_version
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Source
        uses: actions/checkout@v3
        with:
          repository: Harry-Moore-dev/some-other-repo
          token: ${{ secrets.GH_PAT }}
      - name: Bump tf value
        uses: docker://ghcr.io/harry-moore-dev/bump-tf-values:latest
        with:
          filepath: ${{ env.FILEPATH }}
          varname: ${{ env.VARNAME }}
          value: "${{ github.ref_name }}"
      - name: Import GPG key
        uses: crazy-max/ghaction-import-gpg@v5
        with:
          gpg_private_key: ${{ secrets.BOT_GPG }}
          git_user_signingkey: true
          git_commit_gpgsign: true
      - name: Git Signed Commit Changes
        run: |
          git add .
          git commit -S -m "Version bump ${{ env.VARNAME }} to ${{ github.ref_name }}"
      - name: Create Pull Request
        uses: peter-evans/create-pull-request@v5
        with:
          token: ${{ secrets.GH_PAT }}
          title: "Bump pinned version of ${{ env.VARNAME }}"
          branch: "bump-tf/${{ env.VARNAME }}"
```

This example also demonstrates checking out and raising a pull request on a repository outside of the one running the action by using a Github personal access token or Github app token as a repository secret with the correct scopes.

## Inputs

This action takes the following inputs:

| Input      | Description                                                                       | Type   | Usage      |
| ---------- | --------------------------------------------------------------------------------- | ------ | ---------- |
| 'filepath' | Filepath containing Terraform file to be modified. If blank defaults to 'main.tf' | String | Optional   |
| 'varname'  | Name of the local to be modified                                                  | String | \*Required |
| 'value'    | New value to be assigned to the local                                             | String | \*Required |

## Build Instructions

To build your own version of the docker container run `docker build -f -t <\username/repo> Publishing.dockerfile .'

## Future Improvments

* Support multiple variables & values as inputs
* Support other block types than 'locals' such as variable defaults
