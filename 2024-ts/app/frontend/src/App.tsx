import './App.css'
import { useState } from 'react'
import TicTacToe from './TickTacToe'

export default function App() {
  return (
    <>
      <section id="center">
        <h1>hello ziom</h1>
        <Button text="Click me" />
        <TicTacToe />
      </section>
    </>
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