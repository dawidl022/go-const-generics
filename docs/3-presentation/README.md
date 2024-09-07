# Presentation

This presentation uses [Reveal.js](https://revealjs.com/) for rendering. The
contents of the presentation itself are written in Markdown
([`presentation.md`](./presentation.md)), however [Reveal.js specific
KaTeX](https://revealjs.com/math/#markdown) syntax is used for rendering formal
rules.

## Running locally

Ensure you have [Node.js](https://nodejs.org/) and
[Yarn](https://classic.yarnpkg.com/) installed.

To install dependencies, run from the root of this directory:

```bash
yarn
```

To run the dev server, run:

```bash
yarn dev
```

This server will serve the presentation at the address shown in the output of
the above command.

Note: The KaTeX script is loaded dynamically, so an internet connection is
required to render the formal rules.

## Printing to PDF

If vite is serving at localhost:5173, then the URL to get the printable
version of the presentation is:

http://localhost:5173/?print-pdf

You can now use the regular Ctrl+P for printing webpages in your browser. It is
recommended to print using Chrome.
