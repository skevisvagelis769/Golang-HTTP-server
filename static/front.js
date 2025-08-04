let btn = document.getElementById("submit")
let srch = document.getElementById("search")
let input = document.getElementById("input")

btn.addEventListener('click',()=>{
    let userData ={
    name: input.value
    };
    console.log(userData.name)

    fetch("http://localhost:8080/users",{
        method: "POST",
        headers:{
            "Content-Type":"application/json"
        },
        body: JSON.stringify(userData)
        
    })
    .then(response=>response.text())
    .then(data=>{
        console.log("Serber response:",data)
        
    })
    .catch(error=>{
        console.log("err",error)
    })
})
srch.addEventListener('click',()=>{
    fetch(`http://localhost:8080/users/search/${input.value}`,{
        method: "GET",
        headers:{
            "Content-Type":"application/json"
        }
    }).then(res => res.json())
    .then(users=>{
        console.log(users.name)
        const list = document.getElementById("output")
        list.textContent= `${users.name}`
        
    }).catch(err=> console.log("Error: ",err))
})