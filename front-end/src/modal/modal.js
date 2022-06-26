import {useState,useContext} from 'react'
import { createInvestor } from '../api/api-requests'
import AppContext from '../context/context'
import "./modal.css"

const Modal = ({setShowModal}) => {
    const [name,setName] = useState("")
    const [stockList,setStockList] = useState("")
    const {setInvestorList} = useContext(AppContext)

    const handleSubmit = (e) => {
        e.preventDefault()
        const stockArray = stockList.split(",").map(stock => {return {symbol:stock}})
        console.log(stockArray)
        createInvestor(name,stockArray).then(resp => console.log("Resp:",resp) && setInvestorList(resp))
        setName("")
        setStockList("")
    }

  return (
    <div className='modal__container'>
        <form className='modal__form' onSubmit={handleSubmit}>
            <label htmlFor='name-box'>Investor Name</label>
            <input type="text" id="name-box" onChange={(e) => setName(e.target.value)} value={name} />
            <label htmlFor='stocks-box'>Stocks</label>
            <input type="text" id="stocks-box" onChange={(e) => setStockList(e.target.value)} value={stockList} />
            <div className='modal__buttons'>
                <input type="submit"/>
                <button onClick={() => setShowModal(false)}>Close</button>
            </div>
        </form>
    </div>
  )
}

export default Modal