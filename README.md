# jsonfmt

There are many `jsonfmt` utilities written in Go. This is mine.

It is designed to be called from an editor (in my case, [Kakoune](https://kakoune.org)), which should pipe unformatted JSON into `jsonfmt` via `stdin`. Formatted JSON will be output to `stdout`.

It does not read files from disk, nor save formatted JSON to files - it **just** reads from `os.Stdin` and outputs to `os.Stdout`.
