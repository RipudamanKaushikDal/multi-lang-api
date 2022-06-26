import {useContext} from 'react'
import { getInvestorStocks } from '../api/api-requests'
import AppContext from '../context/context'
import Stock from '../stock-info/stock'
import "./investor.css"

const Investor = ({investor,setFetchTime}) => {

  const {stocks,setLoading,setStocks,setRefreshedInvestor} = useContext(AppContext)
  

  const updateStock = (newStock) => {
    const prevStocks = stocks
    for(let i=0;i<prevStocks.length;i++){
      if (prevStocks[i].investorId === newStock.investorId){
        prevStocks[i].stockInfo = newStock.stockInfo
        return prevStocks     
      }
    }
    return false
  }

  const refreshStock = async(investorId) =>{
    setLoading(true)
    setRefreshedInvestor(investorId)
    const startTime = performance.now()
    const resp = await getInvestorStocks(investor.ID)
    const endTime = performance.now()
    setFetchTime(endTime-startTime)
    if (!resp.data) return;
    const updatedStocks = updateStock(resp.data)
    setLoading(false)
  
    if (!updatedStocks) {
      setStocks([...stocks,resp.data])
      return
    }
    setStocks(updatedStocks)
  }


  

  const generateStocks = () => {
    let investorStocks = stocks.find(stock => stock.investorId === investor.ID)

    if (!investorStocks){
      return <Stock stock={investor.stocks} id={investor.ID}/>
    }

    return  <Stock stock={investorStocks.stockInfo} id={investorStocks.investorId} />
  }

  return (
    <section className='investor' key={`investor ${investor.ID}`}>
      <div className='investor__head' >
        <h1>{investor.name}</h1>
        <button onClick={() => refreshStock(investor.ID)}>Refresh</button>
      </div>
      <div className='investor__stocks__container'>
          {generateStocks()}
      </div>
    </section>
  )
}

export default Investor