import {createBrowserRouter as Router,Route, BrowserRouter,Routes} from 'react-router-dom'
import { useEffect,useState} from 'react';
import React from 'react'
import LoginPage from './loginpage'
import CreateNewUser from './createuser'
function App() {
  const [serverOK,setserverOK]=useState(false)
  //const[requestno,setRequestNo]=useState(0)  
  async function fetchServerStatus(){
    try{
      const res= await fetch("http://localhost:8000/checkserver",{
      method:"GET"
      })
      if (res.ok){
        setserverOK(true)
      }
    }
    catch(err){
      console.log(err)
    }
    
  }
  

    useEffect(()=>{
    fetchServerStatus()
    },[])
    
    let requestCount=0
    
    function sendServerRequests(){
      if(serverOK!=true){
       console.log(requestCount) 
       // setRequestNo(requestno+1)
       requestCount++; 
       fetchServerStatus()
        .finally(()=>{
          if(serverOK!==true && requestCount<5){
            setTimeout(sendServerRequests,5000)
          }
        })
      }
    }
    sendServerRequests()
    
  return (
    <div>
      {serverOK?(
        <div>
          <BrowserRouter>
          <Routes>
            <Route exact path="/" element={<LoginPage/>} />
            <Route path="/register" element={<CreateNewUser/>}/> 
          </Routes>
          </BrowserRouter>

        </div>
      ):(<div>
        Server Error.
        </div>)
    }  
    </div>
    );
}

export default App;
