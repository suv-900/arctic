const express=require('express')
const homeController=require('./controllers/homeController')
const checkServerHealth=require('./controllers/checkServer')
const cors=require('cors')

const app=express()

app.use(cors(
    {
        origin:'http://localhost:3000'
    }))

app.use('/home',homeController)
app.use('/checkserver',checkServerHealth)

app.listen(4000,()=>console.log("Server started on port 4000\n"))
 