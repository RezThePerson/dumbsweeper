// add cells
c.innerHTML = '<button></button>'.repeat(81)
B = [...c.children]

// set the mines
for (i = 0; i < 10; i++) B[Math.random() * 81 | 0].m = true

// Game loop
B.map((e, i) => {
	// Count neighbors
	e.n = B.filter((_, j) => Math.hypot(i % 9 - j % 9, (i / 9 | 0) - (j / 9 | 0)) < 1.5 && B[j].m).length

	// Self-calling cascade
	e.onclick = function F() {
		if (e.disabled) return
		e.disabled = true
		e.textContent = e.m ? '💣' : e.n || ''

		// Cascade: if count is 0, trigger click on all neighbors
		if (!e.n && !e.m) {
			B.map((n, j) => Math.hypot(i % 9 - j % 9, (i / 9 | 0) - (j / 9 | 0)) < 1.5 && n.onclick())
		}
	}
})
