import {useContext, useEffect,useState} from "react"
import {getAllInvestors, getAllStocks} from "../api/api-requests"
import AppContext from '../context/context';
import Investor from '../investor-block/investor';

import Modal from "../modal/modal";
import "./landing-page.css"

const LandingPage = () => {   
    const {setStocks,setLoading,investorList,setInvestorList,setRefreshedInvestor} = useContext(AppContext)
    const [showModal, setShowModal] = useState(false)
    const [fetchTime,setFetchTime] = useState()

    useEffect (() => {
        setLoading(true)
        getAllInvestors().then(resp => resp && setInvestorList(resp.data))
        setLoading(false)
      },[setLoading,setInvestorList])
    
      const refreshPrices = async() =>{
        setLoading(true)
        setRefreshedInvestor('all')
        const startTime = performance.now()
        const resp = await getAllStocks()
        const endTime = performance.now()
        if (!resp.data) return;
        setStocks(resp.data)
        setLoading(false) 
        setFetchTime(endTime-startTime)
      }


  
  return (
    <section className="landing-page">
        <h1>Stocks MApp</h1>
        <div className="landing-page__cta">
          <button onClick={refreshPrices}>Refresh Prices</button>
          <button onClick={() => setShowModal(true)}>Create Investor</button>
        </div>
        {fetchTime && <h5>Fetched all stocks in {parseFloat(fetchTime/1000).toFixed(2)}s</h5>}
        {
          investorList.map(investor => (
            <Investor investor={investor} setFetchTime={setFetchTime} />
          )  
        )}
        {showModal && <Modal setShowModal={setShowModal} />}
    </section>
  )
}

export default LandingPage