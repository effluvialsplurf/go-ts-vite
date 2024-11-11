import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
//import App from './App.tsx'

//createRoot(document.getElementById('root')!).render(
//  <StrictMode>
//    <App />
//  </StrictMode>,
//)

import AIHome from './AIHome.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <AIHome />
  </StrictMode>,
)

// import SnakeGame from './SnakeGame.tsx'

// createRoot(document.getElementById('snakeGame')!).render(
//  <StrictMode>
//    <SnakeGame />
//  </StrictMode>
// )


