import { useState } from "react";
import AppContext from "./context";

const AppContextProvider = ({children}) => {
    const [stocks,setStocks] = useState([])
    const [isLoading,setIsLoading] = useState(false)

    return (
        <AppContext.Provider value={{stocks,setStocks,isLoading,setIsLoading}}>
            {children}
        </AppContext.Provider>
    )
}

export default AppContextProvider