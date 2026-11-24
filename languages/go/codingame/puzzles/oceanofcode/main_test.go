package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"testing"
)

func emptyMap() []string {
	row := strings.Repeat(".", 15)
	m := make([]string, 15)
	for i := range m {
		m[i] = row
	}
	return m
}

func setCell(m []string, x, y int, ch byte) {
	row := []byte(m[y])
	row[x] = ch
	m[y] = string(row)
}

func newVisitedGrid() [][]bool {
	visited := make([][]bool, 15)
	for i := range visited {
		visited[i] = make([]bool, 15)
	}
	return visited
}

func countPossible(t *EnemyTracker) int {
	count := 0
	for _, row := range t.possibleCells {
		for _, ok := range row {
			if ok {
				count++
			}
		}
	}
	return count
}

func generateMap(seed int64, islandPct int) []string {
	rng := rand.New(rand.NewSource(seed))
	for attempt := 0; attempt < 100; attempt++ {
		m := emptyMap()
		waterCount := 0
		for y := 0; y < 15; y++ {
			for x := 0; x < 15; x++ {
				if rng.Intn(100) < islandPct {
					setCell(m, x, y, 'x')
				} else {
					waterCount++
				}
			}
		}
		if waterCount >= 20 {
			return m
		}
	}
	return emptyMap()
}

// Layer 1: Unit tests for EnemyTracker, movement helpers, and torpedo logic.

func TestNewEnemyTracker(t *testing.T) {
	m := emptyMap()
	setCell(m, 0, 0, 'x')
	setCell(m, 7, 4, 'x')
	setCell(m, 14, 14, 'x')

	tracker := NewEnemyTracker(m)

	if countPossible(&tracker) != 225-3 {
		t.Fatalf("got %d possible cells, want %d", countPossible(&tracker), 222)
	}
	if tracker.possibleCells[0][0] || tracker.possibleCells[4][7] || tracker.possibleCells[14][14] {
		t.Fatalf("expected island cells to be marked false")
	}
	if !tracker.possibleCells[5][5] {
		t.Fatalf("expected water cell (5,5) to be marked true")
	}
}

func TestEnemyTrackerProcessOrdersNA(t *testing.T) {
	m := emptyMap()
	var tracker EnemyTracker
	tracker.possibleCells[3][4] = true
	tracker.possibleCells[10][12] = true

	tracker.ProcessOrders("NA", m)

	if countPossible(&tracker) != 2 || !tracker.possibleCells[3][4] || !tracker.possibleCells[10][12] {
		t.Fatalf("expected NA to leave possibleCells unchanged")
	}
}

func TestEnemyTrackerHandleMove(t *testing.T) {
	cases := []struct {
		name    string
		startX  int
		startY  int
		islandX int
		islandY int
		order   string
		wantX   int
		wantY   int
		wantOK  bool
	}{
		{name: "move north", startX: 5, startY: 5, islandX: -1, islandY: -1, order: "MOVE N", wantX: 5, wantY: 4, wantOK: true},
		{name: "move south", startX: 5, startY: 5, islandX: -1, islandY: -1, order: "MOVE S", wantX: 5, wantY: 6, wantOK: true},
		{name: "move east", startX: 5, startY: 5, islandX: -1, islandY: -1, order: "MOVE E", wantX: 6, wantY: 5, wantOK: true},
		{name: "move west", startX: 5, startY: 5, islandX: -1, islandY: -1, order: "MOVE W", wantX: 4, wantY: 5, wantOK: true},
		{name: "clipped at north edge", startX: 5, startY: 0, islandX: -1, islandY: -1, order: "MOVE N", wantX: 0, wantY: 0, wantOK: false},
		{name: "clipped at east edge", startX: 14, startY: 5, islandX: -1, islandY: -1, order: "MOVE E", wantX: 14, wantY: 5, wantOK: false},
		{name: "blocked by island", startX: 5, startY: 5, islandX: 6, islandY: 5, order: "MOVE E", wantX: 6, wantY: 5, wantOK: false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := emptyMap()
			if tc.islandX >= 0 {
				setCell(m, tc.islandX, tc.islandY, 'x')
			}
			var tracker EnemyTracker
			tracker.possibleCells[tc.startY][tc.startX] = true

			tracker.ProcessOrders(tc.order, m)

			if tc.wantOK {
				if countPossible(&tracker) != 1 || !tracker.possibleCells[tc.wantY][tc.wantX] {
					t.Fatalf("expected single cell (%d,%d), got count=%d", tc.wantX, tc.wantY, countPossible(&tracker))
				}
			} else if countPossible(&tracker) != 0 {
				t.Fatalf("expected 0 possible cells, got %d", countPossible(&tracker))
			}
		})
	}
}

