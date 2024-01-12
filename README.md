## This is my homepage repo
## Bug list
* Don't support Chinese Tags and Categories.
> Using github action workflow can generate static file correctly.But failed to deploy by vercel-cli as tags or categories has Chinese charactor.

[Reason] The vercel cli cause this. the error ocurrs when `vercel deploy --prebulit`.but `vercel deploy` is ok
