
const APIURL = "http://192.168.0.9:8080"

export const getAllInvestors = async() => {
    try {
        const resp = await fetch(`${APIURL}/investors`)
        return resp.json()
    } catch (error) {
        console.error(error)
    }
    
}

export const getAllStocks = async() => {
    try {
        const resp = await fetch(`${APIURL}/stocks`)
        return resp.json()
    } catch (error) {
        console.error(error)
    }
}

export const getInvestorStocks = async(investorId) => {
    try {
        const resp = await fetch(`${APIURL}/stocks/${investorId}`)
        return resp.json()
    } catch (error) {
        console.error(error)
    }
}

export const createInvestor = async(name,stockList) => {
    try {
        const resp = await fetch(`${APIURL}/investors`,{
            method:"POST",
            body:JSON.stringify({
                name,
                stocks:stockList
            })
        })
        if (resp.status !== 200) {
            throw new Error("Can't Create Investor")
        }
        const response = await getAllStocks()
        return  response.json()
    } catch (error) {
        console.error(error)
    }
}