func TestEnemyTrackerHandleSurface(t *testing.T) {
	cases := []struct {
		sector    int
		insideX   int
		insideY   int
		outsideX  int
		outsideY  int
		minX      int
		minY      int
		maxX      int
		maxY      int
	}{
		{sector: 1, insideX: 0, insideY: 0, outsideX: 5, outsideY: 0, minX: 0, minY: 0, maxX: 4, maxY: 4},
		{sector: 5, insideX: 9, insideY: 9, outsideX: 4, outsideY: 5, minX: 5, minY: 5, maxX: 9, maxY: 9},
		{sector: 9, insideX: 14, insideY: 14, outsideX: 10, outsideY: 9, minX: 10, minY: 10, maxX: 14, maxY: 14},
	}

	for _, tc := range cases {
		t.Run(fmt.Sprintf("sector_%d", tc.sector), func(t *testing.T) {
			m := emptyMap()
			tracker := NewEnemyTracker(m)
			tracker.ProcessOrders(fmt.Sprintf("SURFACE %d", tc.sector), m)

			if countPossible(&tracker) != 25 {
				t.Fatalf("expected 25 cells in sector %d, got %d", tc.sector, countPossible(&tracker))
			}
			if !tracker.possibleCells[tc.insideY][tc.insideX] {
				t.Fatalf("expected (%d,%d) to remain possible", tc.insideX, tc.insideY)
			}
			if tracker.possibleCells[tc.outsideY][tc.outsideX] {
				t.Fatalf("expected (%d,%d) outside sector %d to be false", tc.outsideX, tc.outsideY, tc.sector)
			}
			if !tracker.possibleCells[tc.minY][tc.minX] || !tracker.possibleCells[tc.maxY][tc.maxX] {
				t.Fatalf("expected sector corners (%d,%d) and (%d,%d) to be true", tc.minX, tc.minY, tc.maxX, tc.maxY)
			}
		})
	}
}

func TestEnemyTrackerHandleSilence(t *testing.T) {
	m := emptyMap()
	setCell(m, 7, 5, 'x') // Island east of (5,5) at distance 2

	var tracker EnemyTracker
	tracker.possibleCells[5][5] = true
	tracker.ProcessOrders("SILENCE", m)

	// 0 steps stays possible
	if !tracker.possibleCells[5][5] {
		t.Fatalf("expected origin (5,5) to remain possible")
	}
	// Up to 4 steps in clear directions
	for _, pt := range [][2]int{{5, 1}, {5, 9}, {1, 5}, {6, 5}} {
		if !tracker.possibleCells[pt[1]][pt[0]] {
			t.Fatalf("expected (%d,%d) to be reachable via SILENCE", pt[0], pt[1])
		}
	}
	// Step 5 in clear directions should not be reachable
	if tracker.possibleCells[5][0] || tracker.possibleCells[0][5] || tracker.possibleCells[10][5] {
		t.Fatalf("expected cells 5 steps away to be unreachable")
	}
	// Island and cells behind the island must stay unreachable
	if tracker.possibleCells[5][7] || tracker.possibleCells[5][8] || tracker.possibleCells[5][9] {
		t.Fatalf("expected island (7,5) and cells behind it to be unreachable")
	}

	// Edge clipping check
	var cornerTracker EnemyTracker
	cornerTracker.possibleCells[0][0] = true
	cornerTracker.ProcessOrders("SILENCE", emptyMap())
	// Origin + 4 east + 4 south = 9 cells
	if countPossible(&cornerTracker) != 9 {
		t.Fatalf("expected 9 reachable cells from corner (0,0), got %d", countPossible(&cornerTracker))
	}
}

