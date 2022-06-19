import './App.css';
import {useEffect,useState} from "react"
import {getAllInvestors, getAllStocks} from "./api/api-requests"
import Investor from './investor-block/investor';

const App = () => {

  const [investorList,setInvestorList] = useState([])
  const [stockInfo,setStockInfo] = useState({})

  useEffect (() => {
    getAllInvestors().then(resp => resp && setInvestorList(resp))
  },[])

  const refreshPrices = () =>{
    const queryBody = investorList.map(investor => {return {id:investor.id,stocks:investor.stocks}})
    getAllStocks(queryBody).then(resp => resp && setStockInfo(resp))
  }

  return (
    <div className="App">
      <div>
        <button onClick={refreshPrices}>Refresh Prices</button>
      </div>
      {investorList.map(investor => (
        <Investor name={investor.name} stocks={investor.stocks} id={investor.id} key={investor.id.toString()} stockInfo={stockInfo} />
      ))}
    </div>
  );
}

export default App;
