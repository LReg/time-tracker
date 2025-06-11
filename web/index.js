function displayDay(date) {
    const jsdate = new Date(date)
    console.log(jsdate.getTime())

    fetch(`http://localhost:8888/tt/${jsdate.getTime()}`).then((res) => {
        console.log(res)
    })
}

document.addEventListener("DOMContentLoaded", () => {
   document.getElementById("select-date").addEventListener("change", (e) => {
       displayDay(e.target.value)
   })
    displayDay(new Date())
})