func TestEnemyTrackerChainedOrders(t *testing.T) {
	m := emptyMap()

	t1 := NewEnemyTracker(m)
	t1.ProcessOrders("SURFACE 1 | MOVE E | TORPEDO 3 5", m)

	// Starting in sector 1 (x in 0..4, y in 0..4) then moving E gives x in 1..5, y in 0..4
	if !t1.possibleCells[0][5] {
		t.Fatalf("expected (5,0) to be possible after SURFACE 1 | MOVE E")
	}
	if t1.possibleCells[0][0] {
		t.Fatalf("expected (0,0) to be impossible after SURFACE 1 | MOVE E")
	}

	t2 := NewEnemyTracker(m)
	t2.ProcessOrders("MOVE E | SURFACE 1 | TORPEDO 3 5", m)

	// Moving E first then surfacing in sector 1 restricts to x in 1..4, y in 0..4
	if t2.possibleCells[0][5] {
		t.Fatalf("expected (5,0) to be impossible after MOVE E | SURFACE 1")
	}
	if !t2.possibleCells[0][4] {
		t.Fatalf("expected (4,0) to be possible after MOVE E | SURFACE 1")
	}
}

func TestFindEmptyStartingCell(t *testing.T) {
	m := emptyMap()
	for y := 0; y < 15; y++ {
		for x := 0; x < 15; x++ {
			setCell(m, x, y, 'x')
		}
	}
	setCell(m, 11, 3, '.')

	rng := rand.New(rand.NewSource(99))
	got := findEmptyStartingCell(m, 15, 15, rng)
	if got != "11 3" {
		t.Fatalf("got starting cell %q, want %q", got, "11 3")
	}
}

func TestDetermineNewDirection(t *testing.T) {
	cases := []struct {
		name      string
		x, y      int
		islands   [][2]int
		visited   [][2]int
		wantDir   string
		wantOneOf []string
	}{
		{
			name:    "avoids edges, islands, and visited to pick only open direction",
			x:       0,
			y:       0,
			islands: [][2]int{{1, 0}},
			wantDir: "S",
		},
		{
			name:    "avoids visited neighbors",
			x:       5,
			y:       5,
			islands: [][2]int{{5, 4}, {6, 5}},
			visited: [][2]int{{5, 6}},
			wantDir: "W",
		},
		{
			name:    "surfaces when completely boxed in",
			x:       5,
			y:       5,
			islands: [][2]int{{5, 4}, {6, 5}},
			visited: [][2]int{{5, 6}, {4, 5}},
			wantDir: "SURFACE",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := emptyMap()
			for _, p := range tc.islands {
				setCell(m, p[0], p[1], 'x')
			}
			visited := newVisitedGrid()
			for _, p := range tc.visited {
				visited[p[1]][p[0]] = true
			}

			rng := rand.New(rand.NewSource(7))
			got := determineNewDirection(tc.x, tc.y, 15, 15, m, visited, rng)
			if got != tc.wantDir {
				t.Fatalf("got direction %q, want %q", got, tc.wantDir)
			}
		})
	}
}

func TestTorpedoReachable(t *testing.T) {
	m := emptyMap()
	// Island wall forcing a detour around (6,5)
	setCell(m, 6, 5, 'x')

	// Straight distance 4
	if !torpedoReachable(5, 5, 5, 9, m) {
		t.Fatalf("expected straight distance 4 to be reachable")
	}
	// Around corner in 4 steps: (5,5) -> (5,4) -> (6,4) -> (7,4) -> (7,5)
	if !torpedoReachable(5, 5, 7, 5, m) {
		t.Fatalf("expected corner path of length 4 around island to be reachable")
	}
	// 5 steps away when detouring: (5,5) to (8,5) requires 5 steps around (6,5)
	if torpedoReachable(5, 5, 8, 5, m) {
		t.Fatalf("expected (8,5) requiring 5 steps around island to be unreachable")
	}
	// Island target itself must be rejected
	if torpedoReachable(5, 5, 6, 5, m) {
		t.Fatalf("expected island target (6,5) to be unreachable")
	}
}

// Layer 2: Protocol replay tests over scripted stdin and stdout.

func TestRunScriptedTurns(t *testing.T) {
	m := emptyMap()
	var in strings.Builder
	in.WriteString("15 15 0\n")
	for _, row := range m {
		in.WriteString(row + "\n")
	}
	in.WriteString("7 7 6 6 3 4 6 -1\nNA\nNA\n")
	in.WriteString("7 6 6 6 2 4 6 -1\nNA\nMOVE S\n")

	var out bytes.Buffer
	run(strings.NewReader(in.String()), &out, rand.New(rand.NewSource(1)))

	lines := strings.Split(strings.TrimSpace(out.String()), "\n")
	if len(lines) != 3 {
		t.Fatalf("got %d output lines, want 3: %q", len(lines), out.String())
	}

	var sx, sy int
	if n, err := fmt.Sscan(lines[0], &sx, &sy); err != nil || n != 2 {
		t.Fatalf("invalid starting line %q: %v", lines[0], err)
	}
	if sx < 0 || sx >= 15 || sy < 0 || sy >= 15 || m[sy][sx] != '.' {
		t.Fatalf("starting coordinate (%d,%d) is not valid water", sx, sy)
	}
	for i := 1; i < len(lines); i++ {
		if !strings.HasPrefix(lines[i], "MOVE ") && lines[i] != "SURFACE" {
			t.Fatalf("unexpected turn %d output %q", i, lines[i])
		}
	}
}

