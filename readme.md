# Dumbsweeper

a really small (746 bytes as of last update), simple, and non-feature packed (thats where dumb comes from it has only the essential part) mindsweeper game in a data url made for [shrink](shrink.hackclub.com)

copy paste this to try it:
```
data:text/html,<style>body,button{background:%23000;color:%23fff;font-family:monospace;display:flex;flex-direction:column;align-items:center;justify-content:center}button{height:48px;width:48px;background:gray;border:1px solid %23fff}c{display:grid;grid-template-columns:repeat(9,50px)}</style><h1>Mindsweeper, now packaged as Dumbsweeper!</h1><c id=c></c><script>c.innerHTML="<button></button>".repeat(81),B=[...c.children];for(i=0;i<10;i++)B[Math.random()*81|0].m=!0;B.map((e,t)=>{e.n=B.filter((e,n)=>Math.hypot(t%259-n%259,(t/9|0)-(n/9|0))<1.5&&B[n].m).length,e.onclick=function(){if(e.disabled)return;e.disabled=!0,e.textContent=e.m?"💣":e.n||"",!e.n&&!e.m&&B.map((e,n)=>Math.hypot(t%259-n%259,(t/9|0)-(n/9|0))<1.5&&e.onclick())}})</script>
```

## gourlgolfing (Go URL Golfing)

a tool i made to help me make dumbsweeper

aka warper around a bunch of existing things to make data url website golfing easier.

to use it install go then run: `go install github.com/reztheperson/dumbsweeper/gourlgolfing`.
- `gourlgolfing dev <folder>`: takes the style.css, index.html and script.js puts them together in one file and serves that file at :8080
- `gourlgolfing build <folder>`: takes the style.css, index.html and script.js puts them together in one file, minimized/strips it off useless whitespaces and comments, converts to data url and puts output in uri.txt
