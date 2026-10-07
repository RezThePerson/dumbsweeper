// add the cells
c.innerHTML = '<d></d>'.repeat(81)
let b = [...c.children]

// set the mines
for (i = 0; i < 10; i++) b[Math.random() * 81 | 0].m = true
