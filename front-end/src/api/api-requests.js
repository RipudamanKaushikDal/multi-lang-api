
const APIURL = "http://localhost:8080"

export const getAllInvestors = async() => {
    try {
        const resp = await fetch(`${APIURL}/investors`)
        return resp.json()
    } catch (error) {
        console.error(error)
    }
    
}

export const getAllStocks = async(investorList) => {
    try {
        const resp = await fetch(`${APIURL}/stocks`,{
            method:"POST",
            headers:{
                'Content-Type': 'application/json'
            },
            body:JSON.stringify(investorList)
        })
        return resp.json()
    } catch (error) {
        console.error(error)
    }
}

