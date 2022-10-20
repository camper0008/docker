const updateCounter = async () => {
    const element = document.getElementById("amount")
    const amount = await (await fetch("/api/amount")).text()
    element.textContent = amount
}

const buttonClicked = async () => {
    await fetch("/api/increment", {
        method: "POST",
    })
    await updateCounter()
}

const main = async () => {
    await updateCounter()
    const button = document.getElementById("increment-button")
    button.addEventListener("click", () => buttonClicked())
}

main()
