function displayDay(date) {
    const jsdate = new Date(date)

    fetch(`http://localhost:8888/tt/${jsdate.getTime()}`).then(async (res) => {
        try {
            const body = await res.json()
            renderTimeSlots(body)
        } catch (e) {
            showEmpty()
        }
    })
}

function showEmpty() {
    document.getElementById("timeslot-container").innerHTML = ''
    document.getElementById("timeslot-container").innerHTML = '<h1>Nothing found</h1>'
}

function renderTimeSlots(timeslots) {
    document.getElementById("timeslot-container").innerHTML = ''
    for (let t of timeslots) {
        document.getElementById("timeslot-container").appendChild(createTimeSlotElement(t))
    }
}

document.addEventListener("DOMContentLoaded", () => {
   document.getElementById("select-date").addEventListener("change", (e) => {
       displayDay(e.target.value)
   })
    displayDay(new Date())
})

function createTimeSlotElement(timeSlot) {
    const container = document.createElement('div');
    container.classList.add('time-slot')

    const title = document.createElement('h3');
    title.textContent = timeSlot.Text;

    const start = document.createElement('p');
    start.innerHTML = `<strong>Start:</strong> ${new Date(timeSlot.Start).toLocaleString()}`;

    const end = document.createElement('p');
    end.innerHTML = `<strong>End:</strong> ${new Date(timeSlot.End).toLocaleString()}`;

    container.appendChild(title);
    container.appendChild(start);
    container.appendChild(end);

    return container;
}

function quit() {
    fetch(`http://localhost:8888/quit`).then(() => {
        window.close()
    })
}