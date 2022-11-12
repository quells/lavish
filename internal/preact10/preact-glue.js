import { render as _render } from 'preact-render-to-string';
import { h as _h, Fragment as _Fragment } from 'preact';

// Set these variables in the global scope, busting through the bundle encapsulation.
_preact10_render = _render;
h = _h;
Fragment = _Fragment;