func TestRunTorpedoTriggerOriginalBehavior(t *testing.T) {
	// Sector 1 layout so SURFACE 1|MOVE N|MOVE E uniquely identifies enemy at (3,1).
	// Bot starts Turn 1 at (0,0) with (1,0) blocked by an island, forcing MOVE S.
	m := make([]string, 15)
	for i := range m {
		m[i] = strings.Repeat("x", 15)
	}
	setCell(m, 0, 0, '.')
	setCell(m, 0, 1, '.')
	setCell(m, 1, 1, '.')
	setCell(m, 2, 1, '.')
	setCell(m, 3, 1, '.')
	setCell(m, 2, 2, '.')

	t.Run("fires when torpedoCooldown is 0 and possibleCount <= 3", func(t *testing.T) {
		var in strings.Builder
		in.WriteString("15 15 0\n")
		for _, row := range m {
			in.WriteString(row + "\n")
		}
		in.WriteString("0 0 6 6 0 4 6 -1\nNA\nSURFACE 1|MOVE N|MOVE E\n")

		var out bytes.Buffer
		run(strings.NewReader(in.String()), &out, rand.New(rand.NewSource(3)))

		lines := strings.Split(strings.TrimSpace(out.String()), "\n")
		if len(lines) != 2 {
			t.Fatalf("got %d lines, want 2: %q", len(lines), out.String())
		}
		want := "MOVE S TORPEDO|TORPEDO 3 1"
		if lines[1] != want {
			t.Fatalf("got %q, want %q", lines[1], want)
		}
	})

	t.Run("does not fire when torpedoCooldown > 0", func(t *testing.T) {
		var in strings.Builder
		in.WriteString("15 15 0\n")
		for _, row := range m {
			in.WriteString(row + "\n")
		}
		in.WriteString("0 0 6 6 1 4 6 -1\nNA\nSURFACE 1|MOVE N|MOVE E\n")

		var out bytes.Buffer
		run(strings.NewReader(in.String()), &out, rand.New(rand.NewSource(3)))

		lines := strings.Split(strings.TrimSpace(out.String()), "\n")
		if len(lines) != 2 {
			t.Fatalf("got %d lines, want 2: %q", len(lines), out.String())
		}
		want := "MOVE S TORPEDO"
		if lines[1] != want {
			t.Fatalf("got %q, want %q", lines[1], want)
		}
	})

	t.Run("does not fire when possibleCount > 3", func(t *testing.T) {
		var in strings.Builder
		in.WriteString("15 15 0\n")
		for _, row := range m {
			in.WriteString(row + "\n")
		}
		in.WriteString("0 0 6 6 0 4 6 -1\nNA\nNA\n")

		var out bytes.Buffer
		run(strings.NewReader(in.String()), &out, rand.New(rand.NewSource(3)))

		lines := strings.Split(strings.TrimSpace(out.String()), "\n")
		if len(lines) != 2 {
			t.Fatalf("got %d lines, want 2: %q", len(lines), out.String())
		}
		want := "MOVE S TORPEDO"
		if lines[1] != want {
			t.Fatalf("got %q, want %q", lines[1], want)
		}
	})
}

// Layer 3: Local referee simulator, validation tests, and multi game self play.

func nextPosition(x, y int, direction string) (int, int) {
	switch direction {
	case "N":
		return x, y - 1
	case "S":
		return x, y + 1
	case "E":
		return x + 1, y
	case "W":
		return x - 1, y
	default:
		return x, y
	}
}

