const { useState } = React
function App() {
 const [value, setValue] = useState('')
 const [response, setResponse] = useState('')
function Send(e) {
	e.preventDefault();
 const config = {
    headers: {
      'Content-Type':'application/x-www-form-urlencoded'
    }
}
axios.post('http://192.168.1.1:3000',{send:value}, config)
      .then(res => {
       setResponse(res.data)
       })
}
 return (
  <div>
   <form>
    <label> Name:
    <input type = "text" value={value} onChange={e => setValue(e.target.value)}/>
     </label>
     <button onClick={Send}>Add</button><br/>
     <span>{response}</span>
      </form>
   </div>
   )
}
ReactDOM.render(<App />, document.getElementById('mydiv'))  
