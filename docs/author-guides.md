# Creating authored documentation pages

Authored documentation pages are referred to as <span class="hljs-attr">guides</span>. Guides should be authored in
[Github Flavoured Markdown](https://help.github.com/articles/basic-writing-and-formatting-syntax/) (GFM) since 
using HTML will make them dependant on the theme in use at the time they were written, and therefore resistant to change.
The flexible approach is to use Github Flavoured Markdown.

## Github Flavoured Markdown

Guides written using [Github Flavoured Markdown](https://help.github.com/articles/basic-writing-and-formatting-syntax/)
have a file extension of `.md`. They can either be **global**, in which case they are located within directory
`{local_assets}/templates/guides/`, or **per-specification** and are located within
directory `{local_assets}/sections/{specification_ID}/guides/`.

If you tell Swaggerly where to find your local assets, then on startup, it will locate and build 
pages for all of your guides, maintaining the directory structure it finds below the `/guides/` directory

### Example

The Swaggerly assets example `examples/guides/assets/` contains three guides:

1. `templates/guides/markdown.md`
2. `templates/guides/level2/markdown2.md`
3. `sections/swagger-petstore/guides/PetstoreHelp.md`

The first two are global documents, available outside any specification reference page. The thrid is a guide specific to
the Petstore specification.

To build and make these guides available, pass Swaggerly the configuration parameters:

```bash
-assets-dir=<Swaggerly-source-directory>/examples/guides/assets \
-spec-dir=<Swaggerly-source-directory>/examples/specifications/petstore \
-force-specification-list=true
```

will build these three guides, making them available a `http://localhost:3123/`.

> When presenting a single specification, Swaggerly will make its root *home page* the summary of that API
and not the summary list of available specifications. Since the `markdown` and `markdown2` guides are only available
from the summary list of available specifications, they weill not be visible. The `-force-specification-list=true` option
instructs Swaggerly to show the list of available specifications, from where you can access these global guides.

When viewing the available specifications list the navigation rendered at the side of
the page will show two navigation entries:

```
> level2
  - markdown2
> markdown
```

and when viewing the [Swagger Petstore](http://localhost:3123/swagger-petstore/) specification, the side navigation
shows:

```HTTP
  PetstoreHelp
  API reference
> Everything about your Pets
> Access to Petstore orders
> Operations about user
```

By default, the side navigation will reproduce the directory structure beneath the `guides/` directory,
However there is an enforced limit of two levels in the navigation, which consequently restricts the depth of
your directory structure.

If you need a deeper directory structure, or have a file nameing convention that does not lend itself
to navigation titles, you can take control of the side navigation through [metadata](#controlling-guide-behaviour-with-metadata).

## Controlling guide behaviour with metadata

Swaggerly allows the integration of guides to be controlled with some simple metadata. This metadata is added to the
beginning of GFM files as a block of lines containing `key: value` pairs. Swaggerly will stop processing 
any metadata present at the first blank line, and consider the rest of the file as content.

This metadata allows you to control the side navigation hierarchy, grouping and page naming.

### Example

The example assets `examples/metadata/assets` demonstrate the use of metadata to control guide navigation. To run
these examples, pass the following assets configuration to Swaggerly:

```bash
-assets-dir=<Swaggerly-source-directory>/examples/guides/assets 
```

Three example global guides are provided, each demonstrating a different directory structure:

1. `examples/metadata/assets/templates/guides/markdown.md`
2. `examples/metadata/assets/templates/guides/level2/markdown2.md`
3. `examples/metadata/assets/templates/guides/level2/level3/markdown2.md`

When Swaggerly renders the documentation, it uses the embedded metadata to name, sort and structure the
navigation entries. If loading a single specification, remember to pass the `-force-specification-list=true`
option to enable the specification list summary home page, from where you will be able to see the global
documentation.


The first example file `examples/metadata/assets/templates/guides/markdown.md` contains metadata as follows:

```HTML
Navigation: Examples/A markdown example
SortOrder: 310
Note: This top section is just MetaData and gets stripped to the first blank line.

# This page was written using Git Flavoured Markdown - With metadata
...
```

The third example file `examples/metadata/assets/templates/guides/level2/level3/markdown2.md`, which is **three**
directory levels deep and thus would normally give an error, renders correctly because of it's `Navigation`
metadata.

```HTML
Navigation: Overview/Another example
SortOrder: 210
```

The `Navigation:` entry forces a page side navigation entry that is just **two** levels deep:

```
  Getting Started
> Overview
  - Another example
> Examples
  - A markdown example
```

By using metadata, the side navigation wording and structure is divorced from the underlying file naming
convention, structure and depth.

The ordering of pages within the page side navigation is controlled by the addition of 
[SortOrder](/docs/author-metadata.html#sortorder) metadata. Refer to the 
[metadata](/docs/author-metadata.html) section for further details.

!!!HIGHLIGHT!!!