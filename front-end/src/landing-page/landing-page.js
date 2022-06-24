import {useContext, useEffect,useState} from "react"
import {getAllInvestors, getAllStocks} from "../api/api-requests"
import AppContext from '../context/context';
import Investor from '../investor-block/investor';

const LandingPage = () => {

    const [investorList,setInvestorList] = useState([])
    const {setStocks,isLoading,setIsLoading} = useContext(AppContext)

    useEffect (() => {
        setIsLoading(true)
        getAllInvestors().then(resp => resp && setInvestorList(resp.data))
        setIsLoading(false)
      },[setIsLoading])
    
      const refreshPrices = () =>{
        setIsLoading(true)
        getAllStocks().then(resp => resp && setStocks(resp.data))
        setIsLoading(false)
      }

  
  return (
    <div>
       <div>
        <button onClick={refreshPrices}>Refresh Prices</button>
      </div>
      {investorList.map(investor => (
        <Investor investor={investor} />
      ))} 
    </div>
  )
}

export default LandingPage