import { useState } from "react";
import axios from "axios";
import "./App.css";

const ENDPOINT = "http://127.0.0.1:8080";

const langTemplate = [
  '#include<stdio.h>\nint main(){\n\tprintf("Hello, World!!!");\n\treturn 0;\n}\n',
  'print("Hello, World!!!")',
];
function App() {
  const [nowLang, SetNowLang] = useState("");
  const [textCode, setTextCode] = useState("");
  const [output, setOutput] = useState("");
  const [isCompiling, setIsCompiling] = useState(false);

  async function complie() {
    setIsCompiling(true);
    const clientCode = JSON.stringify({ lang: nowLang, code: textCode })
    const input = await axios.post(`${ENDPOINT}/run`, clientCode, {
      headers: {
        'Content-Type': 'application/json',
      },
      'Access-Control-Allow-Origin': '*'
    });

    const output = await axios.get(`${ENDPOINT}/ouput`)
    console.log(output.data)

    // console.log(textCode)
    // console.log(`${ENDPOINT}/output/${nowLang}`)

    // const output = await fetch(`${ENDPOINT}/output/${nowLang}`, { mode: 'no-cors' });

    // console.log("Output success");
    // console.log(output)

    setIsCompiling(false);
  }

  const changeLang = (e) => {
    const lang = e.target.value;

    SetNowLang(lang);
    if (lang === "cpp") setTextCode(langTemplate[0]);
    else if (lang === "py") setTextCode(langTemplate[1]);
    else setTextCode("");
  };

  return (
    <div className="App">
      <header className="App-header">
        <a>Simple Compiler</a>
      </header>

      <div className="body">
        <div>
          <a>Language: </a>
          <select value={nowLang} onChange={changeLang}>
            <option value="">Choose the Language</option>
            <option value="cpp">C++</option>
            <option value="py">Python</option>
          </select>
        </div>

        <div className="textArea">
          <textarea value={textCode} onChange={(e) => {
            setTextCode(e.target.value)
          }}></textarea>
        </div>

        <button onClick={complie} disabled={isCompiling}>
          Compile
        </button>

        <div className="output">
          <a className="maintext">Output: </a> <br />
          <p>{output}</p>
        </div>
      </div>
    </div>
  );
}

export default App;
