# bookmarks-in-issues

## Overview

This GitHub Action is designed to streamline bookmark management by updating GitHub Issues created with a URL in the title, changing the issue's title to the corresponding web page's title and setting the issue's body to the URL.

## Setup

Incorporate this action into your workflow by adding the following configuration to a new yaml file in your `.github/workflows` directory:

```yml
on:
  issues:
    types: [opened]

jobs:
  bookmarks-in-issues:
    runs-on: ubuntu-latest
    steps:
      - uses: hxrxchang/bookmarks-in-issues@v0.0.4
        with:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          url: ${{ github.event.issue.title }}
          number: ${{ github.event.issue.number }}
```

## Important: Setting Write Permissions for `GITHUB_TOKEN`

To ensure this action can successfully update issues, it's crucial to verify that the GITHUB_TOKEN used has write permissions. You can do this by:

1. Go to your repository's Settings.
2. Click on Actions in the left sidebar, then select General.
3. Under Workflow permissions, check that Read and write permissions is selected. This step is essential for the action to have the necessary permissions to update issue titles and bodies.

## Benefits

- Efficiently organizes bookmarks within GitHub Issues.
- Automates the update of issue titles and bodies with webpage titles and URLs.
