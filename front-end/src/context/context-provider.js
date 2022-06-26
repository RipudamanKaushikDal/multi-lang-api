import { useState } from "react";
import AppContext from "./context";

const AppContextProvider = ({children}) => {
    const [stocks,setStocks] = useState([])
    const [loading,setLoading] = useState(false)
    const [investorList,setInvestorList] = useState([])
    const [refreshedInvestor,setRefreshedInvestor] = useState()

    return (
        <AppContext.Provider value={{stocks,setStocks,loading,setLoading,investorList,setInvestorList,refreshedInvestor,setRefreshedInvestor}}>
            {children}
        </AppContext.Provider>
    )
}

export default AppContextProvider