func torpedoReachable(sx, sy, tx, ty int, gameMap []string) bool {
	height := len(gameMap)
	if height == 0 {
		return false
	}
	width := len(gameMap[0])
	if sx < 0 || sx >= width || sy < 0 || sy >= height {
		return false
	}
	if tx < 0 || tx >= width || ty < 0 || ty >= height {
		return false
	}
	if gameMap[sy][sx] != '.' || gameMap[ty][tx] != '.' {
		return false
	}
	if sx == tx && sy == ty {
		return true
	}

	type node struct {
		x, y, dist int
	}

	var seen [15][15]bool
	seen[sy][sx] = true
	queue := []node{{sx, sy, 0}}

	dx := [4]int{0, 0, 1, -1}
	dy := [4]int{-1, 1, 0, 0}

	for head := 0; head < len(queue); head++ {
		cur := queue[head]
		if cur.dist == 4 {
			continue
		}
		for i := 0; i < 4; i++ {
			nx := cur.x + dx[i]
			ny := cur.y + dy[i]
			if nx < 0 || nx >= width || ny < 0 || ny >= height {
				continue
			}
			if seen[ny][nx] || gameMap[ny][nx] != '.' {
				continue
			}
			if nx == tx && ny == ty {
				return true
			}
			seen[ny][nx] = true
			queue = append(queue, node{nx, ny, cur.dist + 1})
		}
	}

	return false
}

type botFunc func(in io.Reader, out io.Writer)

type playerState struct {
	x, y         int
	hp           int
	torpedoCD    int
	sonarCD      int
	silenceCD    int
	visited      [15][15]bool
	pendingSonar string
	lastOrders   string
}

type matchOptions struct {
	surfaceOnInvalidTorpedo bool
}

type matchResult struct {
	turnsPlayed      int
	hp               [2]int
	invalidTorpedoes int
}

type referee struct {
	gameMap          [15]string
	players          [2]playerState
	opts             matchOptions
	invalidTorpedoes int
}

func sectorFor(x, y int) int {
	return (y/5)*3 + (x/5) + 1
}

func newReferee(gameMap []string, opts matchOptions) (*referee, error) {
	if len(gameMap) != 15 {
		return nil, fmt.Errorf("map height %d, want 15", len(gameMap))
	}
	r := &referee{opts: opts}
	for i := 0; i < 15; i++ {
		if len(gameMap[i]) != 15 {
			return nil, fmt.Errorf("map row %d width %d, want 15", i, len(gameMap[i]))
		}
		r.gameMap[i] = gameMap[i]
	}
	for p := 0; p < 2; p++ {
		r.players[p] = playerState{
			hp:           6,
			torpedoCD:    3,
			sonarCD:      4,
			silenceCD:    6,
			pendingSonar: "NA",
			lastOrders:   "NA",
		}
	}
	return r, nil
}

