# Dumbsweeper

A tiny (746 bytes as of the last update), simple, and feature-less (hence "dumb", it has only the bare essentials) Minesweeper game packed into a data URL, made for [Shrink](https://shrink.hackclub.com).

Copy and paste this into your browser's address bar to try it:

```
data:text/html,<style>body,button{background:%23000;color:%23fff;font-family:monospace;display:flex;flex-direction:column;align-items:center;justify-content:center}button{height:48px;width:48px;background:gray;border:1px solid %23fff}c{display:grid;grid-template-columns:repeat(9,50px)}</style><h1>Mindsweeper, now packaged as Dumbsweeper!</h1><c id=c></c><script>c.innerHTML="<button></button>".repeat(81),B=[...c.children];for(i=0;i<10;i++)B[Math.random()*81|0].m=!0;B.map((e,t)=>{e.n=B.filter((e,n)=>Math.hypot(t%259-n%259,(t/9|0)-(n/9|0))<1.5&&B[n].m).length,e.onclick=function(){if(e.disabled)return;e.disabled=!0,e.textContent=e.m?"💣":e.n||"",!e.n&&!e.m&&B.map((e,n)=>Math.hypot(t%259-n%259,(t/9|0)-(n/9|0))<1.5&&e.onclick())}})</script>
```

## gourlgolfing (Go URL Golfing)

A tool I built to help create Dumbsweeper.

It's essentially a wrapper around existing utilities to make code golfing data URL websites easier.

To use it, install Go and run:

```bash
go install github.com/reztheperson/dumbsweeper/gourlgolfing@latest
```

- `gourlgolfing dev <folder>`: Combines `style.css`, `index.html`, and `script.js` into a single file and serves it at `localhost:8080`.
* `gourlgolfing build <folder>`: Combines `style.css`, `index.html`, and `script.js` into one file, minifies it by stripping unnecessary whitespace and comments, converts it to a data URL, and saves the output to `uri.txt`.
