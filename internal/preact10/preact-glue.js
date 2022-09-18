import render from 'preact-render-to-string';
import { h, Fragment } from 'preact';

// Set these variables in the global scope, busting through the bundle encapsulation.
_render = render;
_h = h;
_Fragment = Fragment;
