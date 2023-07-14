import {useState} from 'react'
export default function LoginPage(){
    const [serverOK,setserverOK]=useState(false)
    async function checkNodeServer(){
        try{
        const res=await fetch("http://localhost:4000/checkserver",{
            method:"GET"
        })
        if(res.ok){
            setserverOK(true)
        }
        }catch(err){
            console.log(err)
        }
    }

    return(
        <div>
            <label>Username</label>
            <input type="text"  />
            <label>Password</label>
            <input type="text"/>
            
            <a href="/register">New User?</a>
            <button onClick={checkNodeServer} >Guest mode</button>
            {serverOK?(
                <h1>Node server running</h1>
            )
            :(
                <div></div>
            )}
        </div>
    )
}