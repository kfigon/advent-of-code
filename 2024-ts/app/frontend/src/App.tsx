import './App.css'

export default function App() {
  const click = () =>  alert('hi there')

  return (
    <>
      <section id="center">
        <h1>hello ziom</h1>
        <Button text="Click me" event={click}/>
      </section>
    </>
  )
}

type ButtonProps = {
  text: string
  event: () => void
}
function Button({text, event}: ButtonProps) {
  return (
    <button onClick={event}>{text}</button>
  )
}