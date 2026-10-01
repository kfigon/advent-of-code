import './App.css'
import { useState } from 'react'
import TicTacToe from './TickTacToe'
import D1 from './advent_of_code/d1'
import D2 from './advent_of_code/d2'
import { BrowserRouter, Routes, Route, Link } from "react-router-dom";

export default function App() {
  return (
    <>
    <BrowserRouter>
      <Navbar/>

      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/tictactoe" element={<TicTacToe />} />
        <Route path="/d1" element={<D1 />} />
        <Route path="/d2" element={<D2 />} />
      </Routes>
    </BrowserRouter>
    </>
  )
}

function Navbar() {
  return (
    <nav>
      <div><Link to="/">Home</Link></div>
      <div><Link to="/tictactoe">TicTacToe</Link></div>
      <div><Link to="/d1">D1</Link></div>
      <div><Link to="/d2">D2</Link></div>
    </nav>
  );
}

function Home() {
  return (
      <section id="center">
        <h1>hello ziom</h1>
        <Button text="Click me" />
      </section>
  )
}

type ButtonProps = {
  text: string
  event?: () => void
}
function Button({text, event}: ButtonProps) {
  // adding memory to our components
  const [count, setCount] = useState(0);

  const handler = () => {
    setCount(count+1);
    if (event) event();
  }

  return (
    <button onClick={handler}>{`${text}, clicked ${count} times`}</button>
  )
}