import { useState } from 'react';
import './App.css';

const ENDPOINT = "https://127.0.0.1:8080"

const langTemplate = [
  "#include<stdio.h>\nint main(){\n\tprintf(\"Hello, World!!!\");\n\treturn 0;\n}\n",
  "print(\"Hello, World!!!\")"
]

function App() {
  const [nowLang, SetNowLang] = useState("");
  const [textCode, setTextCode] = useState("");
  const [output, setOutput] = useState("");
  const [isCompiling, setIsCompiling] = useState(false)

  async function complie() {
    setIsCompiling(true)
    const input = await fetch(`${ENDPOINT}/run/${nowLang}`, {
      method: "POST",
      headers: {
        "Content-Type": "text/plain; charset=utf-8",
      },
      body: textCode
    })

    console.log("Input Success")

    const output = await fetch(`${ENDPOINT}/output/${nowLang}`, {
      method: "GET",
      headers: {
        "Content-Type": "text/plain; charset=utf-8",
      },
    });

    console.log("Output success");

    if (output.ok) {
      let s = await output.text();
      setOutput(s);
      console.log(s);
    } else {
      let err = output.status;
      console.log(err);
    }
    setIsCompiling(false);
  }

  const changeLang = e => {
    const lang = e.target.value

    SetNowLang(lang)
    if (lang === "cpp") setTextCode(langTemplate[0])
    else if (lang === "py") setTextCode(langTemplate[1])
    else setTextCode("")
  }

  return (
    <div className="App">
      <header className="App-header">
        <a>Simple Compiler</a>
      </header>

      <div className='body'>
        <div>
          <a>Language: </a>
          <select value={nowLang} onChange={changeLang}>
            <option value="">Choose the Language</option>
            <option value="cpp">C++</option>
            <option value="py">Python</option>
          </select>
        </div>

        <div className="textArea">
          <textarea defaultValue={textCode}></textarea>
        </div>

        <button onClick={complie} disabled={isCompiling}>Compile</button>

        <div className="output">
          <a className="maintext">Output: </a> <br />
          <p>{output}</p>
        </div>
      </div>
    </div >
  );
}

export default App;
