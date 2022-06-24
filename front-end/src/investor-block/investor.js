import {useContext} from 'react'
import { getInvestorStocks } from '../api/api-requests'
import AppContext from '../context/context'
import Stock from '../stock-info/stock'
import "./investor.css"

const Investor = ({investor}) => {

  const {stocks,isLoading,setIsLoading,setStocks} = useContext(AppContext)


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

  const refreshPrices = () =>{
    setIsLoading(true)
    getInvestorStocks(investor.id).then(resp => {
     
      if (!resp) return;
      const updatedStocks = updateStock(resp.data)
  
      if (!updatedStocks) {
        setStocks([...stocks,resp.data])
        return
      }
      setStocks(updatedStocks)
    })
    setIsLoading(false)
  }



  const generateStocks = () => {
    const investorStocks = stocks.find(stock => stock.investorId === investor.id)
    if (stocks.length === 0 || !investorStocks){
      return <Stock stock={investor.stocks} id={investor.id} isLoading={false} />
    }

    return  <Stock stock={investorStocks.stockInfo} id={investorStocks.investorId} isLoading={isLoading} />
  }

  return (
    <section className='investor'>
      <div className='investor__head' >
        <h1>{investor.name}</h1>
        <button onClick={refreshPrices}>Refresh</button>
      </div>
      <div className='investor__stocks__container'>
          {generateStocks()}
      </div>
    </section>
  )
}

export default Investor