import React,{useContext} from 'react'
import AppContext from '../context/context'
import Loader from '../loader/loader'
import "./stock.css"

const Stock = ({stock,id}) => {
  const {loading,refreshedInvestor} = useContext(AppContext)
  const refreshedStock = () => {
    if (loading && refreshedInvestor === 'all') {
      return true
    }
    return loading && refreshedInvestor === id
  }
 

  return (
    
      stock.map((info,idx )=> (
        <div key={`Investor ${id} stock ${idx} `} className="stock-tile">
            <h4>{info.symbol}</h4>
            {refreshedStock()  ? <Loader />:
            (<span>
                {info.price ? info.price : "-"} 
                  &nbsp;
                {info.change ? info.change: "-"}
            </span>
            )}
        </div>
      ))
      
  
  )
}

export default Stock