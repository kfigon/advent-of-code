import { useState } from "react";

export default function TicTacToe() {
   return(
      <section id="center">
        <Board />
      </section>
   ) 
}

function Board() {
    const css = `
      .button-grid {
        display: grid;
        grid-template-columns: repeat(3, 1fr);
        grid-template-rows: repeat(3, 1fr);
        gap: 10px;
        width: 300px;
      }

      .button-grid button {
        aspect-ratio: 1;
        border: none;
        border-radius: 8px;
        cursor: pointer;
      }
    `;
    
    let [states, setStates] = useState(Array<ButtonState>(9).fill('empty'));
    let [currentPlayer, setCurrentPlayer] = useState<CurrentPlayer>('cross');

   let whoWon = '';
   switch (gameResult(states)) {
    case "pending": 
        whoWon = '';
        break;
    case "draw":
        whoWon = "it's a draw!";
        break;
    case "circleWon":
        whoWon = 'circle won!';
        break;
    case "crossWon":
        whoWon = 'cross won!';
        break;
    }
    let buttons = [];
    for(let i = 0; i < states.length; i++) {
        buttons.push(Button({ 
            text: gameStateToText(states[i]),
            handler: () => {
               if (whoWon){ return }
               if (states[i] !== 'empty'){ return }

               if (currentPlayer === "circle")  {
                    setCurrentPlayer('cross')
                   states[i] = 'circle'
               } else {
                    setCurrentPlayer('circle')
                   states[i] = 'cross'
               }
               setStates(states);
            }
        }));
    }

    return (
    <>
     <style>{css}</style>
      { gameStateToText(currentPlayer)}
      <div className="button-grid">
      { buttons }
      {whoWon && <p>{whoWon}</p>}
      </div>
    </>
    )
}

const gameResult = (states: ButtonState[]): GameResult => {
    const g =(b: ButtonState, allVals: [number,number,number][]) => {
        for (const vals of allVals) {
            if(vals.map(i => states[i]).every(v => v === b)) return true;
        }
        return false
    }

    let combinations: [number,number,number][] = [[0,1,2], [3,4,5],[6,7,8],[0,4,8],[2,4,6]]
    if (g('circle', combinations)) return 'circleWon';
    if (g('cross', combinations)) return 'crossWon';
    else if (states.some(v => v === 'empty')) return 'pending';
    return 'draw';
}

const gameStateToText = (b: ButtonState) => {
    switch (b) {
        case 'empty': return '';
        case 'circle': return 'O';
        case 'cross': return 'X';
        default:
            const _u: never = b;
            return _u;
    }
}

type GameResult = 'pending' | 'draw' | 'circleWon' | 'crossWon'
type CurrentPlayer = 'cross' | 'circle'
type ButtonState = 'empty' | 'circle' | 'cross'

type ButtonProps = {
    text: string
    handler: () => void
}
function Button({text, handler}: ButtonProps) {
    return (
        <button onClick={handler}>{text}</button>
    )
}