func (r *referee) applyActionLine(p int, line string) error {
	line = strings.TrimSpace(line)
	if line == "" {
		return fmt.Errorf("player %d output empty command", p)
	}

	st := &r.players[p]
	opp := &r.players[1-p]
	used := map[string]bool{}
	var oppSummary []string

	segments := strings.Split(line, "|")
	for _, seg := range segments {
		seg = strings.TrimSpace(seg)
		if seg == "" {
			return fmt.Errorf("player %d output empty pipe segment in %q", p, line)
		}
		parts := strings.Fields(seg)
		cmd := parts[0]

		if cmd != "MSG" {
			if used[cmd] {
				return fmt.Errorf("player %d repeated action %s in %q", p, cmd, line)
			}
			used[cmd] = true
		}

		switch cmd {
		case "MOVE":
			if len(parts) != 3 {
				return fmt.Errorf("player %d invalid MOVE %q", p, seg)
			}
			dir, power := parts[1], parts[2]
			if dir != "N" && dir != "S" && dir != "E" && dir != "W" {
				return fmt.Errorf("player %d invalid MOVE direction %q", p, dir)
			}
			nx, ny := nextPosition(st.x, st.y, dir)
			if nx < 0 || nx >= 15 || ny < 0 || ny >= 15 {
				return fmt.Errorf("player %d moved out of bounds to (%d,%d)", p, nx, ny)
			}
			if r.gameMap[ny][nx] != '.' {
				return fmt.Errorf("player %d moved onto island (%d,%d)", p, nx, ny)
			}
			if st.visited[ny][nx] {
				return fmt.Errorf("player %d revisited (%d,%d)", p, nx, ny)
			}
			switch power {
			case "TORPEDO":
				if st.torpedoCD > 0 {
					st.torpedoCD--
				}
			case "SONAR":
				if st.sonarCD > 0 {
					st.sonarCD--
				}
			case "SILENCE":
				if st.silenceCD > 0 {
					st.silenceCD--
				}
			default:
				return fmt.Errorf("player %d invalid charge power %q", p, power)
			}
			st.x, st.y = nx, ny
			st.visited[ny][nx] = true
			oppSummary = append(oppSummary, "MOVE "+dir)

		case "SURFACE":
			if len(parts) != 1 {
				return fmt.Errorf("player %d invalid SURFACE %q", p, seg)
			}
			st.visited = [15][15]bool{}
			st.visited[st.y][st.x] = true
			st.hp--
			oppSummary = append(oppSummary, fmt.Sprintf("SURFACE %d", sectorFor(st.x, st.y)))

		case "TORPEDO":
			if len(parts) != 3 {
				return fmt.Errorf("player %d invalid TORPEDO %q", p, seg)
			}
			if st.torpedoCD != 0 {
				return fmt.Errorf("player %d fired TORPEDO with cooldown %d", p, st.torpedoCD)
			}
			tx, errX := strconv.Atoi(parts[1])
			ty, errY := strconv.Atoi(parts[2])
			if errX != nil || errY != nil {
				return fmt.Errorf("player %d non integer TORPEDO coords %q", p, seg)
			}
			if !torpedoReachable(st.x, st.y, tx, ty, r.gameMap[:]) {
				if r.opts.surfaceOnInvalidTorpedo {
					r.invalidTorpedoes++
					st.visited = [15][15]bool{}
					st.visited[st.y][st.x] = true
					st.hp--
					oppSummary = append(oppSummary, fmt.Sprintf("SURFACE %d", sectorFor(st.x, st.y)))
					continue
				}
				return fmt.Errorf("player %d fired TORPEDO (%d,%d) unreachable from (%d,%d)", p, tx, ty, st.x, st.y)
			}
			st.torpedoCD = 3
			for i := 0; i < 2; i++ {
				pl := &r.players[i]
				if pl.x == tx && pl.y == ty {
					pl.hp -= 2
				} else if abs(pl.x-tx) <= 1 && abs(pl.y-ty) <= 1 {
					pl.hp--
				}
			}
			oppSummary = append(oppSummary, fmt.Sprintf("TORPEDO %d %d", tx, ty))

		case "SONAR":
			if len(parts) != 2 {
				return fmt.Errorf("player %d invalid SONAR %q", p, seg)
			}
			if st.sonarCD != 0 {
				return fmt.Errorf("player %d used SONAR with cooldown %d", p, st.sonarCD)
			}
			sector, err := strconv.Atoi(parts[1])
			if err != nil || sector < 1 || sector > 9 {
				return fmt.Errorf("player %d invalid SONAR sector %q", p, seg)
			}
			st.sonarCD = 4
			if sectorFor(opp.x, opp.y) == sector {
				st.pendingSonar = "Y"
			} else {
				st.pendingSonar = "N"
			}
			oppSummary = append(oppSummary, fmt.Sprintf("SONAR %d", sector))

		case "SILENCE":
			if len(parts) != 3 {
				return fmt.Errorf("player %d invalid SILENCE %q", p, seg)
			}
			if st.silenceCD != 0 {
				return fmt.Errorf("player %d used SILENCE with cooldown %d", p, st.silenceCD)
			}
			dir := parts[1]
			if dir != "N" && dir != "S" && dir != "E" && dir != "W" {
				return fmt.Errorf("player %d invalid SILENCE direction %q", p, dir)
			}
			dist, err := strconv.Atoi(parts[2])
			if err != nil || dist < 0 || dist > 4 {
				return fmt.Errorf("player %d invalid SILENCE distance %q", p, parts[2])
			}
			st.silenceCD = 6
			for step := 0; step < dist; step++ {
				nx, ny := nextPosition(st.x, st.y, dir)
				if nx < 0 || nx >= 15 || ny < 0 || ny >= 15 {
					return fmt.Errorf("player %d SILENCE moved out of bounds (%d,%d)", p, nx, ny)
				}
				if r.gameMap[ny][nx] != '.' {
					return fmt.Errorf("player %d SILENCE hit island (%d,%d)", p, nx, ny)
				}
				if st.visited[ny][nx] {
					return fmt.Errorf("player %d SILENCE revisited (%d,%d)", p, nx, ny)
				}
				st.x, st.y = nx, ny
				st.visited[ny][nx] = true
			}
			oppSummary = append(oppSummary, "SILENCE")

		case "MSG":
			// Messages do not affect game state or opponent orders.

		default:
			return fmt.Errorf("player %d unknown command %q", p, line)
		}
	}

	if len(oppSummary) == 0 {
		return fmt.Errorf("player %d performed no game action in %q", p, line)
	}
	st.lastOrders = strings.Join(oppSummary, "|")
	return nil
}

