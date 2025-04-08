# 🐍A simple and fun **Snake Game** built using HTML5 Canvas, styled with CSS, and served with a Go backend! 🎮✨

## 📦 Features

- Classic Snake gameplay: eat food, grow longer, and avoid crashing into yourself
- Canvas-based rendering
- Responsive keyboard controls (Arrow keys)
- Score tracking
- Game-over detection with auto reset
- Food spawns randomly without overlapping the snake
- Simple backend using Go to serve static files

## 🚀 Getting Started

### 🔧 Prerequisites

- [Go](https://golang.org/dl/) installed

### ▶️ Run the Game

```bash
go run main.go
```

## 📄 Code Highlights

### `main.go`

This simple Go server serves static files (HTML, CSS, JS) from the `static` directory and runs the game on [http://localhost:8080](http://localhost:8080):

```go
fs := http.FileServer(http.Dir("static"))
http.Handle("/", fs)
http.ListenAndServe(":8080", nil)
```
## 🎮 Controls

Use the arrow keys on your keyboard to control the snake:

- 🔼 **Arrow Up**: Move up
- 🔽 **Arrow Down**: Move down
- ◀️ **Arrow Left**: Move left
- ▶️ **Arrow Right**: Move right



