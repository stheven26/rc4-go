const localAuth = () => {
    const auth = localStorage.getItem("login") || JSON.stringify({status: false, name: ""})

    if (localStorage.getItem("login") == undefined) localStorage.setItem("login", auth)

    const data = JSON.parse(auth)
    
    return data
}

const setLocalAuth = (props: {status: boolean}) => localStorage.setItem("login", JSON.stringify(props))

export {localAuth, setLocalAuth};