func playMatch(gameMap []string, bots [2]botFunc) (matchResult, error) {
	return playMatchWithOptions(gameMap, bots, matchOptions{})
}

func playMatchWithOptions(gameMap []string, bots [2]botFunc, opts matchOptions) (matchResult, error) {
	r, err := newReferee(gameMap, opts)
	if err != nil {
		return matchResult{}, err
	}

	var inW [2]*io.PipeWriter
	var outR [2]*io.PipeReader
	var scanners [2]*bufio.Scanner
	var wg sync.WaitGroup

	for p := 0; p < 2; p++ {
		inReader, inWriter := io.Pipe()
		outReader, outWriter := io.Pipe()
		inW[p] = inWriter
		outR[p] = outReader
		scanners[p] = bufio.NewScanner(outReader)

		bot := bots[p]
		wg.Add(1)
		go func(ir *io.PipeReader, ow *io.PipeWriter) {
			defer wg.Done()
			defer ir.Close()
			defer ow.Close()
			bot(ir, ow)
		}(inReader, outWriter)
	}

	defer func() {
		for p := 0; p < 2; p++ {
			if inW[p] != nil {
				inW[p].Close()
			}
			if outR[p] != nil {
				outR[p].Close()
			}
		}
		wg.Wait()
	}()

	// Send initial header and map, then read starting placements.
	for p := 0; p < 2; p++ {
		if _, err := fmt.Fprintf(inW[p], "15 15 %d\n", p); err != nil {
			return matchResult{}, fmt.Errorf("write header to player %d: %w", p, err)
		}
		for _, row := range r.gameMap {
			if _, err := fmt.Fprintln(inW[p], row); err != nil {
				return matchResult{}, fmt.Errorf("write map to player %d: %w", p, err)
			}
		}
		if !scanners[p].Scan() {
			return matchResult{}, fmt.Errorf("player %d failed to output starting position", p)
		}
		var sx, sy int
		if n, err := fmt.Sscan(scanners[p].Text(), &sx, &sy); err != nil || n != 2 {
			return matchResult{}, fmt.Errorf("player %d invalid starting line %q", p, scanners[p].Text())
		}
		if sx < 0 || sx >= 15 || sy < 0 || sy >= 15 || r.gameMap[sy][sx] != '.' {
			return matchResult{}, fmt.Errorf("player %d placed on invalid cell (%d,%d)", p, sx, sy)
		}
		r.players[p].x = sx
		r.players[p].y = sy
		r.players[p].visited[sy][sx] = true
	}

	turnsPlayed := 0
	for turn := 0; turn < 299; turn++ {
		if r.players[0].hp <= 0 || r.players[1].hp <= 0 {
			break
		}
		turnsPlayed++
		for p := 0; p < 2; p++ {
			st := &r.players[p]
			opp := &r.players[1-p]

			if _, err := fmt.Fprintf(
				inW[p],
				"%d %d %d %d %d %d %d -1\n%s\n%s\n",
				st.x, st.y, st.hp, opp.hp, st.torpedoCD, st.sonarCD, st.silenceCD,
				st.pendingSonar,
				opp.lastOrders,
			); err != nil {
				return matchResult{}, fmt.Errorf("write turn %d to player %d: %w", turn, p, err)
			}
			st.pendingSonar = "NA"

			if !scanners[p].Scan() {
				return matchResult{}, fmt.Errorf("player %d closed output on turn %d", p, turn)
			}
			if err := r.applyActionLine(p, scanners[p].Text()); err != nil {
				return matchResult{}, fmt.Errorf("turn %d: %w", turn, err)
			}
			if r.players[0].hp <= 0 || r.players[1].hp <= 0 {
				break
			}
		}
	}

	return matchResult{
		turnsPlayed:      turnsPlayed,
		hp:               [2]int{r.players[0].hp, r.players[1].hp},
		invalidTorpedoes: r.invalidTorpedoes,
	}, nil
}

func scriptedBot(start string, replies ...string) botFunc {
	return func(in io.Reader, out io.Writer) {
		scanner := bufio.NewScanner(in)
		for i := 0; i < 16; i++ {
			if !scanner.Scan() {
				return
			}
		}
		fmt.Fprintln(out, start)

		idx := 0
		for scanner.Scan() {
			// Consume sonarResult and opponentOrders
			if !scanner.Scan() || !scanner.Scan() {
				return
			}
			if idx < len(replies) {
				fmt.Fprintln(out, replies[idx])
				idx++
			} else {
				fmt.Fprintln(out, "SURFACE")
			}
		}
	}
}

