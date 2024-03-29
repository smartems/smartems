# Documentation style guide

This style guide applies to all documentation created for smartEMS products.

## Contributing

The *Documentation style guide* is a living document. Add to it whenever a style decision is made or a question is answered regarding style, grammar, or word choice.

## Published guides

For all items not covered in this guide, refer to the [Microsoft Style Guide](https://docs.microsoft.com/en-us/style-guide/welcome/) and the [Chicago Manual of Style](https://www.chicagomanualofstyle.org/home.html).

## Spelling

The [codespell](https://github.com/codespell-project/codespell) tool is run for every change to catch common misspellings.

## smartEMS-specific style

The following sections provide general guidelines on topics specific to smartEMS documentation. Note that for the most part, these are *guidelines*, not rigid rules. If you have questions, ask in the #docs channel of smartEMS Slack.

### General

* Use active voice. Avoid passive voice.
  - Passive: The heatmap visualization is displayed.
  - Active: smartEMS displays the heatmap visualization.
* Write in the imperative second person. Examples: You can write a query. Click the panel. Close the window.
* Write in present tense.
  - Not: The panel will open.
  - Use: The panel opens. smartEMS opens the panel.
* Do not use an ampersand (&) as an abbreviation for "and." 
  - **Exceptions:** If an ampersand is used in the smartEMS UI, then match the UI.

### File naming conventions

- Files that are displayed in the help system should have names that are all lowercase, no spaces. Use hyphens instead of spaces. Example: glossary.md
- Documentation file names should match the title. **Note:** This only applies to new files at this time. Do not change the names of older files unless directed to do so.
- Internal reference file names should be all uppercase except the file extension. Example: CONTRIBUTING.md

### Headings

* Write headings in sentence case, not title case.
  - This is sentence case
  - This Is Title Case
* Task topic headings start with a verb.
  - Write a query. Create a dashboard.
* Concept and reference topic headings should be nouns or gerunds. Examples: Contributing to docs, Visualizations, Style guide
* Avoid following one heading with another heading.
* Avoid skipping heading levels. For example, an h1 should be followed by an h2 rather than an h3.
* Avoid having just one lower-level heading. For example, h1, h2, h2, h3, h3, h2 is a good order. Do no go h1, h2, h3, h2, h3, h2.
* Don't include parenthetical words like (Important!) in headings.

### Images

* Preferred format is .png
* File extension should be all lowercase.
* Preferred DPI is 72.
* Assume all graphics will be exclusively viewed on the web.
* Maximum image size is 3840px X 2160px.
* Screenshots should be readable, but not too large.

### Capitalization

* smartEMS, Loki, and Prometheus are always capitalized unless part of a code block.
* API names are always Title Case, followed by "API"—for example, "Dashboard Permissions API"
* Git is always capitalized, unless part of a code block.
* Abbreviations are always capitalized (such as API, HTTP, ID, JSON, SQL, or URL) unless they are part of a code block.
* Menu and submenu titles always use sentence case: capitalize the first word, and lowercase the rest.
  - "Dashboards" when referring to the submenu title.
  - "Keyboard shortcuts" when referring to the submenu topic.
* Generic and plural versions are always lowercase.
  - Lowercase "dashboard" when referring to a dashboard generally.
  - Lowercase "dashboards" when referring to multiple dashboards.
* **Exceptions:** If a term is lowercased in the smartEMS UI, then match the UI.

### Links and references

When referencing another document, use "Refer to" rather than alternatives such as "See" or "Check out."

Always give the reader some idea of what to expect in the reference. Avoid blind references, such as, "Refer to [this file](link)."

When possible, use the exact title of the page or section you are linking to as the link text.

**Example**
* Refer to the [Documentation style guide](documentation-style-guide.md) for information about word usage and capitalization guidelines.

### Word usage

smartEMS products has some words, abbreviations, and slang particular to this discourse community.

#### data source

Two words, not one

**Exceptions:**
* "datasource" used as an identifier
* "datasource" in a URL
* "Open source" should be hyphenated when used as an adjective, e.g. *open-source software*. The open form should be preferred when used as a noun, e.g. *smartEMS is open source*.
* Use "data source" instead of "datasource" unless used as an identifier, in code, or as part of a URL.
* Spell out "repository" and avoid the shorter "repo."
* Use "Unix" as the preferred spelling (as opposed to "UNIX", or "unix") when referring to the family of operating systems.

#### display (verb)

*Display* is a transitive verb, which means it always needs a direct object.
* Correct, active voice: smartEMS displays your list of active alarms.
* Correct, but passive voice: Your list of active alarms is displayed.
* Incorrect: The list of active alarms displays.

#### open source, open-source

Hyphenate when used as an adjective. For example: _open-source software._ 

Do not hyphenate when it is used as a noun. For example: _Open source is the best way to develop software._

#### setup, set up

Two words if used as a verb, one word if used as a noun.

**Examples**

* Set up the workspace.
* Initial setup might take five minutes.
