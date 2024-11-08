import Screen from "./assets/Screen.ts"

export default function SnakeGame() {
  const s = new Screen("snakeGame", 12, 12);

  s.drawScreen();

  return (
    <div id="something-rendered">
    </div>
  )
}
