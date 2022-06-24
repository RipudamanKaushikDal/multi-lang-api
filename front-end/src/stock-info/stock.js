import React from 'react'
import Loader from '../loader/loader'
import "./stock.css"

const Stock = ({stock,isLoading,id}) => {
  // console.log(stock)
  // TODO: Map stockInfo
  return (
    
      stock.map((info,idx )=> (
        <div key={`Investor ${id} stock ${idx} `} className="stock-tile">
            <h4>{info.symbol}</h4>
            {isLoading ? (<Loader />) : (
            <span>
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