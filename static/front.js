let submit = document.getElementById("submit")
let input = document.getElementById("input")
let search = document.getElementById("search")

submit.addEventListener('click',()=>{
	console.log("clicked")
	let userData={
		name: input.value
	}
	fetch("http://localhost:8080/users",{
		method:"POST",
		headers:{
			"Content-Type":"application:json"
		},
		body:JSON.stringify(userData)
	})
	.then(response => response.text())
	.then(data=>{
		console.log("Server Response is:",data)
	})
	.catch(error=>{
		console.log("Error occured: ",error)
	})
})

search.addEventListener('click',()=>{
    //console.log("in serach ")
    const list = document.getElementById("output")
	console.log("clicked!!")
	fetch(`http://localhost:8080/users/${input.value}`,{
		method:"GET",
		headers:{
			"Content-Type":"application/json"
		}
	}).then(res=>res.json())
	.then(users=>{
		console.log(users.name)
		list.textContent=`${users.name}`
	}).catch(err=> {
		list.textContent="Could not find!"
		console.log("Error occured: ",err)
	})
})

