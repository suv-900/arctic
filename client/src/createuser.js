import {useState,useEffect} from 'react'


function CreateNewUser(){
    const[nameinput,setnameinput]=useState('')
    const[emailInput,setEmailInput]=useState('')
    const[passInput,setPassInput]=useState('')
    const[userExists,setUserExists]=useState(false) 
    
    async function searchUsername(){
        try{
            const res=await fetch("http://localhost:8000/checkusername",{
                method:"POST",
                body:JSON.stringify({username:nameinput})
            })
            if(res.status===409){
                console.log(res.status)
                setUserExists(true)       
            }
        }catch(err){
            console.log(err)
        }
    }
    useEffect(()=>{
        const delay=setTimeout(()=>{
            searchUsername()
        },1000)
        return ()=>clearTimeout(delay)
    },[nameinput]) 

    return (
        <div>
            <form>
                <label>Username</label>
                <input type="text" onChange={(e)=>{
                        setnameinput(e.target.value)
                        
                        }} />
                {userExists?(
                    <div>User already exists.</div>
                ):
                (<></>) }

                <label>Email</label>
                <input type="text" onChange={(e)=>setEmailInput(e.target.value)} />
                
                <label>Password</label>
                <input type="text" onChange={(e)=>setPassInput(e.target.value)} />
                
                <button type="submit">Submit</button>
            </form>
            <a href="/">Login?</a>
        </div>
    )
}

export default CreateNewUser