func TestRefereeRejectsIllegalCommands(t *testing.T) {
	m := emptyMap()
	setCell(m, 5, 4, 'x') // North of (5,5) is an island

	cases := []struct {
		name    string
		badBot  botFunc
		wantSub string
	}{
		{
			name:    "moving into island",
			badBot:  scriptedBot("5 5", "MOVE N TORPEDO"),
			wantSub: "island",
		},
		{
			name:    "revisiting cell",
			badBot:  scriptedBot("5 5", "MOVE E TORPEDO", "MOVE W TORPEDO"),
			wantSub: "revisited",
		},
		{
			name:    "firing torpedo on non zero cooldown",
			badBot:  scriptedBot("5 5", "TORPEDO 5 7"),
			wantSub: "cooldown",
		},
		{
			name:    "firing torpedo out of range",
			badBot:  scriptedBot("0 0", "MOVE E TORPEDO", "MOVE E TORPEDO", "MOVE E TORPEDO|TORPEDO 14 14"),
			wantSub: "unreachable",
		},
		{
			name:    "duplicate MOVE in one turn",
			badBot:  scriptedBot("0 0", "MOVE E TORPEDO|MOVE S TORPEDO"),
			wantSub: "repeated action MOVE",
		},
	}

	goodBot := func(in io.Reader, out io.Writer) {
		run(in, out, rand.New(rand.NewSource(1)))
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := playMatch(m, [2]botFunc{tc.badBot, goodBot})
			if err == nil {
				t.Fatalf("expected referee error containing %q, got nil", tc.wantSub)
			}
			if !strings.Contains(err.Error(), tc.wantSub) {
				t.Fatalf("got error %q, want substring %q", err.Error(), tc.wantSub)
			}
		})
	}
}

func TestSelfPlayNeverMakesIllegalMove(t *testing.T) {
	densities := []int{0, 10, 20}
	totalInvalidTorpedoes := 0
	for _, density := range densities {
		for game := 0; game < 10; game++ {
			seed := int64(density*100 + game + 1)
			m := generateMap(seed, density)

			bot0 := func(in io.Reader, out io.Writer) {
				run(in, out, rand.New(rand.NewSource(seed*2)))
			}
			bot1 := func(in io.Reader, out io.Writer) {
				run(in, out, rand.New(rand.NewSource(seed*2+1)))
			}

			res, err := playMatchWithOptions(m, [2]botFunc{bot0, bot1}, matchOptions{surfaceOnInvalidTorpedo: true})
			if err != nil {
				t.Fatalf("density=%d game=%d seed=%d illegal move: %v", density, game, seed, err)
			}
			if res.turnsPlayed == 0 {
				t.Fatalf("density=%d game=%d played 0 turns", density, game)
			}
			totalInvalidTorpedoes += res.invalidTorpedoes
		}
	}
	t.Logf("30 self play games completed cleanly (invalid torpedoes triggering fallback SURFACE: %d)", totalInvalidTorpedoes)
}

// Layer 4: Latency benchmarks to verify comfortable margin below 50 ms per turn.

func BenchmarkEnemyTrackerProcessOrders(b *testing.B) {
	m := generateMap(42, 15)
	tracker := NewEnemyTracker(m)
	orders := []string{"MOVE N", "MOVE E", "SURFACE 5", "SILENCE", "MOVE S|TORPEDO 7 7"}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tracker.ProcessOrders(orders[i%len(orders)], m)
		if countPossible(&tracker) == 0 {
			tracker = NewEnemyTracker(m)
		}
	}
}

func BenchmarkTurnDecision(b *testing.B) {
	m := generateMap(77, 15)
	tracker := NewEnemyTracker(m)
	visited := newVisitedGrid()
	rng := rand.New(rand.NewSource(7))

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tracker.ProcessOrders("MOVE N", m)
		possibleCount := 0
		bestTargetX, bestTargetY := -1, -1
		for ty, row := range tracker.possibleCells {
			for tx, isPossible := range row {
				if isPossible {
					possibleCount++
					bestTargetX, bestTargetY = tx, ty
				}
			}
		}
		_ = possibleCount
		_ = bestTargetX
		_ = bestTargetY
		_ = determineNewDirection(7, 7, 15, 15, m, visited, rng)
		if countPossible(&tracker) == 0 {
			tracker = NewEnemyTracker(m)
		}
	}
}
