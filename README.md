# A Lavish HTML Templating System for Go

For when [Plush](https://github.com/gobuffalo/plush) just isn't _nice_ enough.

Have you ever wished that you could write a monolithic web application in Go,
with server-side rendering, but you kinda like JSX for HTML templates?

Finally, the day has come* where this can be a reality!

*if you count alpha-quality software without API stability guarantees

Lavish is built on the incredible foundation laid by
[esbuild](https://github.com/evanw/esbuild)
and
[goja](https://github.com/dop251/goja)
to compile JSX to vanilla Javascript and execute it in a pure Go runtime.

It's faster than you might think - not as fast as Plush, of course, but only a
few thousand times slower! Surely those handfuls of milliseconds are worth the
more comfortable developer experience.

## Rendering Engines

Lavish comes bundled with Preact v10 to turn the JSX into an HTML string.

In theory, any JSX framework which also ships a server-side rendering package
could be used instead - although you can't have any client-side state, anyway,
so what else could you want? Maybe there's a runtime that's even smaller after
tree shaking?

## Reminder that Lavish is currently in an alpha state

I got the basic idea working in a few hours late one evening and then spent a
Saturday setting up this barely presentable package. Use this package at your
own risk.
