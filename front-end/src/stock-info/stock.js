import React from 'react'
import "./stock.css"

const Stock = () => {
  return (
    <div>
        <h4>{stockName}</h4>
        <span>
            {stock.Price}
            {stock.Change}
        </span>
    </div>
  )
}

export default Stock